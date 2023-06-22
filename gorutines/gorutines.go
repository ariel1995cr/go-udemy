package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(time.Second)
		fmt.Println(letra)
	}
}
