package controllers

import "net/http"

func CreateUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}

func ListAllUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Buscando ususarios"))
}

func ListUserById(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("buscando um ususario"))
}

func UpdateUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("atualiando um ususario"))
}

func DeleteUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("deletando ususario"))
}
