package models

import (
	"github.com/deyr02/bnzl_crm/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Activity_Category struct {
	gorm.Model
	Title        string `gorm:"" json:"title"`
	Sequence     int    `gorm:"" json:"sequence"`
	Customizable int    `gorm:"" json:"customizable"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate()

}

func GetActivity_Categories() []Activity_Category {
	var activity_categories []Activity_Category
	db.Find(&activity_categories)
	return activity_categories
}

func GetActivity_CategoryById(Id int64) (*Activity_Category, *gorm.DB) {
	var getActivity_category Activity_Category
	db := db.Where("id=?", Id).Find((&getActivity_category))
	return &getActivity_category, db
}

func (ac *Activity_Category) CreateActivity_Category() *Activity_Category {
	db.NewRecord(ac)
	db.Create(&ac)
	return ac
}

func DeleteActivity_Category(Id int64) Activity_Category {
	var ac Activity_Category
	db.Where("ID=?", Id).Delete(&ac)
	return ac
}
