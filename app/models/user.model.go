package models

import(
	"time"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	ID uint
	Email string `gorm:"primaryKey"`
	Password string
	Post []Post `gorm:"foreignKey:UserRefer"`
	Comment []Comment `gorm:"foreignKey:UserRefer"`
	Role []*Role `gorm:"many2many:user_role;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}