package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"text/template"
)

// Page is page default struct
type Page struct {
	Template string
	Title    string
	Heading  string
}

// HomePage is home page struct
type HomePage struct {
	Page
	Node     string
	Chrome   string
	Electron string
	Go       string
}

// AboutPage is about page struct
type AboutPage struct {
	Page
}

// ErrorPage is error page struct
type ErrorPage struct {
	Page
	Status  int
	Messege string
}

func (p *HomePage) setPage(args ...string) {
	p.Template = "index.html"
	p.Title = "Home"
	p.Heading = "Hello, Electron + Go!"
	p.Go = runtime.Version()[2:]
	switch len(args) {
	case 3:
		p.Electron = args[2]
		fallthrough
	case 2:
		p.Chrome = args[1]
		fallthrough
	case 1:
		p.Node = args[0]
	default:
		p.Node = "undefined"
		p.Chrome = "undefined"
		p.Electron = "undefined"
	}
}
func (p *AboutPage) setPage() {
	p.Template = "about.html"
	p.Title = "About"
	p.Heading = "About page"
}
func (p *ErrorPage) setPage(status int) {
	p.Template = "error.html"
	p.Title = "Error"
	p.Heading = "Error"
	p.Status = status
	if status == http.StatusNotFound {
		p.Messege += "Page Note Found."
	} else {
		p.Messege = ""
	}
}

func statusLog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Request :", r)
	fmt.Println("ResponseWriter :", w)
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		var page HomePage
		page.setPage()
		page.handler(w, r)
	case "/about":
		fallthrough
	case "/about/":
		var page AboutPage
		page.setPage()
		page.handler(w, r)
	case "/favicon.ico":
		fmt.Print()
	default:
		var page ErrorPage
		page.setPage(http.StatusNotFound)
		page.handler(w, r, http.StatusNotFound)
	}
}
func (p *HomePage) handler(w http.ResponseWriter, r *http.Request) {
	statusLog(w, r)
	fmt.Println(*p)
	tmpl, err := template.ParseFiles(p.Template)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, *p)
	if err != nil {
		panic(err)
	}
}
func (p *AboutPage) handler(w http.ResponseWriter, r *http.Request) {
	statusLog(w, r)
	fmt.Println(*p)
	tmpl, err := template.ParseFiles(p.Template)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, *p)
	if err != nil {
		panic(err)
	}
}
func (p *ErrorPage) handler(w http.ResponseWriter, r *http.Request, status int) {
	statusLog(w, r)
	fmt.Println(*p)
	tmpl, err := template.ParseFiles(p.Template)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, *p)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", allHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
