package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dir = kingpin.Flag("directory", "Path to dir which has to be served.").Required().Short('d').String()
)

func main() {
	kingpin.Version("0.5")
	kingpin.Parse()

	s := Server{
		port:    ":8080",
		dirPath: *dir,
	}
	s.Serve()
}
