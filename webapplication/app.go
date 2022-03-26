package webapplication

import (
	"fmt"
	"net/http"

	"github.com/yesleymiranda/go-toolkit/logger"

	"github.com/gorilla/mux"
)

type App struct {
	port   string
	Router *mux.Router
}

func New(port string) *App {
	return &App{
		port:   port,
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (app *App) Initialize() {
	logger.Init()
	logger.Info("start app...")
}

func (app *App) ListenAndServe() error {
	address := fmt.Sprintf(":%v", app.port)
	logger.Info(fmt.Sprintf("app on address:%v", address))
	return http.ListenAndServe(address, app.Router)
}
