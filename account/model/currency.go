package model

import (
	"database/sql"
	"time"
)

type Currency struct {
	Code        string `db:"code"`
	Description string `db:"description"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Currency) TableName() string {
	return "currencies"
}

func (c Currency) Columns() []string {
	return []string{
		"code",
		"description",
	}
}

// to create in DB
func (c Currency) Data() map[string]any {
	return map[string]any{
		"code":        c.Code,
		"description": c.Description,
	}
}
