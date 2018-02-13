package main

import (
	"net"
	"fmt"
	"bufio"
	"log"
)

func main() {
	ln, err1 := net.Listen("tcp", ":9000")
	if err1 != nil {
		log.Fatalln(err1)
	}
	defer ln.Close()

	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan string)
	i := 0

	go func() {
		for {
			conn, err2 := ln.Accept()
			if err2 != nil {
				log.Fatalln(err2)
			}
			defer conn.Close()

			conns <- conn
		}
	}()

	for {
		select {
		
		case conn := <- conns:
			aconns[conn] = i
			i++

			go func(conn net.Conn, i int) {
				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						break
					}
					msgs <- fmt.Sprintf("Cliente %v: %v", i, m)
				}

				dconns <- conn
			}(conn, i)

		case msg := <- msgs:
			for conn := range aconns {
				conn.Write([]byte(msg))
			}

		case dconn := <- dconns:
			log.Printf("Cliente %v saiu\n", aconns[dconn]+1)
			delete(aconns, dconn)
		}
	}

}