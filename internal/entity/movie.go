package entity

import "time"

type Movie struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Duration    int       `gorm:"not null" json:"duration"`
	Artists     string    `gorm:"type:text" json:"artists"`
	Genres      string    `gorm:"type:text" json:"genres"`
	WatchURL    string    `gorm:"type:text;not null" json:"watch_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
