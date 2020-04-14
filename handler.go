package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CheckAut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entro check")
	decoder := json.NewDecoder(r.Body)
	var user Login
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("error %v", err)
		return
	}

	fmt.Println("entro ligin")
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("desde hanedr : " + user.Usuario + user.Contrasena)
	Emple := ValidaUsuario(user.Usuario, user.Contrasena)
	if Emple.Id == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("El usuario no existe"))
		return
	} else {
		result := ResponseToken{GenerateJWT(Emple)}
		token, err := result.ToJson()
		if err != nil {
			fmt.Println("error en la conversion del json del token")
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(token)
	}

}

func GenerateJWT(user Empleado) string {
	privatebyte, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		fmt.Println("Error al leer archivo private")
	}

	privatekey, err := jwt.ParseRSAPrivateKeyFromPEM(privatebyte)
	if err != nil {
		fmt.Println(err)
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	//	w.Write(bye)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "log in",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(privatekey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println(err)
		return "nil"
	}
	return tokenString
}

func ValidaToken(w http.ResponseWriter, r *http.Request) {

}

func RegistroEmpleado(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user Empleado
	fmt.Println("el id es :" + user.Id)

	err := decoder.Decode(&user)
	if err != nil {
		return
	}

	if user.Id == "" {
		if RegistraEmpleado(user) {
			fmt.Fprintf(w, "true")
		}
	} else {

		if "" == user.Contrasena {
			ActulizaEmpleadoSinContrasena(user)
		} else {
			ActulizaEmpleado(user)
		}

		fmt.Fprintf(w, "true")

	}

}

func GetEmpleados(w http.ResponseWriter, r *http.Request) {

	resultado := "["
	var user []Empleado = SeleccionaEmpleados()
	for _, result := range user {
		JSON, err := json.MarshalIndent(result, "", "\t")
		resultado = resultado + string(JSON) + ","
		if err != nil {
			fmt.Println("error:", err)
		} else {
			w.Header().Set("Content-Type", "application/jsons")
			//w.Write(JSON)
		}
	}
	resultado = resultado[:len(resultado)-1] + "]"
	w.Write([]byte(resultado))
	//fmt.Println(resultado)
}
func GetEmpleado(w http.ResponseWriter, r *http.Request) {

	cuerpoRespuesta, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}
	respuestaString := string(cuerpoRespuesta)
	var user = SeleccionaEmpleado(respuestaString)
	JSON, err := user.ToJson()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		w.Header().Set("Content-Type", "application/jsons")
		w.Write(JSON)
	}

}

func BorrarEmpleado(w http.ResponseWriter, r *http.Request) {
	fmt.Println("paso por borrar")
	cuerpoRespuesta, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error leyendo respuesta: %v", err)
	}

	respuString := string(cuerpoRespuesta)
	fmt.Println("la respuesta para borrar: " + respuString)

	BorraEmpleado(respuString)
	GetEmpleados(w, r)

}

func RegistraEntrada(w http.ResponseWriter, r *http.Request) {
	var user Empleado
	if json.NewDecoder(r.Body).Decode(&user) != nil {
		fmt.Println("error al decodificar")
		return
	}
	if EntradaRegistrada(user.Id) == "" {
		RegistroEntrada(user)
		fmt.Println("entrada")
		return
	}

	RegistroSalida(user, EntradaRegistrada(user.Id))
	fmt.Println("salida")

}

func checo(w http.ResponseWriter, r *http.Request) {
	var user Empleado
	if json.NewDecoder(r.Body).Decode(&user) != nil {
		fmt.Println("error al decodificar")
		return
	}
	w.Header().Set("Content-Type", "application/jsons")
	if EntradaRegistrada(user.Id) == "" {
		fmt.Println("false ")
		w.Write([]byte("false"))
		return
	}
	fmt.Println("true")
	w.Write([]byte("true"))

}
