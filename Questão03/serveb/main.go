package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		//var r interface{}
		//r = recover()
		if r := recover(); r != nil {
			fmt.Println(r, "recover")
		}
	}()

	s, err := net.Listen("tcp", ":9091")

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	for {
		conn, er := s.Accept()

		if er != nil {
			log.Println(er)
			continue
		}

		go func(c net.Conn) {
			b := make([]byte, 100)

			fmt.Println("iniciando conexão com", c.RemoteAddr().String())

			c.Read(b)
			var resl float64
			men := string(b)
			var op, n1, n2 = unmachilig(men)
			if op == "op1" {
				resl = op1(n1, n2)
				men = (strconv.FormatFloat(resl, 'f', 2, 64) + "\n")
			} else if op == "op2" {
				go func() {
					men = forward_(men)
				}()

			}
			fmt.Println(men)
			conn.Write([]byte(men))
			c.Close()
		}(conn)
	}

}

func unmachilig(men string) (op string, n1 float64, n2 float64) {
	tipo := strings.Split(men, ";")
	op = tipo[0]
	n1, _ = strconv.ParseFloat(tipo[1], 64)
	n2, _ = strconv.ParseFloat(tipo[2], 64)

	fmt.Println(op)
	fmt.Println(n1)
	fmt.Println(n2)
	return
}

func op1(n1 float64, n2 float64) (resl float64) {
	resl = n1 + n2
	return
}

func forward_(men string) (resl string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "recover")
		}
	}()
	s, err := net.Dial("tcp", ":9092")
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()
	s.Write([]byte(men))
	b := make([]byte, 15)
	s.Read(b)
	resl = string(b)
	return
}
