package models

import "time"

type VoiceActing string

const (
	FemaleVoice VoiceActing = "Female"
	MaleVoice   VoiceActing = "Male"
)

type Streamer struct {
	ID       int64
	URL      string
	Name     string
	Password string
}

type Donation struct {
	ID         int64
	Amount     float64
	Message    string
	Voice      VoiceActing
	Name       string
	StreamerID int64
	CreatedAt time.Time
}

type StreamerSettings struct {
	StreamerID int64
	MinDonat   float64
	Picture    string
	Music      string
}
