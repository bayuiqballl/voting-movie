package entity

import "time"

type Vote struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null;index" json:"user_id"`
	MovieID   int       `gorm:"not null;index" json:"movie_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type VoteRequest struct {
	MovieID int  `json:"movie_id" validation:"required"`
	UserID  int  `json:"user_id"`
	IsVote  bool `json:"is_vote" validation:"required"`
}
