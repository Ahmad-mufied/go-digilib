package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

// WebhookHandler represents a handler for Xendit webhooks
type WebhookHandler struct {
	WebhookSecret string
}

// WebhookPayload represents the structure of a Xendit webhook payload
type WebhookPayload struct {
	Id        string  `json:"id"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
	PaidAt    string  `json:"paid_at"`
	UpdatedAt string  `json:"updated_at"`
	Created   string  `json:"created"`
}

// HandleWebhook processes incoming Xendit webhooks
func HandleWebhook(c echo.Context) error {
	// Read the request body
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error reading request body")
	}

	// Parse the webhook payload
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing webhook payload")
	}

	fmt.Printf("Invoice paid: ID=%s, Status=%s\n", payload.Id, payload.Status)

	// Get the invoice URL FROM invoiceData.ExternalID
	invoiceUrl := fmt.Sprintf("https://checkout-staging.xendit.co/web/%s", payload.Id)

	// Update deposit status
	err = repo.Deposit.UpdateDepositStatus(invoiceUrl, constants.PaymentStatusConfirmed)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
