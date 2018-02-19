package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"encoding/json"
	
)

type Pessoa struct {
	Nome string
	Curso string
	Cidade string
}


func main() {

	conn, err1 := net.Dial("tcp", "localhost:9000")
	if err1 != nil {
		panic(err1)
	}
	defer conn.Close()
	
	for {

		pessoa := Pessoa{}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Nome: ")
		pessoa.Nome, _ = reader.ReadString('\n')
		fmt.Print("Curso: ")
		pessoa.Curso, _ = reader.ReadString('\n')
		fmt.Print("Cidade: ")
		pessoa.Cidade, _ = reader.ReadString('\n')
		
		j, _ := json.Marshal(pessoa)
		
		conn.Write([]byte(j))

		// ouvindo a resposta do servidor (eco)
		mensagem, err2 := bufio.NewReader(conn).ReadString('\n')
		if err2 != nil {
			fmt.Println(err2)
			os.Exit(3)
		}
		// escrevendo a resposta do servidor no terminal
		fmt.Print("Resposta do servidor: " + mensagem)
		
	}
		
}