package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
  "math"
)

func round(num float64) int {
  return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
  output := math.Pow(10, float64(precision))
  return float64(round(num * output)) / output
}

func main() {
  fmt.Println("Start server...")
  
  // listen on port 8000
  ln, _ := net.Listen("tcp", ":8000")
  
  // accept connection
  conn, _ := ln.Accept()
  
  // run loop forever (or until ctrl-c)
  for {

    nomeAluno, _ := bufio.NewReader(conn).ReadString('\n')
    nomeString := (strings.TrimSpace(nomeAluno))


    nota1, _ := bufio.NewReader(conn).ReadString('\n')
    nota1string, _ := strconv.ParseFloat(strings.TrimSpace(nota1), 64)


    nota2, _ := bufio.NewReader(conn).ReadString('\n')
    nota2string, _ := strconv.ParseFloat(strings.TrimSpace(nota2), 64)


    nota3, _ := bufio.NewReader(conn).ReadString('\n')
    nota3string, _ := strconv.ParseFloat(strings.TrimSpace(nota3), 64)
    
    var media = (nota1string + nota2string + nota3string) / 3
    fmt.Println("----------------------------\n\nNome: ", nomeString)
    fmt.Println("\nNotas: ",nota1string,"|",nota2string,"|",nota3string)
    fmt.Println("\nMédia: ", toFixed(media, 2),"\n\n----------------------------")
  }
}