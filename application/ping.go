package application

import (
	"fmt"
	"net/http"
)

func NewPingHandler(app *App) {
	app.Router.Get("/ping", getPing)
}

func getPing(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "pong")
}
