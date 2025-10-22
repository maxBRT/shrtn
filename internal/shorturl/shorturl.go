package shorturl

import (
	"database/sql"
	"errors"
	"net/url"
)

type Url struct {
	BaseUrl  string `json:"base_url"`
	ShortUrl string `json:"short_url"`
}

func (u *Url) Validate() error {
	if u.BaseUrl == "" {
		return errors.New("base_url is required")
	}
	if u.ShortUrl == "" {
		return errors.New("short_url is required")
	}
	_, err := url.Parse(u.BaseUrl)
	if err != nil {
		return errors.New("base_url is not a valid url")
	}
	_, err = url.Parse(u.ShortUrl)
	if err != nil {
		return errors.New("short_url is not a valid url")
	}
	return nil
}

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) CreateLink(url Url) (*Url, error) {
	query := `INSERT INTO urls (base_url, short_url) VALUES ($1, $2) RETURNING *`
	var u Url
	err := r.db.QueryRow(query, url.BaseUrl, url.ShortUrl).Scan(u.BaseUrl, u.ShortUrl)
	return &u, err
}

func (r *UrlRepository) GetLinks() ([]*Url, error) {
	query := `SELECT * FROM urls`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var urls []*Url
	for rows.Next() {
		var u Url
		err := rows.Scan(&u.BaseUrl, &u.ShortUrl)
		if err != nil {
			return nil, err
		}
		urls = append(urls, &u)
	}
	return urls, nil
}

func (r *UrlRepository) GetLink(shortUrl string) (string, error) {
	query := `SELECT base_url FROM urls WHERE short_url = $1`
	var baseUrl string
	err := r.db.QueryRow(query, shortUrl).Scan(&baseUrl)
	return baseUrl, err
}
