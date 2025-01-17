package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:50;not null" json:"username"`
	Email     string    `gorm:"size:100;unique;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
