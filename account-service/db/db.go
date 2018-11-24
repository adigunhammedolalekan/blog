package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func Init(uri string) (*gorm.DB, error) {

	db, err := gorm.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	//Perform DB auto migrate
	db.AutoMigrate(&Account{})
	return db, nil
}

type Account struct {

	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`

	Token string `sql:"-" gorm:"-" json:"token"`
}

func (a *Account) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.New().String()
	return scope.SetColumn("id", uid)
}