package handler

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"github.com/Ahmad-mufied/go-digilib/server/middleware"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func MakeNewBorrow(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	userId := claims.UserID

	// Get the borrow from the request
	borrowPayload := new(model.BorrowRequest)
	err := c.Bind(&borrowPayload)
	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest, "Invalid input")
	}

	// Validate
	err = validate.Struct(borrowPayload)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// convert string startDate and endDate to time.Time
	startDate, _ := time.Parse("2006-01-02", borrowPayload.StartDate)
	endDate, _ := time.Parse("2006-01-02", borrowPayload.EndDate)

	duration := endDate.Sub(startDate)
	days := int(duration.Hours() / 24)

	// 2. Tentukan tipe durasi
	var durationType constants.DurationType
	switch {
	case days <= 7:
		durationType = constants.DurationTypeDaily
	case days <= 30:
		durationType = constants.DurationTypeWeekly
	default:
		durationType = constants.DurationTypeMonthly
	}

	basePrice, err := repo.Book.GetBookBasePrice(borrowPayload.BookID)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	borrowPrice, err := repo.BorrowPrice.GetBorrowPrice(borrowPayload.BookID, durationType)

	// Calculate total price
	totalPrice := basePrice * borrowPrice.PriceMultiplier * float64(days)

	borrow := &data.Borrow{
		UserID:     userId,
		BookID:     borrowPayload.BookID,
		Status:     constants.BookStatusEnum(constants.BorrowStatusPending),
		StartDate:  startDate,
		EndDate:    endDate,
		TotalPrice: totalPrice,
	}

	// Create Borrow
	borrowId, err := repo.Borrow.CreateBorrow(borrow)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	borrow.ID = borrowId

	// Update borrow status asynchronously after  1 minute
	go func() {
		time.Sleep(1 * time.Minute)
		err := repo.Borrow.UpdateBorrowStatusConfirm(borrowId, constants.BookStatusEnum(constants.BorrowStatusSuccess))
		if err != nil {
			log.Printf("Error updating borrow status: %v", err)
		}
	}()

	return c.JSON(http.StatusCreated, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Borrow created successfully",
		Data:    borrow,
	})

}

func GetBorrowById(c echo.Context) error {

	borrowId := utils.StringToUint(c.Param("id"))

	borrow, err := repo.Borrow.GetBorrowById(borrowId)
	if err != nil {
		return utils.HandleError(c, constants.ErrNotFound, "Borrow not found")
	}

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success Getting Borrow",
		Data:    borrow,
	})
}

func GetAllBorrowsByUserID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	userId := claims.UserID

	borrows, err := repo.Borrow.GetAllBorrowsByUserID(userId)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success Getting All Borrows",
		Data:    borrows,
	})
}

func UpdateBorrowReturnedAt(c echo.Context) error {

	returnBookPayload := new(model.ReturnBookRequest)
	err := c.Bind(&returnBookPayload)

	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest, "Invalid input")
	}

	// Validate
	err = validate.Struct(returnBookPayload)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	err = repo.Borrow.UpdateBorrowReturnedAt(returnBookPayload.BorrowID)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Borrow returned successfully",
	})
}
