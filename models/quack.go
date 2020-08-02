package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // imports postgres driver
)

// Quack represents quack data in the database
type Quack struct {
	gorm.Model
	userID uint   `gorm:"not_null;index"`
	text   string `gorm:"not null"`
}

// QuackDB is an interface for interacting with quack data in the database
type QuackDB interface {
	FindByID(id uint) (*Quack, error)
	FindByUserID(id uint) ([]Quack, error)

	Create(q *Quack) error
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
func NewQuackService() {

}

type quackValidator struct {
	QuackDB
}

func (qv *quackValidator) FindByID(id uint) (*Quack, error) {
	return nil, nil
}

func (qv *quackValidator) FindByUserID(id uint) ([]Quack, error) {
	return nil, nil
}

func (qv *quackValidator) Create(q *Quack) error {
	return nil
}

func (qv *quackValidator) Delete(id uint) error {
	return nil
}

func newQuackValidator() {

}

type quackGorm struct {
	db *gorm.DB
}

func (qg *quackGorm) FindByID(id uint) (*Quack, error) {
	return nil, nil
}

func (qg *quackGorm) FindByUserID(id uint) ([]Quack, error) {
	return nil, nil
}

func (qg *quackGorm) Create(q *Quack) error {
	return nil
}

func (qg *quackGorm) Delete(id uint) error {
	return nil
}
