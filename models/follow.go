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
	FindByID(id uint) (*Follow, error)
	FindByUserID(id uint) ([]Follow, error)

	Create(f *Follow) error
	// No update method, not needed.
	Delete(id uint) error
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
	fg := newFollowGorm(db)
	fv := newFollowValidator(fg)

	return followService{
		FollowDB: fv,
	}
}

type followValidator struct {
	FollowDB
}

func newFollowValidator(f FollowDB) *followValidator {
	return &followValidator{
		FollowDB: f,
	}
}

func (fv *followValidator) FindByID(id uint) (*Follow, error) {
	return nil, nil
}

func (fv *followValidator) FindByUserID(id uint) ([]Follow, error) {
	return nil, nil
}

func (fv *followValidator) Create(f *Follow) error {
	return nil
}

func (fv *followValidator) Delete(id uint) error {
	return nil
}

type followGorm struct {
	db *gorm.DB
}

func newFollowGorm(db *gorm.DB) *followGorm {
	return &followGorm{
		db: db,
	}
}

func (fg *followGorm) FindByID(id uint) (*Follow, error) {
	return nil, nil
}

func (fg *followGorm) FindByUserID(id uint) ([]Follow, error) {
	return nil, nil
}

func (fg *followGorm) Create(f *Follow) error {
	return nil
}

func (fg *followGorm) Delete(id uint) error {
	return nil
}
