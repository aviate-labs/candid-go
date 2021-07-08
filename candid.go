package candid

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/allusion-be/candid-go/internal/candid"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
)

// ParseMotoko parses the Motoko (.mo) files and returns the interface description that is defined in it.
func ParseMotoko(path string) (Description, error) {
	moc, err := exec.LookPath("moc")
	if err != nil {
		return Description{}, err
	}
	var (
		dir  = filepath.Dir(path)
		base = filepath.Base(path)
	)
	cmd := exec.Cmd{
		Path: moc,
		Args: []string{moc, base, "--idl"},
		Dir:  dir,
	}
	if err := cmd.Run(); err != nil {
		return Description{}, err
	}
	didFileName := fmt.Sprintf("%s/%s.did", dir, strings.TrimSuffix(base, ".mo"))
	defer func() {
		_ = os.RemoveAll(didFileName)
	}()
	raw, err := ioutil.ReadFile(didFileName)
	if err != nil {
		return Description{}, err
	}
	return ParseDID(raw)
}

// ParseDID parses the given raw .did files and returns the Program that is defined in it.
func ParseDID(raw []byte) (Description, error) {
	p, err := ast.New(raw)
	if err != nil {
		return Description{}, err
	}
	n, err := candid.Prog(p)
	if err != nil {
		return Description{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return Description{}, err
	}
	return convertDescription(n), nil
}
