package main

import (
	"demo-user-service/model"
	"demo-user-service/router"
	"demo-user-service/service"

	"github.com/gin-gonic/gin"
)

// @title Data Center API
// @version 0.0.1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	gin.SetMode(gin.DebugMode)
	model.InitDB()
	r := router.Router()
	service.InitData(r.Routes())
	err := r.Run()
	if err != nil {
		return
	}
}
