package main

import (
	"fmt"

	"github.com/arielt1995/godesde0/variables"
)

func main() {
	variables.MostrarEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConvertToText(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
