package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(name string, canal1 chan bool) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(time.Second)
		fmt.Println(letra)
	}
	canal1 <- true
}
