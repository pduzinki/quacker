package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// Follow represents 'follow' relation in the database
type Follow struct {
	gorm.Model
	UserID        uint `gorm:"not_null;index"`
	FollowsUserID uint `gorm:"not_null"`
}

// FollowDB is an inferface for interacting with follow data in the database
type FollowDB interface {
	FindByID(id uint) (*Follow, error)
	FindByUserID(id uint) ([]Follow, error)

	Create(follow *Follow) error
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
	follow := Follow{}
	follow.ID = id

	err := runFollowValidatorFuncs(&follow,
		fv.idGreaterThanZero)
	if err != nil {
		return nil, err
	}

	return fv.FollowDB.FindByID(id)
}

func (fv *followValidator) FindByUserID(id uint) ([]Follow, error) {
	follow := Follow{}
	follow.UserID = id

	err := runFollowValidatorFuncs(&follow,
		fv.userIDGreaterThanZero)
	if err != nil {
		return nil, err
	}

	return fv.FollowDB.FindByUserID(id)
}

func (fv *followValidator) Create(follow *Follow) error {
	err := runFollowValidatorFuncs(follow,
		fv.userIDGreaterThanZero,
		fv.followsUserIDGraterThanZero)
	if err != nil {
		return err
	}

	return fv.FollowDB.Create(follow)
}

func (fv *followValidator) Delete(id uint) error {
	follow := Follow{}
	follow.ID = id

	err := runFollowValidatorFuncs(&follow,
		fv.userIDGreaterThanZero)
	if err != nil {
		return err
	}

	return fv.FollowDB.Delete(id)
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
	follow := Follow{}
	err := fg.db.Where("id = ?", id).First(&follow).Error
	if err == gorm.ErrRecordNotFound {

	} else if err != nil {

	}
	return &follow, nil
}

func (fg *followGorm) FindByUserID(id uint) ([]Follow, error) {
	follows := make([]Follow, 1)

	err := fg.db.Where("user_id = ?", id).Find(&follows).Error
	if err != nil {
		return nil, err
	}

	return follows, nil
}

func (fg *followGorm) Create(follow *Follow) error {
	return fg.db.Create(follow).Error
}

func (fg *followGorm) Delete(id uint) error {
	follow := Follow{}
	follow.ID = id

	return fg.db.Delete(follow).Error
}
