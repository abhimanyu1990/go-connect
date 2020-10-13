package models

import(
	"time"
	"gorm.io/gorm"
)

type Role struct{
	gorm.Model  
	ID uint
	Name string  `gorm:"unique;not null;"`
	Value string `gorm:"unique;not null;"`
	Description string `gorm:"type:text"`
	User []*User `gorm:"many2many:user_role;"`
	Permission []*Permission `gorm:"many2many:role_permission;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}