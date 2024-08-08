package dto

import "time"

type (
	CreateAccountPayload struct {
		CurrencyCode string `json:"currency_code"`
		ProductID    int    `json:"product_id"`
	}

	AccountListResponseElement struct {
		CurrencyCode string `json:"currency_code"`
		Name         string `json:"name"`
		ProductID    int    `json:"product_id"`
		ID           int    `json:"id"`
	}

	AccountListResponse []AccountListResponseElement

	TransferHistoryResponseElement struct {
		ID            int       `json:"id"`
		FromAccountID int       `json:"from_account_id"`
		ToAccountID   int       `json:"to_account_id"`
		Amount        int       `json:"amount"`
		Action        string    `json:"action"`
		CreatedAt     time.Time `json:"created_at"`
	}

	TransferHistoryResponse struct {
		Histories []TransferHistoryResponseElement `json:"histories"`
		Balance   int                              `json:"balance"`
	}
)
