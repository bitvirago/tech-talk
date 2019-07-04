package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`Bitrise tech talk1`))
	})
	log.Fatal(http.ListenAndServe(":9999", nil))
}
