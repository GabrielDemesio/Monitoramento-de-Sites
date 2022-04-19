package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Laço de repetição for, quando  passado sem  parâmetro fica num loop
	for {
		menu()
		comando := comandoLido()

		/* Condicional switch faz um if por baixo dos panos*/
		switch comando {
		case 0:
			// Executo a função
			iniciandoMonitoramento()
		case 1:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 2:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Desconhecido!")
			os.Exit(-1)
		}
	}
}
func menu() {
	fmt.Println("0 - Iniciar Monitoramento")
	fmt.Println("1 - Exibir Logs")
	fmt.Println("2 - Sair do Programa")
}
func comandoLido() int {
	/* Guardar a variável comando num int */
	var comandoLido int
	/* %d usado para int, &comercial usando no comando para chamar a função */
	fmt.Scanf("%d", &comandoLido)
	/* Condicionais if, else, elif, e switch*/
	return comandoLido
}
func iniciandoMonitoramento() {
	fmt.Println("Iniciando Monitoramento...")
	sites := leArquivo()
	// Range informa a posição permitindo usar no for
	for i, site := range sites {
		fmt.Println("Estou passando na posição ", i, ":", site)
		testaSite(site)
	}
	fmt.Println("")
}
func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("O site: ", site, "funcionando corretamente!, status", resp.StatusCode)
		registraLogs(site, true)
	} else {
		fmt.Println("O site", site, "Nao está funcionando. StatusCode", resp.StatusCode)
		registraLogs(site, false)
	}
}
func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		fmt.Println(linha)
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
func registraLogs(site string, status bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(arquivo)
	arquivo.Close()
}
func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))

}
