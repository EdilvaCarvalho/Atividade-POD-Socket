package main

import (
    "database/sql"
    "fmt"
    _ "postgresql/pq"
    _ "mysql/mysql"
    //"time"
    //_ "github.com/lib/pq"
    "bufio"
    "net"
    "log"
    "encoding/json"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME     = "pod"
)

type Pessoa struct {
	Nome string
	Curso string
	Cidade string
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
        ln := scanner.Text()

        fmt.Println(ln)

        var p Pessoa
		json.Unmarshal([]byte(ln), &p)

		//pq
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Inserindo valores")

        var lastInsertId int
        err = db.QueryRow("INSERT INTO pessoa(nome, curso, cidade) VALUES($1, $2, $3) returning uid;", p.Nome, p.Curso, p.Cidade).Scan(&lastInsertId)
        checkErr(err)
        fmt.Println("Postgres: Último id inserido =", lastInsertId)


        //mysql

        db2, err := sql.Open("mysql", "root:@/pod?charset=utf8")
        checkErr(err)

        // insert
        stmt, err := db2.Prepare("INSERT pessoa SET nome=?, curso=?, cidade=?")
        checkErr(err)

        res, err := stmt.Exec(p.Nome, p.Curso, p.Cidade)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println("Mysql: Último id inserido =", id)

        // envia a resposta para o cliente
		conn.Write([]byte("Pessoa armazenada com sucesso!" + "\n"))
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
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