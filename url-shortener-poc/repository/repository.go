package repository

import (
	"database/sql"
	"url-shortener/models"
)

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) Save(url *models.URL) error {
	_, err := r.DB.Exec("INSERT INTO urls (short_code, original_url) VALUES (?, ?)", url.ShortCode, url.OriginalURL)
	return err
}

func (r *URLRepository) FindByShortCode(code string) (*models.URL, error) {
	row := r.DB.QueryRow("SELECT id, short_code, original_url, created_at FROM urls WHERE short_code = ?", code)
	u := &models.URL{}
	err := row.Scan(&u.ID, &u.ShortCode, &u.OriginalURL, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
