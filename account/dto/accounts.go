package dto

type (
	CreateAccountPayload struct {
		CurrencyCode string `json:"currency_code"`
	}

	AccountListResponseElement struct {
		CurrencyCode string `json:"currency_code"`
		Balance      int    `json:"balance"`
		Name         string `json:"name"`
		Status       string `json:"status"`
	}

	AccountListResponse []AccountListResponseElement
)
