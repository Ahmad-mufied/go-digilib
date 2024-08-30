package converter

import (
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"time"
)

// ConvertToDepositResponse Function to convert Deposit DB to DepositResponse
func ConvertToDepositResponse(deposit *data.Deposit) *model.DepositResponse {
	return &model.DepositResponse{
		DepositID:  deposit.ID,
		Amount:     deposit.Amount,
		InvoiceURL: deposit.InvoiceURL,
		Status:     string(deposit.Status),
		CreatedAt:  deposit.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  deposit.UpdatedAt.Format(time.RFC3339),
	}
}
