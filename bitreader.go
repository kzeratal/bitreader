package bitreader

import (
	"errors"
)

type BitReader struct {
	bytes    []byte
	position int
}

func NewReader(bytes []byte) *BitReader {
	return &BitReader{bytes: bytes}
}

func (r *BitReader) Read(length int) (byte, error) {
	b := byte(0)
	if length > 8 || length < 1 {
		return b, errors.New("cannot read more than 8 bits or less than 1 bit")
	}
	for i := 0; i < length; i++ {
		if r.position == len(r.bytes) * 8 {
			break
		}
		b = b << 1 | r.bytes[r.position / 8] >> (7 - r.position % 8) & 1
		r.position++
	}
	return b, nil
}