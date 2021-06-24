package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	var outputFlag string

	flag.StringVar(&outputFlag, "output", "render.png", "the file to output to")
	flag.Parse()

	if outputFlag == "" {
		flag.Usage()
		return
	}

	f, err := os.Create(outputFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := Render{w: f}
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
