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

func (m *Model) CreateNewURL(url forms.URLRequest) error {
	stmt, err := m.db.Prepare(`
		INSERT INTO urls (original_url, short_url)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&url.OriginalURL, &url.ShortURL)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) GetAllURL() ([]forms.URL, error) {
	var urls []forms.URL

	rows, err := m.db.Query(`
		SELECT url_id, original_url, short_url FROM urls
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url forms.URL
		err := rows.Scan(&url.ID, &url.OriginalURL, &url.ShortURL)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (m *Model) GetAllURLDetail() ([]forms.URLDetail, error) {
	var urls []forms.URLDetail

	rows, err := m.db.Query(`
		SELECT url_id, original_url, short_url, created_at, updated_at FROM urls
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url forms.URLDetail
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
		SELECT url_id, original_url, short_url
		FROM urls
		WHERE url_id = $1
	`)
	if err != nil {
		return url, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&url.ID, &url.OriginalURL, &url.ShortURL)
	if err != nil {
		return url, err
	}

	return url, nil
}

func (m *Model) GetURLDetailByID(id int) (forms.URLDetail, error) {
	var url forms.URLDetail

	stmt, err := m.db.Prepare(`
		SELECT url_id, original_url, short_url, created_at, updated_at
		FROM urls
		WHERE url_id = $1
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
