package main

import (
	"github.com/arielt1995/godesde0/middleware"
)

func main() {
	/* variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvertToText(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "darwin" {
		fmt.Println("Esto no es windows, es", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es LINUX")
	case "darwin":
		fmt.Println("Esto es DARWIN")
	default:
		fmt.Printf("%s \n", os)
	}
	var entero int
	var mensaje string
	entero, mensaje = ejercicios.ReturnValue("100")
	fmt.Printf("El entero es %d y el mensaje es %s \n", entero, mensaje)
	entero, mensaje = ejercicios.ReturnValue("120")
	fmt.Printf("El entero es %d y el mensaje es %s \n", entero, mensaje)

	entero, mensaje = ejercicios.ReturnValue("asddsa")
	fmt.Printf("El entero es %d y el mensaje es %s \n", entero, mensaje)

	teclado.IngresoNumeros()

	fmt.Println(ejercicios.CreateNumericTable())

	funciones.Calculos()
	funciones.LlamarClosure()
	funciones.Exponencia(2)
	arreglos_slices.MuestroArreglos()
	arreglos_slices.MuestroSlice()
	arreglos_slices.Capacidad()
	mapas.MostrarMapas()
	users.AltaUsuario()

	Ariel := new(modelos.Hombre)
	e.HumanosRespirando(Ariel)

	Aldana := new(modelos.Mujer)
	e.HumanosRespirando(Aldana)

	//d.VemosDefer()
	//d.EjemploPanic()
	canal1 := make(chan bool)
	go gorutines.MiNombreLento("Ariel Triviño", canal1)
	defer func() {
		<-canal1
	}()
	fmt.Println("Estoy aqui")*/

	//webserver.MiWebServer()
	middleware.MiMiddleware()
}
