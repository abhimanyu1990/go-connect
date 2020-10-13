package models


import(
	"time"
	"gorm.io/gorm"

)

type RegistrationToken struct{
	gorm.Model
	ID uint
	Email *string  `gorm:"unique;not null;"`
	Token string   `gorm:"unique;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAT time.Time

}