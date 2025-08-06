package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          uint
	Todo        string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
	DueTo       sql.NullTime
}
