package main

import (
	"bufio"
	"bytes"
	"crypto/md5"  //nolint:gosec // Third-party digest choice is beyond my control.
	"crypto/sha1" //nolint:gosec // Third-party digest choice is beyond my control.
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	var (
		listFlag               bool
		algorithmFlag, sumFlag string
	)

	flag.BoolVar(&listFlag, "l", false, "list available algorithms")
	flag.BoolVar(&listFlag, "list", false, "list available algorithms")
	flag.StringVar(&algorithmFlag, "a", "", "the algorithm to use")
	flag.StringVar(&algorithmFlag, "algorithm", "", "the algorithm to use")
	flag.StringVar(&sumFlag, "s", "", "the expected sum, in hex")
	flag.StringVar(&sumFlag, "sum", "", "the expected sum, in hex")
	flag.Parse()

	if listFlag {
		log.Printf("Supported algorithms:\n\tcrc32, crc32c, md5, sha1, sha256, sha512")
		return
	}

	if algorithmFlag == "" || sumFlag == "" {
		flag.Usage()
		return
	}

	expectedSum, err := hex.DecodeString(sumFlag)
	if err != nil {
		log.Fatalf("invalid sum: %v", err)
		return
	}

	e := ensure{
		algorithm:   algorithmFlag,
		expectedSum: expectedSum,
	}

	if err := e.Run(); err != nil {
		log.Fatalln(err)
	}
}

type ensure struct {
	algorithm   string
	expectedSum []byte
}

func (e *ensure) Run() error {
	var (
		h hash.Hash
		r = bufio.NewReader(os.Stdin)
		b bytes.Buffer
	)

	switch e.algorithm {
	case "crc32":
		h = crc32.New(crc32.IEEETable)
	case "crc32c":
		h = crc32.New(crc32.MakeTable(crc32.Castagnoli))
	case "md5":
		h = md5.New() //nolint:gosec // Third-party digest choice is beyond my control.
	case "sha1":
		h = sha1.New() //nolint:gosec // Third-party digest choice is beyond my control.
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		return &invalidAlgorithmError{algorithm: e.algorithm}
	}

	tee := io.TeeReader(r, &b)
	if n, err := io.Copy(h, tee); err != nil {
		return &failedCopyError{byteCount: n}
	}

	computedSum := h.Sum(nil)
	if !bytes.Equal(e.expectedSum, computedSum) {
		return &mismatchedSumError{expected: e.expectedSum, computed: computedSum}
	}

	fmt.Printf("%s", b.String())

	return nil
}

type invalidAlgorithmError struct {
	algorithm string
}

func (e *invalidAlgorithmError) Error() string {
	return fmt.Sprintf("invalid algorithm: %s", e.algorithm)
}

type failedCopyError struct {
	byteCount int64
}

func (e *failedCopyError) Error() string {
	return fmt.Sprintf("failed to copy all input: %d bytes copied", e.byteCount)
}

type mismatchedSumError struct {
	expected, computed []byte
}

func (e *mismatchedSumError) Error() string {
	return fmt.Sprintf("wrong sum: expected %x, computed %x", e.expected, e.computed)
}
