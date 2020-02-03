package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"	
	"strings"
	"time"
)

func opcoes() {
	fmt.Println("[1]-Iniciar Monitoramento")	
	fmt.Println("[0]-Sair do Progrma")
}

func main() {
	for {
		opcoes()
		comando := funcaoParam()

		switch comando {
		case 1:
			mostraMonitoramento()		
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço essa instrução")	
			main()		
		}
	}
}

func funcaoParam() int {
	var opcao int
	fmt.Scan(&opcao)
	return opcao
}

func mostraMonitoramento() {
	const tempo = 4

	fmt.Println("Monitorando...")

	var sites []string

	sites = leArquivo()	

	for i := 0; i < 3; i++ {

		for x := 0; x <= len(sites)-1; x++ {
			testaSite(sites[x])
		}

		time.Sleep(tempo * time.Second)
	}
}

func testaSite(site string) {

	resposta, _ := http.Get(site)
	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")		
	} else {
		fmt.Println("Site:", site, "não pode ser carregado")		
	}

}

func leArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: O arquivo sites.txt não foi encontrado ou está corrompido")
		main()
	}

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