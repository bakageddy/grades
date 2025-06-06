CREATE TABLE IF NOT EXISTS Users(
	register_no INTEGER PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS ScoreSheet(
	register_no INTEGER REFERENCES Users(register_no),
	grade TEXT,
	credits INTEGER,
	subject_code TEXT
);

CREATE TABLE IF NOT EXISTS SubjectDatabase(
	subject_code TEXT PRIMARY KEY,
	subject_name TEXT UNIQUE NOT NULL,
	subject_credits INTEGER NOT NULL
);
