package dto

type (
	SendPayload struct {
		Amount      int    `json:"balance"`
		ToAccountID string `json:"toAccountID"`
	}

	WithdrawPayload struct {
		Amount int `json:"balance"`
	}
)
