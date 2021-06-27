package idl

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/allusion-be/leb128"
)

type Text string

func (Text) BuildTypeTable(*TypeTable) {}

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
	*t = Text(string(bs))
	return nil
}

func (t Text) EncodeType() []byte {
	bs, _ := leb128.EncodeSigned(textType)
	return bs
}

func (t Text) EncodeValue() []byte {
	bs, _ := leb128.EncodeUnsigned(big.NewInt(int64(len(t))))
	bs = append(bs, []byte(t)...)
	return bs
}

func (Text) Name() string {
	return "text"
}

func (t Text) String() string {
	return fmt.Sprintf("text: %s", string(t))
}
