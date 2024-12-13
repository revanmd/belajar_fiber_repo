package services

import (
	"belajar-api/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterPriorityService struct {
	DB *gorm.DB
}

// GetAllMasterPriorities fetches all MasterPriority records
func (s *MasterPriorityService) GetAllMasterPriorities() ([]models.MasterPriority, error) {
	var priorities []models.MasterPriority
	if err := s.DB.Find(&priorities).Error; err != nil {
		return nil, err
	}
	return priorities, nil
}

// GetMasterPriorityByUID fetches a MasterPriority by its UID
func (s *MasterPriorityService) GetMasterPriorityByUID(UID string) (*models.MasterPriority, error) {
	var priority models.MasterPriority
	if err := s.DB.Where("UID = ?", UID).First(&priority).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &priority, nil
}

// CreateMasterPriority creates a new MasterPriority
func (s *MasterPriorityService) CreateMasterPriority(priority *models.MasterPriority) error {
	priority.UID = uuid.New().String() // generate unique UID
	return s.DB.Create(priority).Error
}

// UpdateMasterPriority updates an existing MasterPriority by its UID
func (s *MasterPriorityService) UpdateMasterPriority(UID string, updatedPriority *models.MasterPriority) error {
	var priority models.MasterPriority
	if err := s.DB.Where("UID = ?", UID).First(&priority).Error; err != nil {
		return err
	}

	// Update fields
	priority.Name = updatedPriority.Name

	return s.DB.Save(&priority).Error
}

// DeleteMasterPriority deletes a MasterPriority by its UID
func (s *MasterPriorityService) DeleteMasterPriority(UID string) error {
	var priority models.MasterPriority
	if err := s.DB.Where("UID = ?", UID).First(&priority).Error; err != nil {
		return err
	}
	return s.DB.Delete(&priority).Error
}
