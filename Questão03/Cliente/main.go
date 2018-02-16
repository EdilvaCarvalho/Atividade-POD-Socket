package main

import (
	"fmt"
	"net"
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
		var x int
		fmt.Scanf("%f\n", &x)

		fmt.Print("Informe o segundo número: ")
		var y int
		fmt.Scanf("%f\n", &y)

		// escrevendo a mensagem na conexão (socket)
		n1 := strconv.FormatInt(x, 'f', 2, 64)
		n2 := strconv.FormatFloat(y, 'f', 2, 64)
		fmt.Fprintf(conn, n1 + "," + n2 + "\n")
	}
}