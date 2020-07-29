package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// Server base struct
type Server struct {
	port    string
	dirPath string
}

func (s *Server) initHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := s.dirPath + r.URL.Path
		file, err := os.Lstat(path)
		if err != nil {
			log.Println("File " + path + " not exists")
			http.NotFound(w, r)
			return
		}

		switch mode := file.Mode(); {
		case mode.IsRegular():
			serveFile(w, r, path)
		case mode.IsDir():
			serveDir(w, r, path)
		}
	})

	log.Printf("Serving %v directory\n", s.dirPath)
}

// Serve function
func (s *Server) Serve() {
	s.initHandler()

	log.Print("Listening on", s.port)
	log.Fatal(http.ListenAndServe(s.port, nil))
}

func serveDir(w http.ResponseWriter, r *http.Request, path string) {
	log.Println("Serving " + path + " dir...")
	dirList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var outputDirList []string
	for _, file := range dirList {
		fileURL := strings.TrimRight(r.RequestURI, "/") + "/" + file.Name()
		outputDirList = append(outputDirList, "<a href=\""+fileURL+"\">"+file.Name()+"</a>")
	}

	response := strings.Join(outputDirList, "<br />")
	w.Write([]byte(string(response)))
}

func serveFile(w http.ResponseWriter, r *http.Request, path string) {
	log.Println("Serving " + path + " file...")

	http.ServeFile(w, r, path)
}
