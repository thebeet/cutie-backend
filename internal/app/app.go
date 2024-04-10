package app

import "net/http"

type App struct {
}

type Opts struct {
}

func New(opts Opts) *App {
	return &App{}
}

func (a *App) Bootstrap(mux *http.ServeMux) {
	if mux != nil {
		mux.HandleFunc("GET /hello", a.hello)
	}
}

func (a *App) Shutdown() {
}
