package models

import(
	"time"
	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	ID uint
	Content string `gorm:"not null;type:text"`
	HasAttachment bool 
	Comment []Comment `gorm:"foreignKey:PostRefer"`
	PostAttachment []PostAttachment `gorm:"foreignKey:PostRefer"`
	UserRefer uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}