package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

func (c Product) TableName() string {
	return "products"
}

func (c Product) Columns() []string {
	return []string{
		"id",
		"name",
		"description",
	}
}

// to create in DB
func (c Product) Data() map[string]any {
	return map[string]any{
		"id":          c.ID,
		"name":        c.Name,
		"description": c.Description,
	}
}
