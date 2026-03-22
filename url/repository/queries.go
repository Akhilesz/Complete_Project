package repository

const (
	getShortUrlFromLongUrlQuery = `SELECT shortcode FROM urls WHERE longurl=$1`
	getlongUrlFromShortCode     = `SELECT longurl FROM urls WHERE shortcode=$1`
	insertShortCodeQuery        = `INSERT INTO urls (longurl,shortcode) VALUES ($1,$2)`
)
