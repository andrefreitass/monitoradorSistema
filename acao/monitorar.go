package acao

import (
	"bufio"
	"fmt"
	"io"
	constant "monitoradorSistema/util"
	error "monitoradorSistema/util"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func IniciarMonitoramento() {

	fmt.Println("Monitorando...")

	sites := lerSiteArquivo()

	for i := 0; i < constant.Monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(constant.Delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func lerSiteArquivo() []string {
	var sites []string
	arquivo, err := os.Open("util/sites.txt")

	error.ValidaError(err)

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	resp, err := http.Get(site)
	fmt.Println("")

	error.ValidaError(err)

	validaStatusCode(resp.StatusCode, site)
}

func validaStatusCode(statusCode int, site string) {
	if statusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("O site,", site, "está com problema, Status Code", statusCode)
		registraLog(site, false)
	}
}

func registraLog(site string, statusSite bool) {

	arquivo, err := os.OpenFile("logs/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	error.ValidaError(err)

	arquivo.WriteString(time.Now().Format("02-01-2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(statusSite) + "\n")

	arquivo.Close()
}
