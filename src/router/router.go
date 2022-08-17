package router

import (
	router "book-api/src/router/routes"

	"github.com/gorilla/mux"
)

// Retornando rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return router.Configurate(r)
}
