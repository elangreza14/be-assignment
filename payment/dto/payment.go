package dto

type (
	SendPayload struct {
		Amount      int    `json:"balance"`
		ToAccountID string `json:"to_account_id"`
	}

	WithdrawPayload struct {
		Amount    int `json:"amount"`
		AccountID int `json:"account_id"`
	}
)
