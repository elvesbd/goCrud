package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/elvesbd/goCrud/src/configuration/logger"
	"github.com/elvesbd/goCrud/src/configuration/restErr"
	"github.com/elvesbd/goCrud/src/model"
	"github.com/elvesbd/goCrud/src/model/repository/entity"
	"github.com/elvesbd/goCrud/src/model/repository/entity/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id,
			)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByID"),
			)
			return nil, restErr.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByID"),
		)
		return nil, restErr.NewNotFoundError(errorMessage)
	}

	logger.Info(
		"FindUserByID repository executed successfully",
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByID"),
	)

	return mapper.EntityToDomain(*userEntity), nil
}
