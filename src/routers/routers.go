package routers

import (
	"github.com/gorilla/mux"
	"github.com/njoysubho/go-mux-example/src/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Routers ...
func Routers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", handlers.BookAndTitle)
	r.Handle("/metrics", promhttp.Handler())
	return r
}
