package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func CreateNumericTable() string {
	fmt.Println("Ingrese numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
	}

	return texto
}
