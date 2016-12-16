package main

import(
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type PageData struct {
	Title string
	Body string
}


func main() {

	rt := mux.NewRouter().StrictSlash(true)

	rt.HandleFunc("/", Index)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

func Index(w http.ResponseWriter, r *http.Request){
	pd := PageData{
		Title: "Index Page",
		Body: "This is body of the page",
	}

	tmpl, err := htmlTemplate(pd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tmpl))
}

func htmlTemplate(pd PageData) (string, error) {
	html := `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	{{.Body}}
	</body>
	</HTML>`

	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer

	if err := tmpl.Execute(&out, pd); err != nil {
		return "", err
	}

	return out.String(), nil
}
