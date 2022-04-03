package webapplication

import (
	"fmt"
	"net/http"
)

func getPing(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "pong")
}

func NewPingHandler(app *App) {
	app.Router.HandleFunc("/ping", getPing).Methods("GET")
}
