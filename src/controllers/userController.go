package controllers

import "net/http"

func CreateUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}

func ListAllUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}

func ListUserById(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}

func UpdateUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}

func DeleteUser(req http.ResponseWriter, res *http.Request) {
	req.Write([]byte("Criando ususarios"))
}
