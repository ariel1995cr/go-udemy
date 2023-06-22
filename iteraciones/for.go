package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 100; i++ {
		if i == 5 {
			//continue
			break
		}
		fmt.Println(i)
	}
}
