package db

import (
	"database/sql"
	"donation-app/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateStreamer(db *sql.DB, streamer models.Streamer) (int64, error) {
	query := `
	INSERT INTO streamers (name, url, password)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	hash, err := bcrypt.GenerateFromPassword([]byte(streamer.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	var id int64
	err = db.QueryRow(query, streamer.Name, streamer.URL, hash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetStreamerByID(db *sql.DB, id int64) (models.Streamer, error) {
	query := `
	SELECT id, name, url, password
	FROM streamers
	WHERE id = $1
	`
	var streamer models.Streamer
	err := db.QueryRow(query, id).Scan(&streamer.ID, &streamer.Name, &streamer.URL, &streamer.Password)
	if err != nil {
		return models.Streamer{}, err
	}
	return streamer, nil
}

func GetStreamerByName(db *sql.DB, name string) (models.Streamer, error) {
	query := `
	SELECT id, name, url, password
	FROM streamers
	WHERE name = $1
	`
	var streamer models.Streamer
	err := db.QueryRow(query, name).Scan(&streamer.ID, &streamer.Name, &streamer.URL, &streamer.Password)
	if err != nil {
		return models.Streamer{}, err
	}
	return streamer, nil
}
