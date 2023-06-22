package middleware

import "fmt"

func add(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio Middleware")
	result := operacionesMidd(add)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(substract)(20, 3)
	fmt.Println(result)
	result = operacionesMidd(multiply)(2, 4)
	fmt.Println(result)
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operacion")
		return f(a, b)
	}
}
