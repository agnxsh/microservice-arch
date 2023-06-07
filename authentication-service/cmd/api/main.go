package main

import (
	"authentication/data"
	"database/sql"
	"log"
)

const webPort = "80"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main(){
	log.Println("Starting authetication service...")

	 
}