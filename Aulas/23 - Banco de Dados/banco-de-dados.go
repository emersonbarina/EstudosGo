package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "root:root@tcp(0.0.0.0:3307)/devbook?charset=utf8&parseTime=True&loc=Local"
	//stringConexao := "golang:golang@tcp(0.0.0.0:3307)/devbook?charset=utf8&parseTime=True&loc=Local"
	//db, err:= sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/mydbname?parseTime=true")

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		// fmt.Println("Dentro do sql.Open")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		// fmt.Println("Dentro do db.Ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão aberta")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
