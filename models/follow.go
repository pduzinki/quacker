package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// Follow represents 'follow' relation in the database
type Follow struct {
	gorm.Model
	UserID         uint `gorm:"not_null;index"`
	FollowedUserID uint `gorm:"not_null"`
}

// FollowDB is an inferface for interacting with follow data in the database
type FollowDB interface {
	// TODO
}

// FollowService is an inferface for interacting with follow model
type FollowService interface {
	FollowDB
}

type followService struct {
	FollowDB
}

// NewFollowService creates FollowService instance
func NewFollowService(db *gorm.DB) FollowService {
	return nil
}

type followValidator struct {
	FollowDB
}

type followGorm struct {
	FollowDB
}
