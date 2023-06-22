package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func CreateNumericTable() {
	fmt.Println("Ingrese numero: ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			break
		}
	}

	for i := 0; i <= 10; i++ {
		var res int = numero * i
		fmt.Printf("%d * %d = %d \n", numero, i, res)
	}

}
