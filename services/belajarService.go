package services

import (
	"belajar-api/models"
	"errors"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBelajarService struct {
	DB *gorm.DB
}

// GetAllUserBelajar fetches all UserBelajar records
func (s *UserBelajarService) GetAllUserBelajar() ([]models.UserBelajar, error) {
	var users []models.UserBelajar
	if err := s.DB.Select("username", "password").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserBelajarByUID fetches a UserBelajar by its UID
func (s *UserBelajarService) GetUserBelajarByUID(UID string) (*models.UserBelajar, error) {
	var user models.UserBelajar
	if err := s.DB.Where("UID = ?", UID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUserBelajar creates a new UserBelajar
func (s *UserBelajarService) CreateUserBelajar(user *models.UserBelajar) error {
	user.UID = uuid.New().String() // generate unique UID
	user.CreatedAt = time.Now()    // set creation time
	return s.DB.Create(user).Error
}

// UpdateUserBelajar updates an existing UserBelajar by its UID
func (s *UserBelajarService) UpdateUserBelajar(UID string, updatedUser *models.UserBelajar) error {
	var user models.UserBelajar
	if err := s.DB.Where("UID = ?", UID).First(&user).Error; err != nil {
		return err
	}

	// Update fields
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	return s.DB.Save(&user).Error
}

// DeleteUserBelajar deletes a UserBelajar by its UID
func (s *UserBelajarService) DeleteUserBelajar(UID string) error {
	var user models.UserBelajar
	if err := s.DB.Where("UID = ?", UID).First(&user).Error; err != nil {
		return err
	}
	return s.DB.Delete(&user).Error
}
