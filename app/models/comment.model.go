package models

import(
	"time"
	"gorm.io/gorm"
)

type Comment struct{
	gorm.Model
	ID uint
	Content string `gorm:"not null;type:text"`
	HasAttachment bool 
	CommentAttachment []CommentAttachment `gorm:"foreignKey:CommentRefer"`
	UserRefer uint
	PostRefer uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}