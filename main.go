package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var fdir string
	flag.StringVar(&fdir, "d", "", "path to dir to serve")
	flag.Parse()

	if fdir == "" {
		printHelp()
		os.Exit(0)
	}

	s := Server{
		port:    ":8080",
		dirPath: fdir,
	}
	s.Serve()
}

func printHelp() {
	fmt.Println("Usage: -d path to dir to serve")
}
