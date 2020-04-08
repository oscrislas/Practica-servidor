package main

import (
	"database/sql"
	"fmt"
	"time"

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

func ValidaUsuario(user string, pass string) Empleado {

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
	}
	defer db.Close()
	fmt.Printf("el usuario es 2" + user + pass)
	filas, err := db.Query("SELECT idEmpleados,Nombre,Apellidos,Correo,Admin FROM Empleados where Correo= '" + user + "' and Contrasena=SHA('" + pass + "')")

	if err != nil {
		fmt.Println("error en la consulta")

	}
	defer filas.Close()

	var c Empleado

	for filas.Next() {

		err = filas.Scan(&c.Id, &c.Nombre, &c.Apellidos, &c.Correo, &c.Admin)
		if err != nil {
			fmt.Println("error al scanear")

		}
		fmt.Println("poso por sql")
		fmt.Println(c)
	}
	return c
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

func RegistraEmpleado(user Empleado) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare("INSERT INTO VtgHpFxzCP.Empleados (Nombre, Apellidos, Telefono, Correo, Contrasena, Admin) VALUES(?, ?, ?, ?, SHA(?),?)")
	if err != nil {
		fmt.Println("hubo un error en la incercion")
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(user.Nombre, user.Apellidos, user.Telefono, user.Correo, user.Contrasena, "Empleado")
	if err != nil {
		return false
	}
	return true
}

func SeleccionaEmpleados() []Empleado {
	empleados := []Empleado{}

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")

	}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	defer db.Close()
	filas, err := db.Query("SELECT idEmpleados,Nombre, Apellidos,Telefono,Correo, IFNULL(FechaEntrada, 'Sin Checar') as FechaEntrada, IFNULL(FechaSalida,'Sin Checar') as FechaSalida FROM Empleados left join Registro on idEmpleados=Empleados_idEmpleados and DATE(FechaEntrada) BETWEEN '" + fecha + "'	AND '" + fecha + "' order by idEmpleados")

	if err != nil {
		fmt.Println("error en la consulta")

	}
	defer filas.Close()

	var c Empleado

	for filas.Next() {

		err = filas.Scan(&c.Id, &c.Nombre, &c.Apellidos, &c.Telefono, &c.Correo, &c.FechaInicio, &c.FechaFin)
		if err != nil {
			fmt.Println("error al consulta empleados scanear")

		}

		empleados = append(empleados, c)
	}
	return empleados
}

func SeleccionaEmpleado(id string) Empleado {

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")

	}
	defer db.Close()
	filas, err := db.Query("SELECT idEmpleados,Nombre,Apellidos,Telefono,Correo,Contrasena FROM Empleados where idEmpleados = " + id)

	if err != nil {
		fmt.Println("error en la consulta")

	}
	defer filas.Close()

	var c Empleado

	for filas.Next() {

		err = filas.Scan(&c.Id, &c.Nombre, &c.Apellidos, &c.Telefono, &c.Correo, &c.Contrasena)
		if err != nil {
			fmt.Println("error al scanear")

		}
		fmt.Println("poso por sql")
		fmt.Println(c)
	}
	return c
}

func ActulizaEmpleado(user Empleado) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare("UPDATE Empleados SET Nombre = ?, Apellidos= ?, Telefono= ?, Correo= ?, Contrasena=SHA(?) WHERE idEmpleados=?")
	if err != nil {
		fmt.Println("hubo un error en la Actulizacion")
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(user.Nombre, user.Apellidos, user.Telefono, user.Correo, user.Contrasena, user.Id)
	if err != nil {
		return false
	}
	return true
}

func ActulizaEmpleadoSinContrasena(user Empleado) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	sentenciaPreparada, err := db.Prepare("UPDATE Empleados SET Nombre = ?, Apellidos= ?, Telefono= ?, Correo= ? WHERE idEmpleados=?")
	if err != nil {
		fmt.Println("hubo un error en la Actulizacion")
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(user.Nombre, user.Apellidos, user.Telefono, user.Correo, user.Id)
	if err != nil {
		return false
	}
	return true
}

func BorraEmpleado(id string) {
	db, err := NewConeccion()
	fmt.Println("borrara con el id :" + id)
	if err != nil {
		fmt.Println("hubo un error")

	}
	defer db.Close()
	filas, err := db.Exec("DELETE FROM Empleados WHERE idEmpleados = " + id)

	if err != nil {
		fmt.Println("error al borrar")

	} else {
		fmt.Println(filas)
	}
}

//INSERT INTO Registro (FechaEntrada, Empleados_idEmpleados) VALUES ('2020-04-05 09:30:00', '3');

func RegistroEntrada(user Empleado) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	sentenciaPreparada, err := db.Prepare("INSERT INTO Registro (FechaEntrada, Empleados_idEmpleados) VALUES (?, ?)")
	if err != nil {
		fmt.Println("hubo un error en la incercion")
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(fecha, user.Id)
	if err != nil {
		return false
	}
	return true
}

func RegistroSalida(user Empleado, salida string) bool {
	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")
		return false
	}
	defer db.Close()
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	sentenciaPreparada, err := db.Prepare("UPDATE Registro set FechaSalida = ?, Empleados_idEmpleados = ? where idRegistro=?")
	if err != nil {
		fmt.Println("hubo un error en la incercion")
		return false
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(fecha, user.Id, salida)
	if err != nil {
		return false
	}
	return true
}

func EntradaRegistrada(id string) string {

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")

	}
	defer db.Close()
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	filas, err := db.Query("SELECT idRegistro FROM Registro WHERE Empleados_idEmpleados=" + id + " and DATE(FechaEntrada) BETWEEN '" + fecha + "'	AND '" + fecha + "'")

	if err != nil {
		fmt.Println("error en la consulta")

	}
	defer filas.Close()

	var c string

	for filas.Next() {

		err = filas.Scan(&c)
		if err != nil {
			fmt.Println("error al scanear")

		}
		fmt.Println("poso por sql")
		fmt.Println(c)
	}
	return c
}

func TablaSemana() []Empleado {
	empleados := []Empleado{}

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")

	}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	defer db.Close()
	filas, err := db.Query("SELECT idEmpleados,Nombre, Apellidos,Telefono,Correo, IFNULL(FechaEntrada, 'Sin Checar') as FechaEntrada, IFNULL(FechaSalida,'Sin Checar') as FechaSalida FROM Empleados left join Registro on idEmpleados=Empleados_idEmpleados and DATE(FechaEntrada) BETWEEN '" + fecha + "'	AND '" + fecha + "' order by idEmpleados")
	if err != nil {
		fmt.Println("error en la consulta")
	}

	defer filas.Close()

	var c Empleado

	for filas.Next() {

		err = filas.Scan(&c.Id, &c.Nombre, &c.Apellidos, &c.Telefono, &c.Correo, &c.FechaInicio, &c.FechaFin)
		if err != nil {
			fmt.Println("error al consulta empleados scanear")
		}
		empleados = append(empleados, c)
	}
	return empleados
}

func CambioCantrasena(id string) string {

	db, err := NewConeccion()
	if err != nil {
		fmt.Println("hubo un error")

	}
	defer db.Close()

	filas, err := db.Query("SELECT Contrasena FROM Empleados where idEmpleados=" + id)

	if err != nil {
		fmt.Println("error en la consulta")

	}
	defer filas.Close()

	var c string

	for filas.Next() {

		err = filas.Scan(&c)
		if err != nil {
			fmt.Println("error al scanear")

		}
		fmt.Println("poso por sql")
		fmt.Println(c)
	}
	return c
}
