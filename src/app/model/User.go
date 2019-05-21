package model

import (
	"github.com/jinzhu/gorm"
)


type User struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
	Type   string `json:"user"`
}

func (u *User) Disable(){
	u.Status=false
}
func (u *User) Enable(){
	u.Status=false
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}