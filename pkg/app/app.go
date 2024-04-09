package app

import "github.com/gin-gonic/gin"

type App struct {
}

type Opts struct {
}

func New(opts Opts) *App {
	return &App{}
}

func (a *App) Bootstrap(r gin.IRouter) {
	if r != nil {
		r.GET("/hello", a.hello)
	}
}

func (a *App) Shutdown() {
}
