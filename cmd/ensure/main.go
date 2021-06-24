package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const usage = `Usage:
	$ ... | ensure md5 1a79a4d60de6718e8e5b326e338ae533 | ...
	Check the hash, then pass standard input to standard output.
Options:
	-list	List all supported algorithms.
	-help	Print this help message.
	-quiet	Suppress error messages.

`

func main() {
	log.SetFlags(0)
	var (
		helpFlag  = flag.Bool("help", false, "")
		listFlag  = flag.Bool("list", false, "")
		quietFlag = flag.Bool("quiet", false, "")
	)
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
	}
	flag.Parse()
	if *quietFlag {
		log.SetOutput(ioutil.Discard)
	}
	if *helpFlag || (!*listFlag && flag.NArg() == 0) {
		fmt.Fprint(flag.CommandLine.Output(), usage)
		return
	}
	if !*listFlag && flag.NArg() != 2 {
		log.Fatalln("ERROR: must provide algorithm and digest")
	}
	(&ensure{listMode: *listFlag}).Run(flag.Args())
}

var algorithms = map[string]hash.Hash{
	"md5":    md5.New(),
	"sha1":   sha1.New(),
	"sha256": sha256.New(),
	"sha512": sha512.New(),
}

type ensure struct {
	listMode          bool
	algorithm, digest string
}

func (e *ensure) Run(args []string) {
	if e.listMode {
		e.list()
		return
	}
	e.algorithm = args[0]
	e.digest = args[1]
	h, ok := algorithms[e.algorithm]
	if !ok {
		log.Fatalf("ERROR: %s is not a supported algorithm", e.algorithm)
	}
	r := bufio.NewReader(os.Stdin)
	var saved bytes.Buffer
	tee := io.TeeReader(r, &saved)
	if _, err := io.Copy(h, tee); err != nil {
		log.Fatal(err)
	}
	res := hex.EncodeToString(h.Sum(nil))
	if res != e.digest {
		log.Fatalf("ERROR: expected %s, got %s", res, e.digest)
	}
	fmt.Printf("%s", saved.String())
}

func (e *ensure) list() {
	keys := make([]string, 0, len(algorithms))
	for k := range algorithms {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k)
	}
}
