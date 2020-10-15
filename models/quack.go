package models

import (
	"github.com/jinzhu/gorm"
)

// Quack represents quack data in the database
type Quack struct {
	gorm.Model
	UserID   uint   `gorm:"not_null;index"`
	Username string `gorm:"-"`
	Text     string `gorm:"not null"`
}

// QuackDB is an interface for interacting with quack data in the database
type QuackDB interface {
	FindByID(id uint) (*Quack, error)
	FindByUserID(id uint) ([]Quack, error)
	FindByMultipleUserIDs(ids []uint) ([]Quack, error)
	FindByHashtag(hashtag string) ([]Quack, error)
	// TODO perhaps add later: FindByUserIDWithLimit(id, limit, offset uint) ([]Quack, error)

	Create(quack *Quack) error
	// No Update method, since updating quack is forbidden.
	Delete(id uint) error
}

// QuackService is an interface for interacting with quack model
type QuackService interface {
	QuackDB
}

type quackService struct {
	QuackDB
}

// NewQuackService creates QuackService instance
func NewQuackService(db *gorm.DB) QuackService {
	qg := newQuackGorm(db)
	qv := newQuackValidator(qg)

	return &quackService{
		qv,
	}
}

type quackValidator struct {
	QuackDB
}

func newQuackValidator(q QuackDB) *quackValidator {
	return &quackValidator{
		QuackDB: q,
	}
}

func (qv *quackValidator) FindByID(id uint) (*Quack, error) {
	q := Quack{}
	q.ID = id

	err := runQuackValidatorFuncs(&q,
		qv.idGreaterThanZero)
	if err != nil {
		return nil, err
	}

	return qv.QuackDB.FindByID(id)
}

func (qv *quackValidator) FindByUserID(id uint) ([]Quack, error) {
	q := Quack{}
	q.UserID = id

	err := runQuackValidatorFuncs(&q,
		qv.userIDGreaterThanZero)
	if err != nil {
		return nil, err
	}

	return qv.QuackDB.FindByUserID(id)
}

func (qv *quackValidator) FindByMultipleUserIDs(ids []uint) ([]Quack, error) {
	for _, id := range ids {
		q := Quack{}
		q.UserID = id

		err := runQuackValidatorFuncs(&q,
			qv.userIDGreaterThanZero)
		if err != nil {
			return nil, err
		}
	}

	return qv.QuackDB.FindByMultipleUserIDs(ids)
}

func (qv *quackValidator) FindByHashtag(hashtag string) ([]Quack, error) {
	return qv.QuackDB.FindByHashtag(hashtag)
}

func (qv *quackValidator) Create(quack *Quack) error {
	err := runQuackValidatorFuncs(quack,
		qv.userIDGreaterThanZero,
		qv.TextShorterThan160chars)
	if err != nil {
		return err
	}

	return qv.QuackDB.Create(quack)
}

func (qv *quackValidator) Delete(id uint) error {
	quack := &Quack{}
	quack.ID = id

	err := runQuackValidatorFuncs(quack, qv.idGreaterThanZero)
	if err != nil {
		return err
	}
	return qv.QuackDB.Delete(id)
}

type quackGorm struct {
	db *gorm.DB
}

func newQuackGorm(db *gorm.DB) *quackGorm {
	return &quackGorm{
		db: db,
	}
}

func (qg *quackGorm) FindByID(id uint) (*Quack, error) {
	q := Quack{}

	err := qg.db.Model(Quack{}).
		Select("quacks.id, quacks.created_at, quacks.text, quacks.user_id, users.username").
		Joins("inner join users on quacks.user_id = users.id").
		Where("quacks.id = ?", id).
		First(&q).Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrRecordNotFound
	} else if err != nil {
		return nil, err
	}

	return &q, nil
}

func (qg *quackGorm) FindByUserID(id uint) ([]Quack, error) {
	q := make([]Quack, 1)

	err := qg.db.Model(Quack{}).
		Select("quacks.id, quacks.created_at, quacks.text, users.username").
		Joins("inner join users on quacks.user_id = users.id").
		Where("user_id = ?", id).
		Order("id desc").
		Find(&q).Error
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (qg *quackGorm) FindByMultipleUserIDs(ids []uint) ([]Quack, error) {
	q := make([]Quack, 1)

	err := qg.db.Model(Quack{}).
		Select("quacks.id, quacks.created_at, quacks.text, users.username").
		Joins("inner join users on quacks.user_id = users.id").
		Where("user_id IN (?)", ids).
		Order("id desc").
		Find(&q).Error
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (qg *quackGorm) FindByHashtag(hashtag string) ([]Quack, error) {
	q := make([]Quack, 1)

	err := qg.db.Model(Quack{}).
		Select("quacks.id, quacks.created_at, quacks.text, users.username").
		Joins("inner join users on quacks.user_id = users.id").
		Joins("inner join hashtags on hashtags.quack_id = quacks.id").
		Where("hashtags.text = (?)", hashtag).
		Order("id desc").
		Find(&q).Error
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (qg *quackGorm) Create(quack *Quack) error {
	return qg.db.Create(quack).Error
}

func (qg *quackGorm) Delete(id uint) error {
	quack := Quack{}
	quack.ID = id

	return qg.db.Delete(quack).Error
}
