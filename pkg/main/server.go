package main

import (
	"net/http"
)

var domainMap = map[string] http.Handler {
}


func main() {

	// static website paths
	http.Handle("/creative/", http.StripPrefix("/creative/", http.FileServer(http.Dir("./templates/startbootstrap-creative"))))
	http.Handle("/grayscale/", http.StripPrefix("/grayscale/", http.FileServer(http.Dir("./templates/startbootstrap-grayscale"))))


	http.HandleFunc("/", ping)
	http.ListenAndServe(":3040", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("healthy"))
}


