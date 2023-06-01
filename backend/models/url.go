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

func (m *Model) CheckExistOriginalURL(originalURL string) (bool, error) {
	var exist bool

	stmt, err := m.db.Prepare(`
		SELECT EXISTS (
			SELECT 1 FROM urls WHERE original_url = $1
		)
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(&originalURL).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (m *Model) StoreURL(originalURL string) error {
	stmt, err := m.db.Prepare(`
		INSERT INTO urls (original_url)
		VALUES ($1)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Query(&originalURL)
	if err != nil {
		return err
	}

	return nil
}

func (m *Model) UpdateNewShortenURL(originalURL string, shortURL string) error {
	stmt, err := m.db.Prepare(`
		UPDATE urls
		SET short_url = $1
		WHERE original_url = $2
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(&shortURL, &originalURL)
	if err != nil {
		return err
	}

	return nil
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


func (m *Model) GetURLByShortURL(shortURL string) (forms.URL, error) {
	var url forms.URL

	stmt, err := m.db.Prepare(`
		SELECT url_id, original_url, short_url
		FROM urls
		WHERE short_url = $1
	`)
	if err != nil {
		return url, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(&shortURL)
	if err != nil {
		return url, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&url.ID, &url.OriginalURL, &url.ShortURL)
		if err != nil {
			return url, err
		}
	}

	return url, nil
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

func (m *Model) GetShortURLByID(id int) (string, error) {
	var shortURL string

	stmt, err := m.db.Prepare(`
		SELECT short_url
		FROM urls
		WHERE url_id = $1
	`)
	if err != nil {
		return shortURL, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(&shortURL)
	if err != nil {
		return shortURL, err
	}

	return shortURL, nil
}

func (m *Model) GetShortURLByOriginalURL(originalURL string) (string, error) {
	var shortURL string

	stmt, err := m.db.Prepare(`
		SELECT short_url
		FROM urls
		WHERE original_url = $1
	`)
	if err != nil {
		return shortURL, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(&originalURL)

	err = row.Scan(&shortURL)
	if err != nil {
		return shortURL, err
	}

	return shortURL, nil
}

func (m *Model) GetOriginalURLByShortURL(shortURL string) (string, error) {
	var originalURL string

	stmt, err := m.db.Prepare(`
		SELECT original_url
		FROM urls
		WHERE short_url = $1
		LIMIT 1
	`)
	if err != nil {
		return originalURL, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(&shortURL)

	err = row.Scan(&originalURL)
	if err != nil {
		return originalURL, err
	}

	return originalURL, nil
}


