package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ahmad-mufied/go-digilib/config"
	"io"
	"net/http"
)

// Invoice represents the structure of a Xendit invoice
type Invoice struct {
	ExternalID                     string                 `json:"external_id"`
	Amount                         int64                  `json:"amount"`
	Description                    string                 `json:"description"`
	InvoiceDuration                int                    `json:"invoice_duration"`
	Customer                       Customer               `json:"customer"`
	InvoiceUrl                     string                 `json:"invoice_url"`
	CustomerNotificationPreference NotificationPreference `json:"customer_notification_preference,omitempty"`
	SuccessRedirectURL             string                 `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                 `json:"failure_redirect_url,omitempty"`
	Currency                       string                 `json:"currency"`
	Items                          []Item                 `json:"items"`
	Fees                           []Fee                  `json:"fees,omitempty"`
}

type Customer struct {
	GivenNames   string    `json:"given_names,omitempty"`
	Surname      string    `json:"surname,omitempty"`
	Email        string    `json:"email,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	Addresses    []Address `json:"addresses,omitempty"`
}

type Address struct {
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	State       string `json:"state,omitempty"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
}

type NotificationPreference struct {
	InvoiceCreated  []string `json:"invoice_created,omitempty"`
	InvoiceReminder []string `json:"invoice_reminder,omitempty"`
	InvoicePaid     []string `json:"invoice_paid,omitempty"`
}

type Item struct {
	Name     string `json:"name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	Price    int64  `json:"price,omitempty"`
	Category string `json:"category,omitempty"`
	URL      string `json:"url,omitempty"`
}

type Fee struct {
	Type  string `json:"type,omitempty"`
	Value int64  `json:"value,omitempty"`
}

// CreateInvoice creates a new invoice using the Xendit API
func CreateInvoice(invoice *Invoice) (*Invoice, error) {

	// Xendit API URL
	xenditAPIURL := config.Viper.GetString("XENDIT_API_URL")

	// Get API KEY
	apiKey := config.Viper.GetString("XENDIT_API_KEY")

	jsonData, err := json.Marshal(invoice)
	if err != nil {
		return nil, fmt.Errorf("error marshaling invoice: %v", err)
	}

	req, err := http.NewRequest("POST", xenditAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
	}

	var createdInvoice Invoice
	err = json.Unmarshal(body, &createdInvoice)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &createdInvoice, nil
}
