package routes

import (
	"github.com/elvesbd/goCrud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	controller controller.UserControllerInterface,
) {
	r.GET("/users/id/:id", controller.FindById)
	r.GET("/users/email/:email", controller.FindByEmail)
	r.POST("/users/create", controller.Create)
	r.PUT("/users/id", controller.Update)
	r.DELETE("/users/id", controller.Delete)
}
