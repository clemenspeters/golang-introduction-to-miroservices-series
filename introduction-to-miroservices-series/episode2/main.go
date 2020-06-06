package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clemenspeters/golang-tutorials/introduction-to-miroservices-series/episode2/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
