package models

import (
	util "../utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"`
}

func (contact *Contact) Validate() (map[string] interface{}, bool) {
	if contact.Name == "" && contact.Phone == "" && contact.UserId <= 0 {
		return util.Message(false, "Not valid"), false
	}

	return util.Message(true, "ok"), true
}

func (contact *Contact) Create() (map[string] interface{}) {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}
	GetDB().Create(contact)
	resp := util.Message(true, "Ok")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) (*Contact)  {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) ([]*Contact) {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}