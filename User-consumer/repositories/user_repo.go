package repositories

import (
	"UserConsumer/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) CreateNewUser(user *models.User) error {
	result := userRepository.DB.Create(user)
	return result.Error
}
