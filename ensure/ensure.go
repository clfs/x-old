package ensure

import (
	"bytes"
	"hash"
	"io"
)

// Copier is the interface that wraps the Copy method.
//
// Implementations are free to implement any sum algorithm of choice.
type Copier interface {
	Copy(dst io.Writer, src io.Reader, sum []byte) (int64, error)
}

type copier struct {
	h hash.Hash
}

// NewCopier returns a new Copier.
func NewCopier(h hash.Hash) Copier {
	return &copier{h: h}
}

// Copy copies from src to dst in the same manner as io.Copy, except
// that the copy is only performed if src has the correct sum.
func (c *copier) Copy(dst io.Writer, src io.Reader, sum []byte) (int64, error) {
	defer c.h.Reset()

	if len(sum) == 0 {
		return 0, &NoSumError{}
	}

	b := new(bytes.Buffer)
	tee := io.TeeReader(src, b)

	if _, err := io.Copy(c.h, tee); err != nil {
		return 0, &FailedReadError{Err: err}
	}

	if got := c.h.Sum(nil); !bytes.Equal(sum, got) {
		return 0, &WrongSumError{Want: sum, Got: got}
	}

	n, err := b.WriteTo(dst)
	if err != nil {
		return n, &FailedWriteError{Err: err}
	}

	return n, nil
}
