package main

import (
	"bufio"
	"fmt"
	"os"

	"log"

	"github.com/jessevdk/go-flags"
	anomlog "github.com/m-mizutani/anomlog/lib"
)

type options struct {
	InputModel  string `short:"i" long:"input"`
	OutputModel string `short:"o" long:"output"`
}

func readFile(fpath string, stream *anomlog.Stream) error {
	fp, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Fail to open file: ", fpath, " ", err)
		return err
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	for s.Scan() {
		text := s.Text()
		if len(text) > 0 {
			stream.Read(text)
		}
	}

	return nil
}

func main() {
	var opts options

	args, ParseErr := flags.ParseArgs(&opts, os.Args)
	if ParseErr != nil {
		os.Exit(1)
	}

	stream := anomlog.NewStream()

	if opts.InputModel != "" {
		stream.Load(opts.InputModel)
	}

	for _, fpath := range args[1:] {
		log.Println("Reading file... ", fpath)
		readFile(fpath, &stream)
	}

	log.Println("Done")
	for idx, format := range stream.Formats() {
		fmt.Printf("[%2d] %s\n", idx, format.String())
	}

	if opts.OutputModel != "" {
		stream.Save(opts.OutputModel)
	}
}
