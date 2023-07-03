package repository

import (
	"context"
	"os"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) Create(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init create user repository")

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, restErr.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, restErr.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))
	return userDomain, nil
}
