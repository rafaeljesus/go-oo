package api

import (
	"io"
	"log"
	"net/http"
)

type App struct {
	Mux *http.ServeMux
	Log *log.Logger
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.Log.Println("request to index")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Mux.ServeHTTP(w, r)
}
