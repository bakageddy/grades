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

CREATE TABLE IF NOT EXISTS StudentGradeCache(
	register_no INTEGER REFERENCES Users(register_no),
	total_credits INTEGER NOT NULL,
	subject_credits INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS subject_name_searchidx ON SubjectDatabase(subject_name);
CREATE INDEX IF NOT EXISTS student_grade_cache_fkey ON StudentGradeCache(register_no);
CREATE INDEX IF NOT EXISTS score_sheet_fkey ON ScoreSheet(register_no);
