package main

import (
	"github.com/elvesbd/goCrud/src/controller"
	"github.com/elvesbd/goCrud/src/model/repository"
	"github.com/elvesbd/goCrud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(userService)
}
