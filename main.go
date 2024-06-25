package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"tutoringsystem/config"
	"tutoringsystem/routes"
	"tutoringsystem/utils"
)

func main() {
	config.LoadEnv()
	utils.ConnectDB()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r := gin.Default()
	r.Use(cors.New(corsConfig))
	r.StaticFile("/favicon.ico", "./web/dist/favicon.ico")
	r.Static("/dist", string(http.Dir("./web/dist")))
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	_ = r.Run(":8080")
}
