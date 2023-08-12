package repository

import (
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	Create(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *restErr.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *restErr.RestErr)
}
