package models

import(
	"time"
	"gorm.io/gorm"
)

type Permission struct{
	gorm.Model
	ID uint
	Name string  `gorm:"unique"`
	Value string `gorm:"unique"`
	Role []*Role `gorm:"many2many:role_permission;"`
	Description string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}