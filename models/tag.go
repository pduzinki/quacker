package models

import (
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"

	"quacker/match"
)

// Tag represents tag data in the database
type Tag struct {
	gorm.Model
	Text    string `gorm:"not_null;index"`
	QuackID uint   `gorm:"not_null;index"`
}

// TagDB is an interface for interacting with tag data in the database
type TagDB interface {
	Create(tag *Tag) error
	Delete(id uint) error
}

// TagService is an interface for interacting with tag model
type TagService interface {
	TagDB
	ParseTags(text string) []string
}

type tagService struct {
	TagDB
	tagRegex *regexp.Regexp
}

func (hs tagService) ParseTags(text string) []string {
	uniqueTags := make([]string, 0)
	keys := make(map[string]bool)
	tags := make([]string, 0)

	words := strings.Split(text, " ")

	for _, word := range words {
		if match := hs.tagRegex.FindString(word); match != "" {
			tags = append(tags, match)
		}
	}

	for _, tag := range tags {
		if _, prs := keys[tag]; !prs {
			keys[tag] = true
			uniqueTags = append(uniqueTags, tag)
		}
	}

	return uniqueTags
}

// NewTagService creates TagService instance
func NewTagService(db *gorm.DB) TagService {
	hg := newTagGorm(db)
	hv := newTagValidator(hg)

	return tagService{
		TagDB:    hv,
		tagRegex: regexp.MustCompile(match.Tag()),
	}
}

type tagValidator struct {
	TagDB
}

func newTagValidator(h TagDB) *tagValidator {
	return &tagValidator{
		TagDB: h,
	}
}

func (hv *tagValidator) Create(tag *Tag) error {
	err := runTagValidatorFuncs(tag,
		hv.quackIDGreaterThanZero,
		hv.truncateHash,
		hv.quackExists)
	if err != nil {
		return err
	}

	return hv.TagDB.Create(tag)
}

func (hv *tagValidator) Delete(id uint) error {
	tag := Tag{}
	tag.ID = id

	err := runTagValidatorFuncs(&tag, hv.idGreaterThanZero)
	if err != nil {
		return err
	}

	return hv.TagDB.Delete(id)
}

type tagGorm struct {
	db *gorm.DB
}

func newTagGorm(db *gorm.DB) *tagGorm {
	return &tagGorm{
		db: db,
	}
}

func (hg *tagGorm) Create(tag *Tag) error {
	return hg.db.Create(tag).Error
}

func (hg *tagGorm) Delete(id uint) error {
	tag := Tag{}
	tag.ID = id

	return hg.db.Delete(tag).Error
}
