package menu

import (
	"fmt"
	acao "monitoradorSistema/acao"
	"os"
)

func AcessaMenu() {

	exibeIntroducao()

	for {

		exibeOpcoes()

		opção := LerComando()

		acessandoMenu(opção)
	}

}

func exibeIntroducao() {

	nome := "Monitorador"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeOpcoes() {

	fmt.Println("1 - Iniciar Monitoramente")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func acessandoMenu(comandoRecebido int) {
	switch comandoRecebido {
	case 1:
		acao.IniciarMonitoramento()
	case 2:
		acao.ImprimiLogs()
	case 0:
		acao.SairPrograma()
	default:
		fmt.Println("Informe uma opção valida!")
		os.Exit(-1)
	}
}
