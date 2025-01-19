package entity

import "time"

type Movie struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255;not null" json:"title" validation:"required"`
	Description string    `gorm:"type:text" json:"description" validation:"required"`
	Duration    int       `gorm:"not null" json:"duration" validation:"required"`
	Artists     string    `gorm:"type:text" json:"artists" validation:"required"`
	Genres      string    `gorm:"type:text" json:"genres" validation:"required"`
	WatchURL    string    `gorm:"type:text;not null" json:"watch_url" validation:"required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UploadMovieResponse struct {
	WatchURL string `json:"watch_url"`
}

type GetListMovieRequest struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

type GetMostDataMovie struct {
	MostView  MostView  `json:"most_view"`
	MostGenre MostGenre `json:"most_genre"`
	MostVoted MostVoted `json:"most_voted"`
}

type MostView struct {
	MovieID int    `json:"movie_id"`
	Title   string `json:"title"`
	Count   int    `json:"count"`
}

type MostGenre struct {
	Genre string `json:"genre"`
	Count int    `json:"count"`
}

type MostVoted struct {
	MovieID int    `json:"movie_id"`
	Title   string `json:"title"`
	Count   int    `json:"count"`
}
