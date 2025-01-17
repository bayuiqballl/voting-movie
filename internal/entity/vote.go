package entity

import "time"

type Vote struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null;index" json:"user_id"`
	MovieID   int       `gorm:"not null;index" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
