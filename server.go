package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type headers map[string]string

// Server base struct
type Server struct {
	port    string
	dirPath string
	cors    headers
	cache   headers
}

var dCacheHdrs = headers{
	"control": "no-cache",
}

var dCorsHdrs = headers{
	"origin":  "*",
	"methods": "*",
	"headers": "*",
}

func cacheHandler(hdrs headers, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range hdrs {
			w.Header().Add("Cache-"+strings.Title(k), v)
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func corsHandler(hdrs headers, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range hdrs {
			w.Header().Add("Access-Control-Allow-"+strings.Title(k), v)
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (s *Server) serve() {
	s.parseArgs()

	mux := http.NewServeMux()
	mux.Handle("/", cacheHandler(s.cache, corsHandler(s.cors, s.handle())))

	log.Print("Listening on", s.port)
	log.Fatal(http.ListenAndServe(s.port, mux))
}

func (s *Server) handle() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := s.dirPath + r.URL.Path
		file, err := os.Lstat(path)
		if err != nil {
			log.Println("File " + path + " not exists")
			http.NotFound(w, r)
			return
		}

		switch mode := file.Mode(); {
		case mode.IsDir():
			handleDir(path, w, r)
		case mode.IsRegular():
			handleFile(path, w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func (s *Server) parseArgs() {
	for k, v := range dCorsHdrs { // parse cors into headers
		if s.cors[k] == "" { // if empty
			s.cors[k] = v
		}
	}

	for k, v := range dCacheHdrs {
		if s.cache[k] == "" { // if empty
			s.cache[k] = v
		}
	}
}

func handleDir(path string, w http.ResponseWriter, r *http.Request) {
	log.Println("Serving " + path + " dir...")
	dirList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	var outputDirList []string
	for _, file := range dirList {
		fileURL := strings.TrimRight(r.RequestURI, "/") + "/" + file.Name()
		outputDirList = append(outputDirList, "<a href=\""+fileURL+"\">"+file.Name()+"</a>")
	}
	response := strings.Join(outputDirList, "<br />")
	w.Write([]byte(string(response)))
}

func handleFile(path string, w http.ResponseWriter, r *http.Request) {
	log.Println("Serving " + path + " file...")

	http.ServeFile(w, r, path)
}
