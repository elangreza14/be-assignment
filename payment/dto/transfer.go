package dto

import "time"

type (
	TransferHistoryResponse struct {
		ID            int       `json:"id"`
		FromAccountID int       `json:"from_account_id"`
		ToAccountID   int       `json:"to_account_id"`
		Amount        int       `json:"amount"`
		Action        string    `json:"action"`
		CreatedAt     time.Time `json:"created_at"`
	}
)
