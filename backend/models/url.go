package models

import (
	"database/sql"

	"github.com/lebrancconvas/ShortenURL/forms"
)

type Model struct {
	db *sql.DB
}

func NewModel(db *sql.DB) *Model {
	return &Model{db: db}
}


func (m *Model) GetAllURL() ([]forms.URL, error) {
	var urls []forms.URL

	rows, err := m.db.Query(`
		SELECT url_id, original_url, short_url, created_at, updated_at FROM url
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url forms.URL
		err := rows.Scan(&url.ID, &url.OriginalURL, &url.ShortURL, &url.CreatedAt, &url.UpdatedAt)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (m *Model) GetURLByID(id int) (forms.URL, error) {
	var url forms.URL

	stmt, err := m.db.Prepare(`
		SELECT url_id, original_url, short_url, created_at, updated_at
		FROM url
		WHERE id = $1
	`)
	if err != nil {
		return url, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&url.ID, &url.OriginalURL, &url.ShortURL, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return url, err
	}

	return url, nil
}
