package main

import (
	"linhx1999.com/gin-blog/config"
	"linhx1999.com/gin-blog/models"
	"linhx1999.com/gin-blog/routers"
)

//var db = make(map[string]string)

func main() {
	//config.LoadConf()
	models.InitDB()

	r := routers.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(config.HttpPort)
}
