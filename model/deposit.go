package model

// DepositRequest /users/deposit
// request body -> { amount }
type DepositRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

// DepositResponse /users/deposit
// response body -> { deposit_id, amount, invoice_url, status, created_at, updated_at}
type DepositResponse struct {
	DepositID  uint    `json:"deposit_id"`
	Amount     float64 `json:"amount"`
	InvoiceURL string  `json:"invoice_url"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
