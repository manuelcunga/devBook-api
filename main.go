package main

import (
	"book-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Bem vindo jovem")

	app := router.Gerar()
	log.Fatal(http.ListenAndServe(":4000", app))

}
