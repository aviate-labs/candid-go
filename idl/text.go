package idl

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Text struct {
	v string
	primType
}

func NewText(s string) *Text {
	return &Text{
		v: s,
	}
}

func (t *Text) Decode(r *bytes.Reader) error {
	n, err := leb128.DecodeUnsigned(r)
	if err != nil {
		return err
	}
	bs := make([]byte, n.Int64())
	i, err := r.Read(bs)
	if err != nil {
		return nil
	}
	if i != int(n.Int64()) {
		return io.EOF
	}
	*t = Text{v: string(bs)}
	return nil
}

func (t Text) EncodeType(_ *TypeTable) ([]byte, error) {
	return leb128.EncodeSigned(big.NewInt(textType))
}

func (t Text) EncodeValue() ([]byte, error) {
	bs, err := leb128.EncodeUnsigned(big.NewInt(int64(len(t.v))))
	if err != nil {
		return nil, err
	}
	return append(bs, []byte(t.v)...), nil
}

func (Text) Name() string {
	return "text"
}

func (t Text) String() string {
	return fmt.Sprintf("text: %s", t.v)
}
