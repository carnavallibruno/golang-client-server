package main

import (
	"bufio"
	"fmt"
	"net"
	// "strconv"
)

func main() {
  fmt.Println("Start server...")
  
  // listen on port 8000
  ln, _ := net.Listen("tcp", ":8000")
  
  // accept connection
  conn, _ := ln.Accept()
  
  // run loop forever (or until ctrl-c)
  for {

    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Nome do aluno: ", string(message))

    message1, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Nota 1 do aluno: ", string(message1))

    message2, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Nota 2 do aluno: ", message2)

    message3, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Println("Nota 3 do aluno: ", message3)
  }
}