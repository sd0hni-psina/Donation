CREATE TABLE streamers (
	id       BIGSERIAL PRIMARY KEY,
	url      TEXT,
	name     TEXT,
	password TEXT
);
CREATE TABLE donations (
	id         BIGSERIAL PRIMARY KEY,
	amount     DECIMAL(10, 2),
	message    TEXT,
	voice      TEXT,
	name       TEXT,
	streamer_id BIGINT,
	created_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY (streamer_id) REFERENCES streamers(id)
);

CREATE TABLE streamer_settings (
	streamer_id BIGINT PRIMARY KEY,
	min_donat   DECIMAL(10, 2),
	picture     TEXT,
	music       TEXT,
	FOREIGN KEY (streamer_id) REFERENCES streamers(id)
);
