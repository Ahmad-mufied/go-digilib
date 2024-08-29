package converter

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"time"
)

// ConvertAddBookRequestToBook converts AddBookRequest to Book and BookPhysicalDetails.
func ConvertAddBookRequestToBook(req model.AddBookRequest) (*data.Book, *data.BookPhysicalDetails, error) {
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
