package model

import (
	"gorm.io/gorm"
)

type Fail struct {
	Id        uint      `json:"id"`
	Content      string    `json:"content" gorm:"type:varchar(128);not null"`
	User_id       uint      `json:"password" gorm:"not null"`
}

func (p *Fail) FirstByFailId(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}

func (p *Fail) GetALLFail(id uint) (tx *gorm.DB) {
	var fail Fail
	return DB.Find(&fail)
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