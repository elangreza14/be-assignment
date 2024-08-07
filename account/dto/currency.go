package dto

type (
	CurrencyListResponseElement struct {
		Code        string `json:"Code"`
		Description string `json:"Description"`
	}

	CurrencyListResponse []CurrencyListResponseElement
)
