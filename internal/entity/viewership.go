package entity

import "time"

type Viewership struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null;index" json:"user_id"`
	MovieID   int       `gorm:"not null;index" json:"movie_id"`
	WatchedAt time.Time `gorm:"autoCreateTime" json:"watched_at"`
}
