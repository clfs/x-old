package main

import (
	"crypto/md5"  //nolint:gosec // Third-party digest choice is beyond my control.
	"crypto/sha1" //nolint:gosec // Third-party digest choice is beyond my control.
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"hash"
	"hash/crc32"
	"log"
	"os"

	"github.com/clfs/x/ensure"
)

func main() {
	log.SetFlags(0)

	var (
		listFlag = flag.Bool("list", false, "list supported algorithms")
		algFlag  = flag.String("alg", "", "algorithm to check with")
		sumFlag  = flag.String("sum", "", "required sum, in hex")
	)

	if *listFlag {
		log.Printf("supported algorithms:\n\tcrc32, crc32c, md5, sha1, sha256, sha512")
		return
	}

	expected, err := hex.DecodeString(*sumFlag)
	if err != nil {
		log.Fatalf("invalid sum: %v", err)
		return
	}

	var h hash.Hash

	switch *algFlag {
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
		flag.Usage()
		return
	}

	c := ensure.NewCopier(h)
	if _, err := c.Copy(os.Stdout, os.Stdin, expected); err != nil {
		log.Fatalf("error: %v", err)
	}
}
