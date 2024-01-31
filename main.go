package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"log"
	"second_hand/controller"
	_ "second_hand/docs"
	"second_hand/logic"
)

// @title Your API Title
// @description This is a sample API for demonstration purposes. It includes multiple endpoints for various operations.
// @version 1.0
// @host localhost:8080
func main() {
	r := gin.Default()
	//err := login.Register()
	//if err != nil {
	//	log.Fatal(err)

	r.Static("/css", "./templates/css")
	r.Static("/js", "./templates/js")
	r.Static("/imgs", "./imgs")
	r.LoadHTMLGlob("./templates/html/*")
	_ = logic.InitMySQL()
	//controller.CreateRecords(logic.DB)
	controller.CollectRoutes(r)
	defer logic.Close()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(":8080"))
}
