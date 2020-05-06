package decorator

import (
	"bytes"
	"io"
)

// Декоратор перевода букв в нижний регистр.
type lowCaseReader struct {
	reader io.Reader
}

// Считать данные.
func (r *lowCaseReader) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	lower := bytes.ToLower(p)

	copy(p, lower)

	return n, err
}

// NewLowCaseReader Задекорировать Reader декоратором lowCaseReader.
func NewLowCaseReader(reader io.Reader) io.Reader {
	r := new(lowCaseReader)
	r.reader = reader

	return r
}
