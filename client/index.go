package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	fmt.Print(args)
	//retorna erro caso não tenha sido declarada nenhuma porta
	if len(args) == 1 {
		fmt.Println("Por favor digite a porta que deseja se comunicar com o servidor")
		return
	}
	// se conecta ao servidor pela porta declarada
	conn, _ := net.Dial("tcp", args[1])
	defer conn.Close()
	fmt.Println("To exit, digit EXIT and press key emter.")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Comandos: sair ou resgistrar\n")
		comando, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.ToUpper(strings.TrimSpace(string(comando))) == "SAIR" {
			fmt.Println("Saindo do servidor...")
			return
		}
		if strings.ToUpper(strings.TrimSpace(string(comando))) == "REGISTRAR" {
			fmt.Print("Quantos alunos deseja registrar?\n")
			quantidade, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			quantidadeNumber, _ := strconv.Atoi(strings.TrimSpace(quantidade))

			for i := 0; i < quantidadeNumber; i++ {
				fmt.Print("Nome do aluno ", i+1, ": ")
				text, _ := reader.ReadString('\n')

				// mensagem para o servidor

				fmt.Fprintf(conn, text)

				for i := 0; i < 3; i++ {
					fmt.Println("Nota", i+1, "do aluno: ")

					text, _ := reader.ReadString('\n')

					fmt.Fprintf(conn, text)
				}
				message, _ := bufio.NewReader(conn).ReadString('\n') // Aguarda resposta do servidor
				fmt.Print("RCV: " + message)
			}
		}
	}

}
