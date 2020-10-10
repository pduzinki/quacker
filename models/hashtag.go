package models

import (
	"github.com/jinzhu/gorm"
)

// Hashtag ...
type Hashtag struct {
	gorm.Model
	Text    string `gorm:"not_null;index"`
	QuackID uint   `gorm:"not_null;index"`
}

// HashtagDB ...
type HashtagDB interface {
	Create(hashtag *Hashtag) error
	Delete(id uint) error
}

// HashtagService ...
type HashtagService interface {
	HashtagDB
}

type hashtagService struct {
	HashtagDB
}

// NewHashtagService ...
func NewHashtagService(db *gorm.DB) HashtagService {
	hg := newHashtagGorm(db)
	hv := newHashtagValidator(hg)

	return hashtagService{
		HashtagDB: hv,
	}
}

type hashtagValidator struct {
	HashtagDB
}

func newHashtagValidator(h HashtagDB) *hashtagValidator {
	return &hashtagValidator{
		HashtagDB: h,
	}
}

func (hv *hashtagValidator) Create(hashtag *Hashtag) error {
	// TODO
	return nil
}

func (hv *hashtagValidator) Delete(id uint) error {
	// TODO
	return nil
}

type hashtagGorm struct {
	db *gorm.DB
}

func newHashtagGorm(db *gorm.DB) *hashtagGorm {
	return &hashtagGorm{
		db: db,
	}
}

func (hg *hashtagGorm) Create(hashtag *Hashtag) error {
	// TODO
	return nil
}

func (hg *hashtagGorm) Delete(id uint) error {
	// TODO
	return nil
}
