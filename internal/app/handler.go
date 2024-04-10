package app

import (
	"fmt"
	"net/http"
)

func (a *App) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "world")
}
