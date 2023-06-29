package main

import (
	"log"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/controller"
	"github.com/elvesbd/goCrud/src/controller/routes"
	"github.com/elvesbd/goCrud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init dependencies
	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
