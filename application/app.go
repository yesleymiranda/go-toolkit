package application

import (
	"github.com/go-chi/chi"
	"github.com/yesleymiranda/go-toolkit/logger"
)

type Config struct {
	Port     string
	WithPing bool
}

type App struct {
	port     string
	withPing bool
	Router   *chi.Mux
}

// New create a new web application
func New(config *Config) *App {
	return &App{
		port:     config.Port,
		withPing: config.WithPing,
		Router:   chi.NewRouter(),
	}
}

// Initialize initialize a app
// Logger rs/zerolog
// GET ping handler
func (app *App) Initialize() {
	logger.Init()
	logger.Info("start app...")

	if app.withPing {
		NewPingHandler(app)
	}
}
