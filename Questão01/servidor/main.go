package main

import (
	"net"
	"fmt"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		var str []string = strings.Split(ln, ",")
		fmt.Println("Números recebidos: " + str[0] + " e " + str[1])

		n1, _ := strconv.ParseFloat(str[0], 64)
		n2, _ := strconv.ParseFloat(str[1], 64)

		n1 = n1 * 100
		n2 = n2 * 100

		soma := 0.0

		for num := n2; num > 0; num-- {
			soma = soma + n1
		}

		result := soma * 0.0001

		// envia a resposta para o cliente
		conn.Write([]byte(strconv.FormatFloat(result, 'f', 2, 64) + "\n"))
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}