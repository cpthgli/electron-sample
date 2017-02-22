package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	PostText string
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

func sliceToString(sliceString []string, unspecified string) string {
	s := ""
	switch {
	case len(sliceString) == 0:
		s += unspecified
	case len(sliceString) == 1 && sliceString[0] == "":
		s += unspecified
	default:
		for _, v := range sliceString {
			s += v
		}
	}
	return s
}

func (p *HomePage) setPage(r *http.Request, args []string) {
	r.ParseForm()
	p.Template = "index.html"
	p.Title = "Home"
	p.Heading = "Hello, Electron + Go!"
	p.Node = sliceToString(r.Form["node"], "unspecified")
	p.Chrome = sliceToString(r.Form["chrome"], "unspecified")
	p.Electron = sliceToString(r.Form["electron"], "unspecified")
	p.Go = runtime.Version()[2:]
	p.PostText = sliceToString(r.Form["text"], "unspecified")
}
func (p *AboutPage) setPage() {
	p.Template = "about.html"
	p.Title = "About"
	p.Heading = "About page"
}
func (p *ErrorPage) setPage(status int, messege string) {
	p.Template = "error.html"
	p.Title = "Error"
	p.Heading = "Error"
	p.Status = status
	p.Messege = messege
}

func statusLog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Printf("\n%s\n", "===============================Request================================")
	fmt.Println("Method           :", r.Method)
	fmt.Println("URL              :", r.URL)
	fmt.Println("Proto            :", r.Proto)
	fmt.Println("ProtoMajor       :", r.ProtoMajor)
	fmt.Println("ProtoMinor       :", r.ProtoMinor)
	fmt.Println("Header           :", r.Header)
	fmt.Println("Body             :", r.Body)
	fmt.Println("ContentLength    :", r.ContentLength)
	fmt.Println("TransferEncoding :", r.TransferEncoding)
	fmt.Println("Close            :", r.Close)
	fmt.Println("Host             :", r.Host)
	fmt.Println("Form             :", r.Form)
	fmt.Println("PostForm         :", r.PostForm)
	fmt.Println("MultipartForm    :", r.MultipartForm)
	fmt.Println("Trailer          :", r.Trailer)
	fmt.Println("RemoteAddr       :", r.RemoteAddr)
	fmt.Println("RequestURI       :", r.RequestURI)
	fmt.Println("TLS              :", r.TLS)
	fmt.Println("Cancel           :", r.Cancel)
	fmt.Println("Response         :", r.Response)
	fmt.Printf("%s\n", "======================================================================")
}

func allHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		var page HomePage
		page.setPage(r, os.Args)
		page.handler(w, r)
	case "/about/":
		var page AboutPage
		page.setPage()
		page.handler(w, r)
	case "/favicon.ico":
		fmt.Print()
	default:
		var page ErrorPage
		page.setPage(http.StatusNotFound, "HTTP 404 not found error.")
		page.handler(w, r, http.StatusNotFound)
	}
}
func (p *HomePage) handler(w http.ResponseWriter, r *http.Request) {
	statusLog(w, r)
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
