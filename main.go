package main

/*
type cal struct{}

func (cal) hello() {
	fmt.Println("Hello word")
}

*/

func main() {

	server := NewServer(":3000")
	server.Handle("POST", "/login", CheckAut)
	server.Handle("GET", "/Empleados", GetEmpleados)
	server.Handle("POST", "/Empleado", RegistroEmpleado)
	server.Handle("POST", "/borrarEmpleado", BorrarEmpleado)
	server.Handle("POST", "/getEmpleado", GetEmpleado)
	server.Handle("POST", "/registroEntrada", RegistraEntrada)
	server.Handle("POST", "/check", checo)
	server.Listen()

}
