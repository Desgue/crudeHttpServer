package main

import (
	"log"

	"github.com/Desgue/crudeHttpServer/http"
)

func main() {
	server := http.NewHTTPServer(":3000")
	log.Fatal(server.ListenAndServe())
}
