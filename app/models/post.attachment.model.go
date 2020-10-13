package models

import(
	"time"
	"gorm.io/gorm"
)

type PostAttachment struct{
	gorm.Model
	ID uint
	Url string  `gorm:"not null"`
	AttachmentType string // need to implement enum here
	PostRefer uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}