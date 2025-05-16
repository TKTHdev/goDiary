package database

import (
	"database/sql"

	"diary/models"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func InitDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "./diary.db")
	if err != nil {
		return nil, err
	}

	// テーブルの作成
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS diaries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			hash TEXT NOT NULL UNIQUE,
			created_at DATETIME NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func CreateDiary(db *DB, diary models.Diary) error {
	_, err := db.Exec(
		"INSERT INTO diaries (content, hash, created_at) VALUES (?, ?, ?)",
		diary.Content, diary.Hash, diary.CreatedAt,
	)
	return err
}

func GetAllDiaries(db *DB) ([]models.Diary, error) {
	rows, err := db.Query("SELECT id, content, hash, created_at FROM diaries ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var diaries []models.Diary
	for rows.Next() {
		var diary models.Diary
		err := rows.Scan(&diary.ID, &diary.Content, &diary.Hash, &diary.CreatedAt)
		if err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}
	return diaries, nil
}

func DeleteDiary(db *DB, hash string) error {
	_, err := db.Exec("DELETE FROM diaries WHERE hash = ?", hash)
	return err
}
