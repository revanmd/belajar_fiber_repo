package services

import (
	"belajar-api/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterOrganizationService struct {
	DB *gorm.DB
}

// GetAllMasterOrganizations fetches all MasterOrganization records
func (s *MasterOrganizationService) GetAllMasterOrganizations() ([]models.MasterOrganization, error) {
	var MasterOrganizations []models.MasterOrganization
	if err := s.DB.Find(&MasterOrganizations).Error; err != nil {
		return nil, err
	}
	return MasterOrganizations, nil
}

// GetMasterOrganizationByUID fetches a MasterOrganization by its UID
func (s *MasterOrganizationService) GetMasterOrganizationByUID(UID string) (*models.MasterOrganization, error) {
	var MasterOrganization models.MasterOrganization
	if err := s.DB.Where("UID = ?", UID).First(&MasterOrganization).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &MasterOrganization, nil
}

// CreateMasterOrganization creates a new MasterOrganization
func (s *MasterOrganizationService) CreateMasterOrganization(MasterOrganization *models.MasterOrganization) error {
	MasterOrganization.UID = uuid.New().String()
	return s.DB.Create(MasterOrganization).Error
}

// UpdateMasterOrganization updates an existing MasterOrganization by its UID
func (s *MasterOrganizationService) UpdateMasterOrganization(UID string, updatedMasterOrganization *models.MasterOrganization) error {
	var MasterOrganization models.MasterOrganization
	if err := s.DB.Where("UID = ?", UID).First(&MasterOrganization).Error; err != nil {
		return err
	}
	MasterOrganization.UID = UID
	MasterOrganization.Name = updatedMasterOrganization.Name
	MasterOrganization.MessageCode = updatedMasterOrganization.MessageCode
	return s.DB.Save(&MasterOrganization).Error
}

// DeleteMasterOrganization deletes a MasterOrganization by its UID
func (s *MasterOrganizationService) DeleteMasterOrganization(UID string) error {
	var MasterOrganization models.MasterOrganization
	if err := s.DB.Where("UID = ?", UID).First(&MasterOrganization).Error; err != nil {
		return err
	}
	return s.DB.Delete(&MasterOrganization).Error
}
