package repository

import (
	"errors"
	"projeto-metas/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	var input models.User

	if err := u.db.Where("email = ?", email).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &input, nil

}

func (u *UserRepository) CreateUser(input *models.User) (*models.User, error) {
	if err := u.db.Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}
