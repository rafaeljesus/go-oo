package main

import (
	"github.com/rafaeljesus/go-oo/api"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	log := log.New(os.Stdout, "web ", log.LstdFlags)
	app := &api.App{mux, log}

	mux.HandleFunc("/v1/healthz", app.Index)

	http.ListenAndServe(":3000", app)
}
