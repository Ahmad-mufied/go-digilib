package converter

import (
	"fmt"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"strconv"
	"time"
)

func ConvertDepositToXenditInvoice(deposit *data.Deposit) (*utils.Invoice, error) {
	amountInRupiah := int64(deposit.Amount)

	// Generate a unique external ID using deposit ID and timestamp
	externalID := fmt.Sprintf("deposit_%d_%d", deposit.ID, time.Now().Unix())

	// Create the Invoice struct
	invoice := &utils.Invoice{
		ExternalID:      externalID,
		Amount:          amountInRupiah,
		Description:     fmt.Sprintf("Deposit #%d", deposit.ID),
		InvoiceDuration: 86400, // 24 hours in seconds, adjust as needed
		InvoiceUrl:      deposit.InvoiceURL,
		Currency:        "IDR", // Assuming Indonesian Rupiah, change if different
		Customer:        utils.Customer{
			// Fill in customer details if available in your system
		},
		Items: []utils.Item{
			{
				Name:     fmt.Sprintf("Deposit #%d", deposit.ID),
				Quantity: 1,
				Price:    amountInRupiah,
				Category: "deposit",
			},
		},
	}

	return invoice, nil
}

// Helper function to convert string to int64
func stringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
