package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI          string
	Methods      string
	Function     func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

func Configurate(router *mux.Router) *mux.Router {
	rota := userRouter

	for route := range rota {
		router.HandleFunc(route.URI, route.Function).Methods(route.Methods)
	}

	return router

}
