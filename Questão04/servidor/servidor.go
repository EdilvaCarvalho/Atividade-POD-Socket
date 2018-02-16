package main

import (
    "database/sql"
    "fmt"
    //"postgresql/pq"
    //"mysql/mysql"
    //"time"
    _ "github.com/lib/pq"
    "bufio"
    "net"
    "log"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "postgres"
    DB_NAME     = "test"
)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()

		//pq
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        checkErr(err)
        defer db.Close()

        fmt.Println("# Inserting values")

        var lastInsertId int
        err = db.QueryRow("INSERT INTO message(mensagem) VALUES($1) returning uid;", ln).Scan(&lastInsertId)
        checkErr(err)
        fmt.Println("Postgres: last inserted id =", lastInsertId)


        //mysql

        db2, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
        checkErr(err)

        // insert
        stmt, err := db2.Prepare("INSERT message SET mensagem=?")
        checkErr(err)

        res, err := stmt.Exec(ln)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)

        fmt.Println("Mysql: last inserted id =", id)
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