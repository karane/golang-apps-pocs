package models

import "time"

type URL struct {
	ID          int
	ShortCode   string
	OriginalURL string
	CreatedAt   time.Time
}
