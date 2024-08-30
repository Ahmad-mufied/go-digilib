package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Ahmad-mufied/go-digilib/config"
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// WebhookHandler represents a handler for Xendit webhooks
type WebhookHandler struct {
	WebhookSecret string
}

// WebhookPayload represents the structure of a Xendit webhook payload
type WebhookPayload struct {
	Event      string          `json:"event"`
	Created    string          `json:"created"`
	Data       json.RawMessage `json:"data"`
	BusinessID string          `json:"business_id"`
}

// HandleWebhook processes incoming Xendit webhooks
func HandleWebhook(c echo.Context) error {
	// Read the request body
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error reading request body")
	}

	// Verify the webhook signature
	if !verifySignature(c.Request().Header.Get("X-Callback-Token"), body) {
		return c.String(http.StatusUnauthorized, "Invalid signature")
	}

	// Parse the webhook payload
	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing webhook payload")
	}

	// Handle different event types
	switch payload.Event {
	case "invoice.paid":
		return handleInvoicePaid(c, payload)
	case "invoice.expired":
		return handleInvoiceExpired(c, payload)
	// Add more cases for other event types as needed
	default:
		fmt.Printf("Unhandled event type: %s\n", payload.Event)
	}

	// Respond with a 200 OK to acknowledge receipt of the webhook
	return c.NoContent(http.StatusOK)
}

// verifySignature verifies the webhook signature
func verifySignature(token string, body []byte) bool {
	webhookSecret := config.Viper.GetString("XENDIT_WEBHOOK_SECRET")
	mac := hmac.New(sha256.New, []byte(webhookSecret))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(token), []byte(expectedMAC))
}

// handleInvoicePaid processes a paid invoice event
func handleInvoicePaid(c echo.Context, payload WebhookPayload) error {
	var invoiceData struct {
		ID         string `json:"id"`
		ExternalID string `json:"external_id"`
		Amount     int64  `json:"amount"`
		Status     string `json:"status"`
	}

	if err := json.Unmarshal(payload.Data, &invoiceData); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing invoice data")
	}

	fmt.Printf("Invoice paid: ID=%s, ExternalID=%s, Amount=%d\n",
		invoiceData.ID, invoiceData.ExternalID, invoiceData.Amount)

	// Add your business logic here for handling a paid invoice
	log.Printf("invoice paid: ID=%s, ExternalID=%s, Amount=%d\n", invoiceData.ID, invoiceData.ExternalID, invoiceData.Amount)

	// Get the invoice URL FROM invoiceData.ExternalID
	invoiceUrl := fmt.Sprintf("https://checkout-staging.xendit.co/web/%s", invoiceData.ExternalID)

	// Update deposit status
	err := repo.Deposit.UpdateDepositStatus(invoiceUrl, constants.PaymentStatusConfirmed)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// handleInvoiceExpired processes an expired invoice event
func handleInvoiceExpired(c echo.Context, payload WebhookPayload) error {
	var invoiceData struct {
		ID         string `json:"id"`
		ExternalID string `json:"external_id"`
		Status     string `json:"status"`
	}

	if err := json.Unmarshal(payload.Data, &invoiceData); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing invoice data")
	}

	fmt.Printf("Invoice expired: ID=%s, ExternalID=%s\n",
		invoiceData.ID, invoiceData.ExternalID)

	// Add your business logic here for handling an expired invoice

	return c.NoContent(http.StatusOK)
}

// ExtractInvoiceIDFromURL extracts the invoice ID from a Xendit invoice URL
func ExtractInvoiceIDFromURL(url string) string {
	// Split the URL by '/' and return the last part
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
