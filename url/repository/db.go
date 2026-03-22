package repository

import (
	"database/sql"
	"log"
)

type URLRepo struct {
	DB *sql.DB
}

func (repo *URLRepo) Save(longURL, shortCode string) (string, error) {
	var existingCode string
	err := repo.DB.QueryRow(getShortUrlFromLongUrlQuery, longURL).Scan(&existingCode)
	if err == nil {
		log.Println("URL already exists ", longURL)
		return existingCode, nil
	}
	_, err = repo.DB.Exec(insertShortCodeQuery, longURL, shortCode)
	return shortCode, err
}

// for redirecting
func (repo *URLRepo) Get(shortcode string) (string, error) {
	var longUrl string
	err := repo.DB.QueryRow(getlongUrlFromShortCode, shortcode).Scan(&longUrl)
	return longUrl, err
}
