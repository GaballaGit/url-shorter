-- : one to one 
CREATE TABLE IF NOT EXISTS shortlink (
	shortlink TEXT PRIMARY KEY UNIQUE,
	url TEXT NOT NULL
	);
