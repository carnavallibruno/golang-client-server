package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// connect to server
	conn, _ := net.Dial("tcp", ":8000")
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Nome do Aluno: ")
	// name, _ := reader.ReadString('\n')

	for {
		// what to send?
		reader := bufio.NewReader(os.Stdin)
		for i := 0; i < 3; i++ {

			fmt.Print("nota prova ", i+1, ": ")
			text, _ := reader.ReadString('\n')
      // send to server
      fmt.Fprintf(conn, text+"\n")
		}
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
