package model

import (
	"database/sql"
	"time"
)

type Transfer struct {
	ID            int `db:"id"`
	ToAccountID   int `db:"to_account_id"`
	FromAccountID int `db:"from_account_id"`
	Amount        int `db:"amount"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Transfer) TableName() string {
	return "transfers"
}

func (c Transfer) Columns() []string {
	return []string{
		"id",
		"to_account_id",
		"from_account_id",
		"amount",
	}
}

// to create in DB
func (c Transfer) Data() map[string]any {
	return map[string]any{
		"id":              c.ID,
		"to_account_id":   c.ToAccountID,
		"from_account_id": c.FromAccountID,
		"amount":          c.Amount,
	}
}
