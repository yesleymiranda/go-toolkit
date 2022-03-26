package webapplication

import (
	"fmt"
	"net/http"

	"github.com/yesleymiranda/go-toolkit/logger"

	"github.com/gorilla/mux"
)

type ApplicationConfig struct {
	Port     string
	WithPing bool
}

type App struct {
	port     string
	withPing bool
	Router   *mux.Router
}

func New(config *ApplicationConfig) *App {
	return &App{
		port:     config.Port,
		withPing: config.WithPing,
		Router:   mux.NewRouter().StrictSlash(true),
	}
}

func (app *App) Initialize() {
	logger.Init()
	logger.Info("start app...")

	if app.withPing {
		NewPingHandler(app)
	}
}

func (app *App) ListenAndServe() error {
	address := fmt.Sprintf(":%v", app.port)
	logger.Info(fmt.Sprintf("app on address:%v", address))
	return http.ListenAndServe(address, app.Router)
}
