package main

import (
	"github.com/kawamurasorachi/Gin-Postgres/db"
	"github.com/kawamurasorachi/Gin-Postgres/server"
)

func main() {
	db.Init()
	server.Init()
}
