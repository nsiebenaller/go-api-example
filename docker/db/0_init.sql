
-- Setup extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

-- Setup tables
CREATE FUNCTION updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE members (
	id TEXT NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TRIGGER member_updated_at
	BEFORE UPDATE
	ON members
	FOR EACH ROW
EXECUTE PROCEDURE updated_at();

CREATE TABLE books (
	id TEXT NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	title TEXT NOT NULL,
	author TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TRIGGER book_updated_at
	BEFORE UPDATE
	ON books
	FOR EACH ROW
EXECUTE PROCEDURE updated_at();

CREATE TABLE member_books (
	member_id TEXT NOT NULL,
	book_id TEXT NOT NULL,
	FOREIGN KEY (member_id) REFERENCES members(id),
	FOREIGN KEY (book_id) REFERENCES books(id)
);

-- Seed data
INSERT INTO members(id, first_name, last_name) VALUES('1', 'Nick', 'Siebenaller');
INSERT INTO books(id, title, author) VALUES('1', 'My Big ole Book', 'Gary E.');