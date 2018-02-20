package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
)

func main() {
	defer func() {
		//var r interface{}
		//r = recover()
		if r := recover(); r != nil {
			fmt.Println(r, "recover")
		}
	}()

	fmt.Print("Informe o primeiro número: ")
	var x float64
	fmt.Scanf("%f\n", &x)

	fmt.Print("Informe o segundo número: ")
	var y float64
	fmt.Scanf("%f\n", &y)

	// escrevendo a mensagem na conexão (socket)
	var men1 = machiling("op1", x, y)
	var men2 = machiling("op2", x, y)
	fmt.Println(men1)
	fmt.Println(men2)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {

		s, err := net.Dial("tcp", ":9091")
		if err != nil {
			log.Fatal(err)
		}
		s.Write([]byte(men1))
		b := make([]byte, 15)
		s.Read(b)
		defer s.Close()
		fmt.Println(string(b))
		fmt.Println("servidoB men1")
		wg.Done()
	}()
	go func() {

		s, err := net.Dial("tcp", ":9091")
		if err != nil {
			log.Fatal(err)
		}
		s.Write([]byte(men2))
		b := make([]byte, 15)
		s.Read(b)
		defer s.Close()
		fmt.Println(string(b))
		fmt.Println("servidoB men2")
		wg.Done()
	}()
	go func() {
		s, err := net.Dial("tcp", ":9092")
		if err != nil {
			log.Fatal(err)
		}

		s.Write([]byte(men2))
		b := make([]byte, 15)
		s.Read(b)
		fmt.Println(string(b))
		defer s.Close()
		fmt.Println("servidoC men2")
		wg.Done()
	}()
	go func() {
		s, err := net.Dial("tcp", ":9093")
		if err != nil {
			log.Fatal(err)
		}
		defer s.Close()
		s.Write([]byte(men1))
		b := make([]byte, 15)
		s.Read(b)
		fmt.Println(string(b))
		fmt.Println("servidoD men1")
		wg.Done()
	}()
	wg.Wait()

}

func machiling(op string, n1 float64, n2 float64) (men string) {
	men = op + ";" + strconv.FormatFloat(n1, 'f', 2, 64) + ";" + strconv.FormatFloat(n2, 'f', 2, 64) + ";"
	return
}
