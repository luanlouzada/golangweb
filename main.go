package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello World!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Contact Page")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>FAQ Page</h1>
<ul><li>Is this page useful?</li><li>Why is this page not useful?</li></ul>
    `)
}

/*func pathHandler(w http.ResponseWriter, r *http.Request) {*/
/*switch r.URL.Path {*/
/*case "/":*/
/*homeHandler(w, r)*/
/*case "/contact":*/
/*contactHandler(w, r)*/
/*default:*/
/*http.Error(w, "Page Not Found", http.StatusNotFound)*/
/*}*/
/*}*/

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Server is listening...")
	http.ListenAndServe(":3000", router)
}
