package main

import (
	"linhx1999.com/gin-blog/routes"
	setting "linhx1999.com/gin-blog/utils"
)

//var db = make(map[string]string)

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(setting.HttpPort)
}
