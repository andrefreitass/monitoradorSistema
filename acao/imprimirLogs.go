package acao

import (
	"fmt"
	"io/ioutil"
	error "monitoradorSistema/util"
)

func ImprimiLogs() {

	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("logs/log.txt")

	error.ValidaError(err)

	fmt.Println(string(arquivo))
}
