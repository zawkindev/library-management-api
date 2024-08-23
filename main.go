package main

import (
	"fmt"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", mainMux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome man!")
}
