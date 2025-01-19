package entity

import "time"

type Viewership struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null;index:idx_user_movie,unique" json:"user_id" validation:"required"`
	MovieID   int       `gorm:"not null;index:idx_user_movie,unique" json:"movie_id" validation:"required"`
	Duration  int       `gorm:"not null" json:"duration"`
	WatchedAt time.Time `gorm:"autoCreateTime" json:"watched_at"`
}
