package main

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
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

type Claims struct {
	Username Empleado
	jwt.StandardClaims
}

type Empleado struct {
	Id          string `json:"id"`
	Nombre      string `json:"nombre"`
	Apellidos   string `json:"apellidos"`
	Telefono    string `json:"telefono"`
	Correo      string `json:"correo"`
	FechaInicio string `json:"fechaInicio"`
	FechaFin    string `json:"fechaFin"`
	Contrasena  string `json:"contrasena"`
	Admin       string `json:"admin"`
}

type Semana struct {
	Lunes     string `json:"lunes"`
	Martes    string `json:"martes"`
	Miercoles string `json:"miercoles"`
	Jueves    string `json:"jueves"`
	Viernes   string `json:"viernes"`
}

type ResponseToken struct {
	Token string `json:"token"`
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

func (r ResponseToken) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

type MetaData interface{}
