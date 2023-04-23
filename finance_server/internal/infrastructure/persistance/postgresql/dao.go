package postgresql

import (
	"context"
	"errors"

	"finance_server/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgrePersistanceService struct {
	client *gorm.DB
}

func (s PostgrePersistanceService) RegisterUserDAO(ctx context.Context, us entity.Sale) (entity.Sale, error) {

	var existingUser entity.Sale
	err := s.client.First(&existingUser, "username = ?", us.SubRegion).Error
	if err == nil {
		return entity.Sale{}, errors.New("username already exists")
	}

	// username does not exist, create a new user record
	err = s.client.Create(&us).Error
	if err != nil {
		return entity.Sale{}, err
	}
	return us, nil
}
