package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		dir   = kingpin.Flag("directory", "Path to dir which has to be served.").Required().Short('d').String()
		port  = kingpin.Flag("port", "Port to run at").Default("8080").Short('p').String()
		cors  = kingpin.Flag("cors", "Add CORS headers").Short('c').StringMap()
		cache = kingpin.Flag("cache", "Add Cache headers").StringMap()
	)

	kingpin.Version("0.5")
	kingpin.Parse()

	s := Server{
		port:    ":" + *port,
		dirPath: *dir,
		cors:    *cors,
		cache:   *cache,
	}
	s.serve()
}
