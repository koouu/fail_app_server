package model

import (
	"gorm.io/gorm"
)

type Fail struct {
	Id        uint      `json:"id"`
	Content      string    `json:"content" gorm:"type:varchar(128);not null"`
	User_id       uint      `json:"user_id" gorm:"not null"`
}

func (p *Fail) FirstByFailId(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}


func GetAllFail()  []Fail{
	fails := []Fail{}
	DB.Find(&fails)
	return fails
}

func GetUserFail(user_id uint)  []Fail{
	fails := []Fail{}
	DB.Where("user_id = ?", user_id).Find(&fails)
	return fails
}

func (p *Fail) GetUserFail(id uint) (tx *gorm.DB) {
	return DB.Where("user_id = ?", id).First(&p)
}

func (p *Fail) CreateFail() (tx *gorm.DB) {
	return DB.Create(&p)
}

func (p *Fail) SaveFail() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *Fail) UpdatesFail() (tx *gorm.DB) {
	return DB.Model(&p).Updates(p)
}

func (p *Fail) DeleteFail() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *Fail) DeleteByFailId(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&p)
}