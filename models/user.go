package models

import (
	"time"
)

type User struct {
	ID         uint64     `gorm:"primary_key" form:"id"`
	Username   string     `sql:"size:255" form:"username" binding:"required"`
	Password   string     `sql:"size:255" form:"password" binding:"required"`
	Realname   string     `sql:"size:255" form:"realname"`
	Department string     `sql:"size:255" form:"department"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
