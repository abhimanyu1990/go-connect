package models


import(
	"time"
	"gorm.io/gorm"
)

type Profile struct{
	gorm.Model
	ID uint
	UserRefer string  `gorm:"unique"`
	User User `gorm:"foreignKey:UserRefer; references:Email; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  // belongTo relation to establish one to one relationship
	FirstName string `gorm:"not null"`
	LastName string  `gorm:"not null"`
	AddressLine1 string
	AddressLine2 string
	City string
	Pin string
	Country string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

}