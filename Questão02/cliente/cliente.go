package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {

	conn, err1 := net.Dial("tcp", "localhost:9000")
	if err1 != nil {
		panic(err1)
	}
	defer conn.Close()

	go readConnection(conn)
	
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		if len(text) >= 10 && len(text) <= 100 {
			conn.Write([]byte(text))
		} else {
			fmt.Println("A mensagem deve conter no mínimo 10 e no máximo 100 caracteres.")
		}
		
	}
		
}

func readConnection(conn net.Conn) {
	for {
		// ouvindo a resposta do servidor (eco)
		mensagem2, err3 := bufio.NewReader(conn).ReadString('\n')
		if err3 != nil {
			fmt.Println(err3)
			os.Exit(3)
		}
		// escrevendo a resposta do servidor no terminal
		fmt.Print(mensagem2)
	}
}