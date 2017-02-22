package main

import (
	"net/http"
	"os"
	"runtime"
	"text/template"
)

// Page is web page struc
type Page struct {
	Title    string
	Heading  string
	Node     string
	Chrome   string
	Electron string
	Go       string
}

func setPage() Page {
	Title := "Hello, World!"
	Heading := "Hello, Electron + Go!"
	Node, Chrome, Electron := "undefined", "undefined", "undefined"
	Go := runtime.Version()[2:]
	if len(os.Args) == 4 {
		Node = os.Args[1]
		Chrome = os.Args[2]
		Electron = os.Args[3]
	}
	page := Page{Title, Heading, Node, Chrome, Electron, Go}
	return page
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := setPage()
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
