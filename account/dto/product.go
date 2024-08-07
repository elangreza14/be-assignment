package dto

type (
	ProductListResponseElement struct {
		ID          int    `json:"id"`
		Name        string `json:"Name"`
		Description string `json:"Description"`
	}

	ProductListResponse []ProductListResponseElement
)
