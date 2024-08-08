package dto

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
)
