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
	server.Handle("POST", "/valida", ValidaToken)
	server.Handle("POST", "/valida", server.AddMiddleware(prueba, CheckAuth()))
	server.Handle("GET", "/Empleados", server.AddMiddleware(GetEmpleados, CheckAuth()))
	server.Handle("POST", "/Empleado", server.AddMiddleware(RegistroEmpleado, CheckAuth()))
	server.Handle("POST", "/borrarEmpleado", server.AddMiddleware(BorrarEmpleado, CheckAuth()))
	server.Handle("POST", "/getEmpleado", server.AddMiddleware(GetEmpleado, CheckAuth()))
	server.Handle("POST", "/registroEntrada", server.AddMiddleware(RegistraEntrada, CheckAuth()))
	server.Handle("POST", "/check", server.AddMiddleware(checo, CheckAuth()))
	server.Listen()

}
