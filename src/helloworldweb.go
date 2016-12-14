package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

type Response struct {
	Greeting string
}

func main() {
	http.HandleFunc("/english", helloEnglish)
	http.HandleFunc("/chinese", helloChinese)

	addr := "localhost:9999"
	log.Printf("Starting server on http://%s", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}

func helloEnglish(w http.ResponseWriter, r *http.Request){
	if err := json.NewEncoder(w).Encode(Response{Greeting: "Hello World"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}
	simpleLog(w, r)
}

func helloChinese(w http.ResponseWriter, r *http.Request){
	if err := json.NewEncoder(w).Encode(Response{Greeting:"Hello World"}); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}
	simpleLog(w, r)
}

func simpleLog(w http.ResponseWriter, r *http.Request){

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	uri := r.URL.RequestURI()

	referer := r.Referer()
	if referer == ""{
		referer = "-"
	}

	userAgent := r.UserAgent()
	if referer == ""{
		userAgent = "-"
	}

	username := "-"

	if u, _, ok := r.BasicAuth(); ok {
		username = u
	}

	fields := []string{
		host,
		username,
		r.Method,
		uri,
		r.Proto,
		referer,
		userAgent,
	}

	log.Println(strings.Join(fields, " "))
}
