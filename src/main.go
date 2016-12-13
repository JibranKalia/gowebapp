package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"text/template"
	"log"
	"os"
	"bufio"

)
func sayhelloName(w http.ResponseWriter, r *http.Request) {

}

func main() {
	serveWeb()
}

var themeName = getThemeName()
var staticPages = populateStaticPages()
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

func getThemeName(){

}

func populateStaticPages() *template.Template{
	result := templates.New("templates")
	templatePaths := new([]string)

	basePath := "pages"
	tempalateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	for _, pathInfo := range templatePathsRaw{
		log.Println(pathInfo.Name)
		*templatePaths = append(*templatePaths, basePath + "/" + pathInfo.Name())
	}

	result.ParseFiles(*templatePaths...)
	return result
}


func serveResource(w http.ResponseWriter, req *http.Request){
	path := "public" + themeName + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css"){
		contentType = "text/css, charset=utf-8"
	} else if strings.HasSuffix(path, ".png"){
		contentType = "image/png , charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg"){
		contentType = "image/jpg, charset=utf-8"
	} else if strings.HasSuffix(path, ".js "){
		contentType = "application/javascript, charset=utf-8"
	} else {
		contentType = "text/plain, charset=utf-8"
	}

	log.Println(path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.Writeto(w)
	} else {
		w.WriteHeader(404 )
	}

}
