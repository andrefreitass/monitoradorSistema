package menu

import "fmt"

func LerComando() int {

	var comandoLido int

	fmt.Scan(&comandoLido)
	fmt.Println("O comando selecionado foi: ", comandoLido)

	return comandoLido
}
