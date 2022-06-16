package model

import (
//	"time"
	"gorm.io/gorm"

)

type Post struct {
//	ULIDBaseModel
	Ulid string `json:"ulid" gorm:"PrimaryKey;type:varchar(26);not null"`
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	Age uint `json:"age" gorm:"not null;default:0"`
	Bloodtype string `json:"bloodtype" gorm:"type:varchar(10);not null"`
	Origin string `json:"origin" gorm:"type:varchar(100);not null"`
}

func (p *Post) FirstById(id string) (tx *gorm.DB) {
	return DB.Where("ulid = ?", id).First(&p)
}

func (p *Post) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

// all column update
func (p *Post) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *Post) Updates(id string) (tx *gorm.DB) {
	return DB.Model(&p).Where("ulid = ?", id).Updates(
		Post{
			Name: p.Name,
			Age: p.Age,
			Bloodtype:p.Bloodtype,
			Origin: p.Origin,
		})
}

func (p *Post) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *Post) DeleteById (id string) (tx *gorm.DB) {
	return DB.Where("ulid = ?", id).Delete(&p)
}

func (p *Post) DropTable() (tx *gorm.DB) {
	DB.Migrator().DropTable(&p)

	return nil
}

func (p *Post) CreateTable() (tx *gorm.DB) {
	DB.Migrator().CreateTable(&p)

	return nil
}
