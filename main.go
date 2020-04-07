package main

/*
type cal struct{}

func (cal) hello() {
	fmt.Println("Hello word")
}

*/

func main() {

	server := NewServer(":3000")

	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/login", CheckAut)
	server.Handle("POST", "/registro", Registro)

	server.Handle("GET", "/Empleados", GetEmpleados)
	server.Handle("POST", "/Empleado", RegistroEmpleado)
	server.Handle("POST", "/borrarEmpleado", BorrarEmpleado)
	server.Handle("POST", "/getEmpleado", GetEmpleado)
	server.Handle("POST", "/registroEntrada", RegistraEntrada)
	server.Handle("POST", "/check", checo)
	server.Listen()
	/*
		fmt.Println("hola mundo")
		nu := cal{}
		nu.hello()

	*/

}
