package main

import (
	"github.com/gin-gonic/gin"
	"tutoringsystem/config"
	"tutoringsystem/routes"
	"tutoringsystem/utils"
)

func main() {
	config.LoadEnv()
	utils.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	_ = r.Run(":8080")
}
