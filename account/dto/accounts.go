package dto

type (
	CreateAccountPayload struct {
		CurrencyCode string `json:"currency_code"`
		ProductID    int    `json:"product_id"`
	}

	AccountListResponseElement struct {
		CurrencyCode string `json:"currency_code"`
		Balance      int    `json:"balance"`
		Name         string `json:"name"`
		Status       string `json:"status"`
		ProductID    int    `json:"product_id"`
	}

	AccountListResponse []AccountListResponseElement
)
