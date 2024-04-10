package main

import (
	"flag"
	"net/http"

	"github.com/thebeet/cutie-backend/internal/app"
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	webApp := app.New(app.Opts{})
	webApp.Bootstrap(mux)

	http.ListenAndServe("localhost:8080", mux)

	webApp.Shutdown()
}
