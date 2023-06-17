package routes

import (
	"github.com/elvesbd/goCrud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/id/:id", controller.FindById)
	r.GET("/user/email/:email", controller.FindByEmail)
	r.POST("/user/create", controller.Create)
	r.PUT("/user/id", controller.Update)
	r.DELETE("/user/id", controller.Delete)
}
