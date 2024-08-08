package dto

type (
	SendPayload struct {
		Amount      int `json:"amount"`
		ToAccountID int `json:"to_account_id"`
		AccountID   int `json:"account_id"`
	}

	WithdrawPayload struct {
		Amount    int `json:"amount"`
		AccountID int `json:"account_id"`
	}
)
