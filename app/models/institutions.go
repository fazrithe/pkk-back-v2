package models

import (
	"fmt"
	u "pkk-back-v2/utils"

	"github.com/jinzhu/gorm"
)

type Institution struct {
	gorm.Model
	Name    string `gorm:"size:255; not null;unique" json:"name"`
	Address string `gorm:"type:text;" json:"address"`
}

func (institution *Institution) Validate() (map[string]interface{}, bool) {
	if institution.Name == "" {
		return u.Message(false, "Name should be on the payload"), false
	}

	return u.Message(true, "success"), true
}

func GetInstitutions() []*Institution {
	institutions := make([]*Institution, 0)
	err := GetDB().Find(&institutions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return institutions
}

func (institution *Institution) Create() map[string]interface{} {
	if resp, ok := institution.Validate(); !ok {
		return resp
	}

	GetDB().Create(institution)

	resp := u.Message(true, "success")
	resp["institution"] = institution
	return resp
}

func GetInstitution(id int) *Institution {
	institution := &Institution{}
	err := GetDB().First(&institution, id).Error
	if err != nil {
		return nil
	}

	return institution
}

func UpdateInstitution(institution *Institution) (err error) {
	err = GetDB().Save(institution).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func GetInstitutionForUpdateOrDelete(id int, institution *Institution) (err error) {
	if err := db.Where("id = ?", id).First(&institution).Error; err != nil {
		return err
	}

	return nil
}
