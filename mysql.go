package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Name      string
	Usuario   string
	Contra    string
	Url       string
	Port      string
	BaseDatos string
}

func NewConeccion() (*sql.DB, error) {
	db, err := sql.Open("mysql", "VtgHpFxzCP:V7sFv16RgJ@tcp(remotemysql.com:3306)/VtgHpFxzCP")
	if err != nil {
		return nil, err
	}
	fmt.Println("exito ")
	return db, nil
}

func ValidaUsuario(user string, pass string) bool {
	contactos := []User{}

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	fmt.Printf("el usuario es 2" + user + pass)
	filas, err := db.Query("SELECT usuario,contrasena FROM VtgHpFxzCP.Administrador where usuario= '" + user + "' and contrasena=SHA('" + pass + "')")

	if err != nil {
		fmt.Println("error en la consulta")
		return false
	}
	defer filas.Close()

	var c User

	for filas.Next() {

		err = filas.Scan(&c.Usuario, &c.Contrasena)
		if err != nil {
			fmt.Println("error al scanear")
			return false
		}

		contactos = append(contactos, c)
	}
	if len(contactos) != 0 {
		fmt.Printf("se encontro")
		return true
	}
	fmt.Printf("no se encontro")
	return false
}

func RegistraUsuario(user User) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare("INSERT INTO Administrador (nombre, Usuario, Contrasena) VALUES(?, ?, SHA(?))")
	if err != nil {
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(user.Nombre, user.Usuario, user.Contrasena)
	if err != nil {
		return false
	}
	return true
}
