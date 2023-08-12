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

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email,
			)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, restErr.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, restErr.NewNotFoundError(errorMessage)
	}

	logger.Info(
		"FindUserByEmail repository executed successfully",
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByEmail"),
	)

	return mapper.EntityToDomain(*userEntity), nil
}
