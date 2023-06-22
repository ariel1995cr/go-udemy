package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 37,
		"Chivas":      30,
		"Boke Jrs":    30,
	}

	fmt.Println(campeonato)

	delete(campeonato, "Real Madrid")

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s tiene un puntaje %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Boke Jrs"]
	fmt.Printf("El puntaje capturado es %d y el equipo es existe = %t \n", puntaje, existe)
}
