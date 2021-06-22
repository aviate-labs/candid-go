package idl

import "math"

func concat(bs ...[]byte) []byte {
	var c []byte
	for _, b := range bs {
		c = append(c, b...)
	}
	return c
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func log2(n uint8) uint8 {
	return uint8(math.Log2(float64(n)))
}

func zeros(n uint8) []byte {
	var z []byte
	for i := 0; i < int(n); i++ {
		z = append(z, 0)
	}
	return z
}

func pad0(n uint8, bs []byte) []byte {
	for len(bs) != int(n) {
		bs = append(bs, 0)
	}
	return bs
}
