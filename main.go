package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comando int
	fmt.Print("\nEscolha um comando: ")
	fmt.Scanf("%d", &comando)

	fmt.Printf("Você escolheu o comando %d\n", comando)
	return comando
}

func iniciarMonitoramento(site string) {
	fmt.Println("Monitorando...")

	for i := 0; i < 10; i++ {
		res, err := http.Get(site)
		if err != nil {
			fmt.Println("Erro ao fazer requisição:", err)
			return
		}

		if res.StatusCode == 200 {
			fmt.Println("Site:", site, "está no ar! ✅")
		} else {
			fmt.Println("ERRO: Site", site, "retornou status %d\n", res.StatusCode)
		}

		time.Sleep(10 * time.Second)
	}
}

func main() {
	fmt.Println("=======================================")
	fmt.Println("Bem-vindo ao programa de monitoramento!")
	fmt.Println("=======================================")

	var site string
	fmt.Print("\nDigite a URL do seu site: ")
	fmt.Scanf("%s", &site)

	for {
		fmt.Println()
		exibeMenu()

		comando := leComando()

		if comando == 1 {
			iniciarMonitoramento(site)
		} else if comando == 0 {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("Comando inválido")
			os.Exit(-1)
		}
	}
}
