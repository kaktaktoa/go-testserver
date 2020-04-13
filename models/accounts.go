package models

import (
	util "../utils"

	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func (account *Account) Validate() (map[string] interface{}, bool)	{
	if !strings.Contains(account.Email, "@") {
		return util.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return util.Message(false, "Password is required"), false
	}

	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return util.Message(false, "Connection error"), false
	}

	if temp.Email != "" {
		return util.Message(false, "Email already"), false
	}

	return util.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{})  {
	if resp, ok  := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return util.Message(false, "Failed to create account, connection error.")
	}

	//Создать новый токен JWT для новой зарегистрированной учётной записи
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //удалить пароль

	response := util.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?").First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return util.Message(false, "Not found")
		}
		return util.Message(false, "Error connect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		return util.Message(false, "Error connect")
	}
	account.Password = ""
	tk := &Token{}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	account.Token, _ = token.SignedString([]byte(os.Getenv("token_password")))
	resp := util.Message(true, "Logged")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" {
		return nil
	}
	acc.Password = ""
	return acc
}
