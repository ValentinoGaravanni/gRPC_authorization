package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" //sqlite
)

type Storage struct {
	db *sql.DB
}

// DB init
func NewStorage(databasePath string) *Storage {
	db, err := sql.Open("sqlite", databasePath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	//New table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			original_url TEXT NOT NULL,
			short_url TEXT UNIQUE NOT NULL,
			click_count INTEGER DEFAULT 0
		);
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	return &Storage{db: db}
}

// Short URL add
func (s *Storage) AddURL(originalURL, shortURL string) error {
	_, err := s.db.Exec("INSERT INTO urls (original_url, short_url) VALUES (?,?)", originalURL, shortURL)
	return err
}

// Receiving our originalURL
func (s *Storage) GetOrginalURL(shortURL string) (string, error) {
	var originalURL string
	err := s.db.QueryRow("SELECT original_url from urls where short_url = ?", shortURL).Scan(&originalURL)
	return originalURL, err
}

// Add +1 to counter
func (s *Storage) ClickCountInc(shortURL string) error {
	_, err := s.db.Exec("UPDATE urls SET click_count + 1 where short_url = ?", shortURL)
	return err
}

// Stats
func (s *Storage) GetStats(shortURL string) (int, error) {
	var ClickCount int
	err := s.db.QueryRow("SELECT click_count from urls where short_url = ?", shortURL).Scan(&ClickCount)
	return ClickCount, err
}
