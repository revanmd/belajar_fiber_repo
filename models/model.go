package models

import (
	"time"

	"gorm.io/gorm"
)

// MasterNature represents the nature of a master entry
type MasterNature struct {
	ID   uint   `gorm:"primaryKey"`
	UID  string `gorm:"unique;not null"`
	Name string `gorm:"not null"`
}

// MasterPriority represents the priority of a master entry
type MasterPriority struct {
	ID   uint   `gorm:"primaryKey"`
	UID  string `gorm:"unique;not null"`
	Name string `gorm:"not null"`
}

// MasterClassification represents the classification of a master entry
type MasterOrganization struct {
	ID          uint   `gorm:"primaryKey"`
	UID         string `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	MessageCode string
}

type UserBelajar struct {
	UID       string         `gorm:"primaryKey"`
	ID        int            `gorm:"unique;not null"`
	Username  string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
