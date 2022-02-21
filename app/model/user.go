package model

import (
	"gorm.io/gorm"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(20);not null"`
	Password       string      `json:"password" gorm:"type:varchar(64);not null"`
}

func (p *User) FirstByUserId(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}

func GetAllUser()  []User{
	users := []User{}
	DB.Find(&users)
	return users
}

func (p *User) GetLogIn(name string,password string) (tx *gorm.DB){
	return DB.Where("name = ? AND password = ?", name,password).First(&p)
}

func (p *User) CreateUser() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *User) SaveUser() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *User) UpdatesUser() (tx *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *User) DeleteUser() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *User) DeleteByUserId(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&p)
}