package main

import (
	"log"

	"github.com/Desgue/crudeHttpServer/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", []string{"text/html"})
		w.WriteHeader(200)
		w.Write([]byte("<h1>Hello World</h1>"))
	})

	server := http.NewHTTPServer(":3000", nil)

	log.Fatal(server.ListenAndServe())
}
