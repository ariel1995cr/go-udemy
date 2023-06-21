package ejercicios

import (
	"strconv"
)

func ReturnValue(numero string) (int, string) {
	var numeroInt int
	var message string
	if numeroInt, error := strconv.Atoi(numero); error != nil {
		return 0, "Hubo un error" + error.Error()
	} else {
		if numeroInt > 100 {
			message = "Es mayor a 100"
		} else {
			message = "Es menor a 100"
		}
	}

	return numeroInt, message
}
