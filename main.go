package main

import (
	"go-crud-mongo/database"
	"go-crud-mongo/helper"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(`.env`)
	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()
	helper.InitLogger()
	database.ConnectMongo()

	startEchoServer()
}
