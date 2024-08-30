package handler

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"github.com/Ahmad-mufied/go-digilib/model/converter"
	"github.com/Ahmad-mufied/go-digilib/server/middleware"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func CreateDeposit(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	userId := claims.UserID

	// Get the deposit from the request
	depositPayload := new(model.DepositRequest)
	err := c.Bind(&depositPayload)
	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest, "Invalid input")
	}

	// Validate
	err = validate.Struct(depositPayload)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Get Wallet id
	walletId, err := repo.Wallet.GetWalletIdByUserID(userId)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	deposit := &data.Deposit{
		Amount:     depositPayload.Amount,
		WalletID:   walletId,
		InvoiceURL: "https://www.google.com",
		Status:     constants.PaymentStatusPending,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Create Deposit
	depositId, err := repo.Deposit.CreateDeposit(deposit)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	deposit.ID = depositId

	// Create Xendit invoice
	xenditInvoice, err := converter.ConvertDepositToXenditInvoice(deposit)
	if err != nil {
		return err
	}

	invoice, err := utils.CreateInvoice(xenditInvoice)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	// Set invoice URL
	invoiceUrl := invoice.InvoiceUrl
	deposit.InvoiceURL = invoiceUrl

	// Update invoice URL
	err = repo.Deposit.UpdateDepositInvoiceURL(depositId, invoiceUrl)

	depositResponse := converter.ConvertToDepositResponse(deposit)

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Deposit created successfully",
		Data:    depositResponse,
	})

}

func GetDepositById(c echo.Context) error {
	depositId := utils.StringToUint(c.Param("id"))

	deposit, err := repo.Deposit.GetDepositById(depositId)
	if err != nil {
		return utils.HandleError(c, constants.ErrNotFound, "Deposit not found")
	}

	depositResponse := converter.ConvertToDepositResponse(deposit)

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Deposit retrieved successfully",
		Data:    depositResponse,
	})
}
