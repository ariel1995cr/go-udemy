package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 0, 55}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[3] = 54

	tabla2 := [10]string{"ariel", "aldana", "bautista"}
	fmt.Println(tabla)
	fmt.Println(tabla2)
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[6][24] = 15
	fmt.Println(matriz)
}
