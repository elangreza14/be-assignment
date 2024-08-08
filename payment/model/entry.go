package model

import (
	"database/sql"
	"time"
)

type Entry struct {
	ID        int `db:"id"`
	AccountID int `db:"account_id"`
	Amount    int `db:"amount"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Entry) TableName() string {
	return "entries"
}

func (c Entry) Columns() []string {
	return []string{
		"id",
		"account_id",
		"amount",
	}
}

// to create in DB
func (c Entry) Data() map[string]any {
	return map[string]any{
		"id":         c.ID,
		"account_id": c.AccountID,
		"amount":     c.Amount,
	}
}
