package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dir  = kingpin.Flag("directory", "Path to dir which has to be served.").Required().Short('d').String()
	port = kingpin.Flag("port", "Port to run at").Default("8080").String()
)

func main() {
	kingpin.Version("0.5")
	kingpin.Parse()

	s := Server{
		port:    ":" + *port,
		dirPath: *dir,
	}
	s.Serve()
}
