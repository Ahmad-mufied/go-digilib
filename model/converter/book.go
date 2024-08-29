package converter

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"time"
)

// ConvertAddBookRequestToBook converts BooksDetails to Book and BookPhysicalDetails.
func ConvertAddBookRequestToBook(req *model.BooksDetails) (*data.Book, *data.BookPhysicalDetails, error) {
	// Parse PublishedAt string to time.Time
	var publishedAt *time.Time
	if req.PublishedAt != nil {
		parsedTime, err := time.Parse("2006-01-02", *req.PublishedAt)
		if err != nil {
			return &data.Book{}, &data.BookPhysicalDetails{}, err
		}
		publishedAt = &parsedTime
	}

	// Convert status from string to constants.BookStatusEnum
	var status constants.BookStatusEnum
	switch req.Status {
	case "available":
		status = constants.BookStatusAvailable
	default:
		status = constants.BookStatusNotAvailable
	}

	book := data.Book{
		CategoryID:    req.CategoryID,
		ISBN:          req.ISBN,
		SKU:           req.SKU,
		Author:        req.Author,
		Title:         req.Title,
		Image:         req.Image,
		Pages:         req.Pages,
		Language:      req.Language,
		Description:   req.Description,
		Stock:         req.Stock,
		Status:        status,
		BorrowedCount: req.BorrowedCount,
		PublishedAt:   publishedAt,
		BasePrice:     req.BasePrice,
	}

	physicalDetails := data.BookPhysicalDetails{
		Weight: req.PhysicalDetails.Weight,
		Height: req.PhysicalDetails.Height,
		Width:  req.PhysicalDetails.Width,
	}

	return &book, &physicalDetails, nil
}

type PhysicalDetails struct {
	Weight float64
	Height int
	Width  int
}

// ConvertToBooksDetails Function to convert Book and BookPhysicalDetails to BooksDetails
func ConvertToBooksDetails(book *data.Book, physicalDetails *data.BookPhysicalDetails) (*model.BooksDetails, error) {
	details := new(model.BooksDetails)
	newPhysicalDetails := new(model.PhysicalDetails)

	newPhysicalDetails.Weight = physicalDetails.Weight
	newPhysicalDetails.Height = physicalDetails.Height
	newPhysicalDetails.Width = physicalDetails.Width

	details.Id = book.ID
	details.CategoryID = book.CategoryID
	details.ISBN = book.ISBN
	details.SKU = book.SKU
	details.Author = book.Author
	details.Title = book.Title
	details.Image = book.Image
	details.Pages = book.Pages
	details.Language = book.Language
	details.Description = book.Description
	details.Stock = book.Stock
	details.Status = string(book.Status)
	details.BorrowedCount = book.BorrowedCount
	if book.PublishedAt != nil {
		publishedAtStr := book.PublishedAt.Format(time.RFC3339)
		details.PublishedAt = &publishedAtStr
	}
	details.BasePrice = book.BasePrice

	details.PhysicalDetails = *newPhysicalDetails

	return details, nil
}
