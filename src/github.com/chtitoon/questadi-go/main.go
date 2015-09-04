package main
import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"html"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}