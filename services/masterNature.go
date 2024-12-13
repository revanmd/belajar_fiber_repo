package services

import (
	"belajar-api/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterNatureService struct {
	DB *gorm.DB
}

// GetAllMasterPriorities fetches all MasterNature records
func (s *MasterNatureService) GetAllMasterNatures() ([]models.MasterNature, error) {
	var priorities []models.MasterNature
	if err := s.DB.Find(&priorities).Error; err != nil {
		return nil, err
	}
	return priorities, nil
}

// GetMasterNatureByUID fetches a MasterNature by its UID
func (s *MasterNatureService) GetMasterNatureByUID(UID string) (*models.MasterNature, error) {
	var Nature models.MasterNature
	if err := s.DB.Where("UID = ?", UID).First(&Nature).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &Nature, nil
}

// CreateMasterNature creates a new MasterNature
func (s *MasterNatureService) CreateMasterNature(Nature *models.MasterNature) error {
	Nature.UID = uuid.New().String() // generate unique UID
	return s.DB.Create(Nature).Error
}

// UpdateMasterNature updates an existing MasterNature by its UID
func (s *MasterNatureService) UpdateMasterNature(UID string, updatedNature *models.MasterNature) error {
	var Nature models.MasterNature
	if err := s.DB.Where("UID = ?", UID).First(&Nature).Error; err != nil {
		return err
	}

	// Update fields
	Nature.Name = updatedNature.Name

	return s.DB.Save(&Nature).Error
}

// DeleteMasterNature deletes a MasterNature by its UID
func (s *MasterNatureService) DeleteMasterNature(UID string) error {
	var Nature models.MasterNature
	if err := s.DB.Where("UID = ?", UID).First(&Nature).Error; err != nil {
		return err
	}
	return s.DB.Delete(&Nature).Error
}
