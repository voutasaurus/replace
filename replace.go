package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// TODO: investigate stream search with buffer

func main() {
	flag.Parse()
	var (
		inStr   = flag.Arg(0)
		outStr  = flag.Arg(1)
		inFile  = flag.Arg(2)
		outFile = flag.Arg(3)
	)
	if outFile == "" {
		outFile = inFile
	}
	f, err := os.Open(inFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	b = bytes.Replace(b, []byte(inStr), []byte(outStr), -1)
	if err := ioutil.WriteFile(outFile, b, 0644); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
