package core

type UserType string

const (
	Basic   UserType = "basic"
	Admin   UserType = "admin"
	Premium UserType = "premium"
)

type Film struct {
	Title                string `json:"title"`
	PosterFile           string `json:"poster_file"`
	CompressedPosterFile string `json:"compressed_poster_file"`
	Url                  string `json:"url"`
}

type Films struct {
	Results []Film `json:"results"`
}

type ManePageResponse struct {
	NewFilms         Films `json:"newFilms"`
	IBMRatingFilms   Films `json:"IBMRatingFilms"`
	LocalRatingFilms Films `json:"localRatingFilms"`
}
