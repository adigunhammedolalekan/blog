package handler

import (
	"context"
	"github.com/adigunhammedolalekan/blog/account-service/db"
	account "github.com/adigunhammedolalekan/blog/account-service/proto/account"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/errors"
	"golang.org/x/crypto/bcrypt"
)

var JWT_SECRET  = "_)(*&^%$%^&*()(*&^$%^&*(*&NOTAREALKEY"

type AccountHandler struct {
	db *gorm.DB
}

func NewAccountHandlerService(db *gorm.DB) *AccountHandler {

	return &AccountHandler{
		db:db,
	}
}

func (h *AccountHandler) CreateAccount(ctx context.Context, req *account.Account, res *account.Response) error {

	acc := GetAccountByAttribute(h.db, "email", req.Email)
	if acc != nil {

		res.Status = false
		res.Account = nil
		res.Message = "Email address already exists"
		return errors.Conflict("", "Email address already exists")
	}

	newAccount := &models.Account{
		Name: req.Name,
		Email: req.Token,
		Password: GenerateHashedPassword(req.Password),
	}

	if err := h.db.Create(newAccount).Error; err != nil {
		res.Status = false
		return errors.InternalServerError("", "Failed to create account %v", err.Error())
	}

	req.Id = newAccount.Id
	req.Token = GenerateHashedPassword(req.Id)
	res.Status = true
	res.Account = req
	return nil
}

func (h *AccountHandler) Authenticate(ctx context.Context, req *account.Account, res *account.Response) error {

	authAccount := GetAccountByAttribute(h.db, "email", req.Email)
	if authAccount == nil {
		res.Status = false
		res.Message = "Invalid email/password combination"
		return errors.Parse("Invalid email/password combination")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authAccount.Password), []byte(req.Password)); err != nil {
		res.Status = false
		res.Message = "Invalid email/password combination"
		return errors.Parse("Invalid email/password combination")
	}

	authAccount.Password = ""
	authAccount.Token = GenerateJWT(authAccount.Id)
	res.Status = true
	res.Account = &account.Account{
		Name:authAccount.Name, Email: authAccount.Email, Token:authAccount.Token, Id:authAccount.Id,
	}

	return nil
}

func (h *AccountHandler) GetAccount(ctx context.Context, req *account.GetAccountRequest, res *account.Response) error {

	value := GetAccountByAttribute(h.db, "id", req.UserId)
	res.Account = &account.Account{
		Name: value.Name, Email:value.Email, Token: value.Token, Id: value.Id,
	}

	return nil
}

func GetAccountByAttribute(db *gorm.DB, attribute string, value interface{}) *models.Account {

	result := &models.Account{}
	err := db.Table("accounts").Where(attribute + " = ?", value).First(result).Error
	if err != nil {
		return nil
	}

	return result
}

type Token struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func GenerateJWT(id string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{Account:id})
	tokenString, _ := token.SignedString(JWT_SECRET)
	return tokenString
}

func GenerateHashedPassword(password string) (hashed string) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashed = string(bytes)
	return
}