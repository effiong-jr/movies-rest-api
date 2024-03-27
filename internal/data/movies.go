package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   int32     `json:"runtime"`
	Genres    []string  `json:"genres"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}
