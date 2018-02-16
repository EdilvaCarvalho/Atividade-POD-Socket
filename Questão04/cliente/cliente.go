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
	
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		
		conn.Write([]byte(text))
		
	}
		
}