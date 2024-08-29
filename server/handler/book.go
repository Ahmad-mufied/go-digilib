package handler

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/model"
	"github.com/Ahmad-mufied/go-digilib/model/converter"
	"github.com/Ahmad-mufied/go-digilib/server/middleware"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func CreateBook(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JWTCustomClaims)
	role := claims.Role

	if role != "admin" {
		return utils.HandleError(c, constants.ErrForbidden, "Only admin can create a book")
	}

	var addBookRequest model.AddBookRequest
	err := c.Bind(&addBookRequest)
	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest, "Invalid input")
	}

	// Validate
	err = validate.Struct(addBookRequest)
	if err != nil {
		// Format the validation errors
		errors := utils.FormatValidationErrors(err)
		return utils.HandleValidationError(c, errors)
	}

	// Check if the book already exists using SKU
	isExist, err := repo.Book.CheckBookBySKU(addBookRequest.SKU)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	if isExist {
		return utils.HandleError(c, constants.ErrConflict, "Book already exists")
	}

	// Custom validation for PublishedAt (if needed)
	if addBookRequest.PublishedAt != nil {
		_, err := time.Parse("2006-01-02", *addBookRequest.PublishedAt)
		if err != nil {
			return utils.HandleError(c, constants.ErrBadRequest, "published_at must be in the format YYYY-MM-DD")
		}
	}

	// Convert AddBookRequest to Book and BookPhysicalDetails
	book, d, err := converter.ConvertAddBookRequestToBook(addBookRequest)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	// Create Book
	bookID, err := repo.Book.CreateBook(book, d)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	// Get Current Stock
	stock, err := repo.Book.GetStockById(bookID)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success create book",
		Data: model.AddBookResponse{
			ID:     bookID,
			Title:  book.Title,
			Status: string(book.Status),
			Stock:  stock,
		},
	})

}

func GetAllBooks(c echo.Context) error {

	books, err := repo.Book.GetAllBooks()
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success Getting All Books",
		Data:    books,
	})
}
