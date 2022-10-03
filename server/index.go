package main

import (
	"bufio"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

// função para arrendondar caracteres do tipo number
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// função para fixer um numero de caracteres para um number
func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	fmt.Println("Abrindo servidor...")
	//declara a porta
	args := os.Args
	fmt.Print(args)
	//retorna erro caso não tenha sido declarada nenhuma porta
	if len(args) == 1 {
		fmt.Println("Por favor digite a porta que deseja abrir o servidor")
		return
	}
	// escuta na porta passada via terminal
	ln, err := net.Listen("tcp", args[1])
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("TCP Server initialized at port", args[1])
	// aceita conexão na porta seleciona
	conn, err := ln.Accept()
	if err != nil {
		fmt.Print(err)
	}
	//// ao final fecha a conexão
	defer ln.Close()
	// se tudo ocorreu corretamente servidor é aberto
	fmt.Print("Cliente conectado")

	// loop eterno
	for {
		//prepara o servidor apra ler as mensagens do client
		mensagem, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.ToUpper(strings.TrimSpace(string(mensagem))) == "SAIR" {
			fmt.Println("servidor encerrado.")
			return
		}
		// cria as variaveis com o valor dado pelo cliente
		nomeAluno, _ := bufio.NewReader(conn).ReadString('\n')
		nomeString := (strings.TrimSpace(nomeAluno))

		nota1, _ := bufio.NewReader(conn).ReadString('\n')
		nota1string, _ := strconv.ParseFloat(strings.TrimSpace(nota1), 64)

		nota2, _ := bufio.NewReader(conn).ReadString('\n')
		nota2string, _ := strconv.ParseFloat(strings.TrimSpace(nota2), 64)

		nota3, _ := bufio.NewReader(conn).ReadString('\n')
		nota3string, _ := strconv.ParseFloat(strings.TrimSpace(nota3), 64)
		//imprime a mesangem no servidor
		var media = (nota1string + nota2string + nota3string) / 3

		mediaCalculada := toFixed(media, 2)
		mediaCalculadaString := strconv.FormatFloat(mediaCalculada, 'E', -1, 32)

		var message = "----------------------------\n\nNome: " + nomeString + "\nNotas: " + nota1 + "|" + nota2 + "|" + nota3 + "\nMédia: " + mediaCalculadaString + "\n\n----------------------------"


		//imprime a mensagem no cliente
		conn.Write([]byte(message))
	}
}
