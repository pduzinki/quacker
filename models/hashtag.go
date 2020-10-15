package models

import (
	"regexp"

	"github.com/jinzhu/gorm"
)

// Hashtag represents hashtag data in the database
type Hashtag struct {
	gorm.Model
	Text    string `gorm:"not_null;index"`
	QuackID uint   `gorm:"not_null;index"`
}

// HashtagDB is an interface for interacting with hashtag data in the database
type HashtagDB interface {
	Create(hashtag *Hashtag) error
	Delete(id uint) error
}

// HashtagService is an interface for interacting with hashtag model
type HashtagService interface {
	HashtagDB
	ParseHashtags(text string) []string
}

type hashtagService struct {
	HashtagDB
	hashtagRegex *regexp.Regexp
}

func (hs hashtagService) ParseHashtags(text string) []string {
	hashtags := hs.hashtagRegex.FindAllString(text, -1)
	uniqueHashtags := make([]string, 0)
	keys := make(map[string]bool)

	for _, hashtag := range hashtags {
		if _, prs := keys[hashtag]; !prs {
			keys[hashtag] = true
			uniqueHashtags = append(uniqueHashtags, hashtag)
		}
	}

	return uniqueHashtags
}

// NewHashtagService creates HashtagService instance
func NewHashtagService(db *gorm.DB) HashtagService {
	hg := newHashtagGorm(db)
	hv := newHashtagValidator(hg)

	return hashtagService{
		HashtagDB:    hv,
		hashtagRegex: regexp.MustCompile(`#[a-zA-Z0-9_]+`),
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
	err := runHashtagValidatorFuncs(hashtag,
		hv.quackIDGreaterThanZero,
		hv.truncateHash,
		hv.quackExists)
	if err != nil {
		return err
	}

	return hv.HashtagDB.Create(hashtag)
}

func (hv *hashtagValidator) Delete(id uint) error {
	hashtag := Hashtag{}
	hashtag.ID = id

	err := runHashtagValidatorFuncs(&hashtag, hv.idGreaterThanZero)
	if err != nil {
		return err
	}

	return hv.HashtagDB.Delete(id)
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
	return hg.db.Create(hashtag).Error
}

func (hg *hashtagGorm) Delete(id uint) error {
	hashtag := Hashtag{}
	hashtag.ID = id

	return hg.db.Delete(hashtag).Error
}
