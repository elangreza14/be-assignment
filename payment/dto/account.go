package dto

import "github.com/google/uuid"

type CreateAccountPayload struct {
	UserID       uuid.UUID
	Name         string
	CurrencyCode string
	ProductID    int
}
