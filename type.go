package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Id         string `json:"id"`
	Nombre     string `json:"nombre"`
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

type Login struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}
type Empleado struct {
	Id          string `json:"id"`
	Nombre      string `json:"nombre`
	Apellidos   string `json:"apellidos`
	Telefono    string `json:"telefono`
	Correo      string `json:"correo`
	FechaInicio string `json:"fechaInicio`
	FechaFin    string `json:"fechaFin`
	Contrasena  string `json:"contrasena`
	Admin       string `json:"admin"`
}

func (e Empleado) ToJson() ([]byte, error) {
	return json.Marshal(e)
}

func (u User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

func (u Login) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

type MetaData interface{}
