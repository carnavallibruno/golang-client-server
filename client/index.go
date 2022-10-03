package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	for {
		// mensagem para o servidor
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < 3; i++ {
			fmt.Print("Nome do aluno ", i+1, ": ")
			text, _ := reader.ReadString('\n')

			// send to server
			fmt.Fprintf(conn, text)

			for i := 0; i < 3; i++ {
				fmt.Println("Nota", i+1, "do aluno: ")

				text, _ := reader.ReadString('\n')

				fmt.Fprintf(conn, text)
			}
		}
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
