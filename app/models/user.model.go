package models

import(
	"time"
	"gorm.io/gorm"
)

type User struct{
	Id uint `gorm:"autoIncrement;not null;"`
	Email string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
	Salt string `gorm:"not null"`
	AccountLocked bool `gorm:"DEFAULT:true"`
	Post []Post `gorm:"foreignKey:UserRefer"`
	Comment []Comment `gorm:"foreignKey:UserRefer"`
	Role []*Role `gorm:"many2many:user_role;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}