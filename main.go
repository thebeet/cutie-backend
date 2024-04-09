package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/thebeet/cutie-backend/pkg/app"
)

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	webApp := app.New(app.Opts{})
	webApp.Bootstrap(r)

	r.Run()

	webApp.Shutdown()
}
