package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arielt1995/godesde0/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func SaveTableInFile() {
	var texto string = ejercicios.CreateNumericTable()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo ", err.Error())
		return
	}
	fmt.Fprintln(file, texto)
	file.Close()
}

func SumTable() {
	var texto string = ejercicios.CreateNumericTable()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(fileName string, text string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(text)
	if err != nil {
		fmt.Println("Error al escribir" + err.Error())
		return false
	}

	arch.Close()

	return true
}

func LeoArchivo() {
	/* file, err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println("Erro al leer el archivo " + err.Error())
		return
	}

	fmt.Println(string(file)) */

	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro al leer el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
