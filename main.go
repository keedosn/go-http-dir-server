package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dir = kingpin.Flag("directory", "Path to dir which has to be served.").Short('d').String()
)

func main() {
	kingpin.Parse()

	if *dir == "" {
		printHelp()
		os.Exit(0)
	}

	s := Server{
		port:    ":8080",
		dirPath: *dir,
	}
	s.Serve()
}

func printHelp() {
	fmt.Println("Usage: -d path to dir to serve")
}
