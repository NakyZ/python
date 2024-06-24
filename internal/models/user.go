package models

import "gorm.io/gorm"

type User struct{
	gorm.Model

	Admin bool 
	Usename string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
	Nama string `gorm:"type:varchar(50)"`
}

type UserModel struct{
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) *UserModel {
	return &UserModel{
		db: connection,
	}
}	

func (um *UserModel) Login(username string, password string) (User, error) {
	var result User
	err := um.db.Where("username = ? and password = ?", username, password).First(&result).Error
	if err != nil{
		return User{}, err
	}
	return result, nil 
}