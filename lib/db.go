package lib

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var confuguracao = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "123g",
	Database: "membros",
}

//Sess que é uma variavel que faxz conexão com o banco de dados
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(confuguracao)
	if err != nil {
		log.Fatal(err.Error())
	}
}
