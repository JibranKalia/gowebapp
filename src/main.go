package main

import (
	"net/http"
	"github.com/gorilla/mux"

)
func sayhelloName(w http.ResponseWriter, r *http.Request) {

}

func main() {
	serveWeb()
}

func serveWeb(){

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", serveContent)
	gorillaRoute.HandleFunc("/{page_alias}", serveContent)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}

func serveContent(w http.ResponseWriter, r *http.Request){
	urlParams := mux.Vars(r)
	page_alias := urlParams["page_alias"]
	if page_alias == "" {
		page_alias = "Home"
	}

	w.Write([]byte("Hello World! Here is the test: " + page_alias))
}
