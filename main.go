package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kchernenko/eventssample/db"
	"github.com/kchernenko/eventssample/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRotes(server)

	server.Run(":8080")
}
