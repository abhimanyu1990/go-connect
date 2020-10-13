package models


import(
	"time"
	"gorm.io/gorm"

)

type ForgotPasswordToken struct{
	gorm.Model
	ID uint
	Email *string  `gorm:"not null;unique"`
	Token string   `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAT time.Time

}