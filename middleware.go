package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				publicbyte, err := ioutil.ReadFile("./public.rsa.pub")
				if err != nil {
					fmt.Println("Error al leer archivo public")
				}

				publickey, err := jwt.ParseRSAPublicKeyFromPEM(publicbyte)
				if err != nil {
					fmt.Println("Error al convertir public key")
				}
				fmt.Println(publickey)
				return publickey, nil
			})

			if err != nil {
				//fmt.Println("error en la validacion")
			}

			if token.Valid {
				w.WriteHeader(http.StatusAccepted)
				//fmt.Fprintln(w, "aceptado")
				f(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "no aceptado")
				return
			}

		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
