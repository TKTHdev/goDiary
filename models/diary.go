package models

import "time"

type Diary struct {
	ID        int64
	Content   string
	Hash      string
	CreatedAt time.Time
}
