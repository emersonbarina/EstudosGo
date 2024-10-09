package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@tcp(docker.for.mac.localhost:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	//db, err:= sql.Open("mysql", "root:password@tcp(docker.for.mac.localhost:3306)/mydbname?parseTime=true")

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão aberta")

}
