package main

import (
	"book-api/src/config"
	"book-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Loading()
	fmt.Sprintf(config.ConnectionStringDb)
	fmt.Println("server running...")

	app := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), app))

}
