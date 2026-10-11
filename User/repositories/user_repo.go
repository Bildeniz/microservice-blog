package repositories

import (
	"User/models"

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

func (userRepository *UserRepository) GetAllUsers() ([]models.User, int64, error) {
	var users []models.User
	results := userRepository.DB.Find(&users)

	return users, results.RowsAffected, results.Error
}

func (userRepository *UserRepository) GetUser(id int) (models.User, int64, error) {
	var user models.User
	result := userRepository.DB.Find(&user, id)

	return user, result.RowsAffected, result.Error
}
