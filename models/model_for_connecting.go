package models

// it's for fields in database
type Song struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	Artist      string `json:"artist"`
	Title       string `json:"title"`
	Music_genre string `json:"music_genre"`
}
