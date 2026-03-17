package database

import (
	"database/sql"
	"donation-app/models"
)

func GetSettings(db *sql.DB, streamerID int64) (models.StreamerSettings, error) {
	query := `
	SELECT streamer_id, min_donat, picture, music FROM streamer_settings WHERE streamer_id = $1
	`
	var settings models.StreamerSettings
	err := db.QueryRow(query, streamerID).Scan(&settings.StreamerID, &settings.MinDonat, &settings.Picture, &settings.Music)
	return settings, err
}

func UpdateSettings(db *sql.DB, streamerID int64, minDonat float64, picture, music string) error {
	query := `
	UPDATE streamer_settings
	SET min_donat = $2, picture = $3, music = $4
	WHERE streamer_id = $1
	`
	_, err := db.Exec(query, streamerID, minDonat, picture, music)
	return err
}
