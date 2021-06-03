package candid

import (
	"fmt"
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	spec "github.com/internet-computer/candid-go/internal/grammar"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ParseMotoko parses the Motoko (.mo) files and returns the Program that is defined in it.
func ParseMotoko(path string) (Program, error) {
	moc, err := exec.LookPath("moc")
	if err != nil {
		return Program{}, err
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
		return Program{}, err
	}
	didFileName := fmt.Sprintf("%s/%s.did", dir, strings.TrimSuffix(base, ".mo"))
	defer func() {
		_ = os.RemoveAll(didFileName)
	}()
	raw, err := ioutil.ReadFile(didFileName)
	if err != nil {
		return Program{}, err
	}
	return ParseDID(raw)
}

// ParseDID parses the given raw .did files and returns the Program that is defined in it.
func ParseDID(raw []byte) (Program, error) {
	p, err := ast.New(raw)
	if err != nil {
		return Program{}, err
	}
	n, err := spec.Prog(p)
	if err != nil {
		return Program{}, err
	}
	if _, err := p.Expect(parser.EOD); err != nil {
		return Program{}, err
	}
	return convertProgram(n), nil
}
