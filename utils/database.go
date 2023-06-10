package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 6543
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func Getconnection() *sqlx.DB {

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err, "-:🔴 ERROR OCCUERED IN CONNECTING DATABASE 🔴 ")
		panic(err)
	}
	log.Println("💚 DATABSE CONNECTION ESTABLISHED 💚")
	return db

}
