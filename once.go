package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var seen = make(map[string]bool)

func once(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if seen[line] == false {
			seen[line] = true
			fmt.Println(line)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: once [file...]\n")
	os.Exit(2)
}

func main() {
	log.SetPrefix("once: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		once(os.Stdin)
	} else {
		for _, arg := range flag.Args() {
			f, err := os.Open(arg)
			if err != nil {
				log.Print(err)
				continue
			}
			once(f)
			f.Close()
		}
	}
}
