// Package ensure checks piped data for the correct hash.
package main

import (
	"bufio"
	"bytes"
	"crypto/md5"  //nolint:gosec
	"crypto/sha1" //nolint:gosec
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

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tensure -help\n")
	fmt.Fprintf(os.Stderr, "\tensure -list\n")
	fmt.Fprintf(os.Stderr, "\t... | ensure -alg crc32 -sum 2747fc56 | ...\n")
	os.Exit(1)
}

func list() {
	fmt.Printf("supported algorithms (-alg):\n")
	fmt.Printf("\tcrc32, crc32c, md5, sha1, sha256, sha512\n")
	os.Exit(0)
}

func main() {
	log.SetFlags(0)

	var (
		listFlag = flag.Bool("list", false, "")
		algFlag  = flag.String("alg", "", "")
		sumFlag  = flag.String("sum", "", "")
	)

	flag.Usage = usage
	flag.Parse()

	if *listFlag {
		list()
	}

	if *sumFlag == "" {
		flag.Usage()
	}

	var h hash.Hash

	switch *algFlag {
	case "crc32":
		h = crc32.New(crc32.IEEETable)
	case "crc32c":
		h = crc32.New(crc32.MakeTable(crc32.Castagnoli))
	case "md5":
		h = md5.New() //nolint:gosec
	case "sha1":
		h = sha1.New() //nolint:gosec
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		flag.Usage()
	}

	expected, err := hex.DecodeString(*sumFlag)
	if err != nil {
		log.Fatalf("invalid sum: %v", err)
	}

	var (
		r = bufio.NewReader(os.Stdin)
		w = bufio.NewWriter(os.Stdout)
		b = new(bytes.Buffer)
		t = io.TeeReader(r, b)
	)

	if _, err := io.Copy(h, t); err != nil {
		log.Fatal("failed to hash stdin")
	}

	if got := h.Sum(nil); !bytes.Equal(got, expected) {
		log.Fatalf("mismatched sum: got %x", got)
	}

	if _, err := b.WriteTo(w); err != nil {
		log.Fatal("failed to write to stdout")
	}

	w.Flush()
}
