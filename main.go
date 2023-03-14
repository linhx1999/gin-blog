package main

import (
	"linhx1999.com/gin-blog/config"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/routers"
)

//var db = make(map[string]string)

func main() {
	models.InitDB()
	config.LoadConf()

	r := routers.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(config.HttpPort)
}
