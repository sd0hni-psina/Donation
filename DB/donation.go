package db

import (
	"database/sql"
	"donation-app/models"
)

func CreateDonation(db *sql.DB, donation models.Donation) error {
	query := `
	INSERT INTO donations (amount, message, voice, name, streamer_id)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, donation.Amount, donation.Message, donation.Voice, donation.Name, donation.StreamerID)
	return err
}

func GetDonationsByStreamerID(db *sql.DB, id int64) ([]models.Donation, error) {
	query := `
	SELECT * FROM donations
	WHERE streamer_id = $1
	`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var donations []models.Donation
	for rows.Next() {
		var donation models.Donation
		if err := rows.Scan(&donation.ID, &donation.Amount, &donation.Message, &donation.Voice, &donation.Name, &donation.StreamerID, &donation.CreatedAt); err != nil {
			return nil, err
		}
		donations = append(donations, donation)
	}
	return donations, nil
}
