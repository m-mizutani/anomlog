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

func readFile(fpath string, engine *anomlog.Engine) error {
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
			engine.Read(text)
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

	engine := anomlog.NewEngine()

	if opts.InputModel != "" {
		engine.Load(opts.InputModel)
	}

	for _, fpath := range args[1:] {
		log.Println("Reading file... ", fpath)
		readFile(fpath, &engine)
	}

	log.Println("Done")
	for idx, format := range engine.Formats() {
		fmt.Printf("[%2d] %s\n", idx, format.String())
	}

	if opts.OutputModel != "" {
		engine.Save(opts.OutputModel)
	}
}
