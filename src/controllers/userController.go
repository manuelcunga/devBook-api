package controllers

import (
	"book-api/src/database"
	"book-api/src/entities"
	"book-api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(req.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var usuario entities.User

	if erro = json.Unmarshal(bodyRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connection()

	if erro != nil {
		log.Fatal(erro)
	}

	repo := repositories.NewUserRepository(db)

	repo.Create(usuario)
}

func ListAllUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Buscando ususarios"))
}

func ListUserById(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("buscando um ususario"))
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("atualiando um ususario"))
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("deletando ususario"))
}
