package mapper

import (
	"github.com/elvesbd/goCrud/src/model"
	"github.com/elvesbd/goCrud/src/model/repository/entity"
)

func DomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
