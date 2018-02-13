package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"bufio"
	"os"
	"strconv"
)

func main() {
	conn, err1 := net.Dial("tcp", "localhost:9000")
	if err1 != nil {
		panic(err1)
	}
	defer conn.Close()

	for {
		// lendo entrada do terminal
		fmt.Print("Informe o primeiro número: ")
		var x float64
		fmt.Scanf("%f\n", &x)

		fmt.Print("Informe o segundo número: ")
		var y float64
		fmt.Scanf("%f\n", &y)

		// escrevendo a mensagem na conexão (socket)
		n1 := strconv.FormatFloat(x, 'f', 2, 64)
		n2 := strconv.FormatFloat(y, 'f', 2, 64)
		fmt.Fprintf(conn, n1 + "," + n2 + "\n")

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