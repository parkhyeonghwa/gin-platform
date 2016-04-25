package models

import (
	"time"
)

type User struct {
	ID         uint64     `gorm:"primary_key" json:"id"`
	Username   string     `sql:"size:255" json:"username"`
	Password   string     `sql:"size:255" json:"password"`
	Realname   string     `sql:"size:255" json:"realname"`
	Department string     `sql:"size:255" json:"department"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
