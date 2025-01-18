package entity

import "time"

type Admin struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"size:100;unique;not null" json:"email" validation:"required"`
	Password  string    `gorm:"size:255;not null" json:"password" validation:"required"`
	Role      string    `gorm:"size:20;default:admin" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type LoginAdmin struct {
	Email    string `json:"email" validation:"required"`
	Password string `json:"password" validation:"required"`
}

type LoginAdminResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
