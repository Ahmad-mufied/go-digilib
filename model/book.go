package model

import "github.com/jmoiron/sqlx/types"

type BooksDetails struct {
	Id              uint            `json:"id,omitempty"`
	CategoryID      uint            `json:"category_id" validate:"required"`
	ISBN            string          `json:"isbn" validate:"required"`
	SKU             string          `json:"sku" validate:"required"`
	Author          types.JSONText  `json:"author" validate:"required"`
	Title           string          `json:"title" validate:"required"`
	Image           string          `json:"image" validate:"required"`
	Pages           int             `json:"pages" validate:"required"`
	Language        string          `json:"language" validate:"required"`
	Description     string          `json:"description" validate:"required"`
	Stock           int             `json:"stock" validate:"required"`
	Status          string          `json:"status" validate:"required"`
	BorrowedCount   int             `json:"borrowed_count"`
	PublishedAt     *string         `json:"published_at,omitempty"` // Omit if not provided
	BasePrice       float64         `json:"base_price" validate:"required"`
	PhysicalDetails PhysicalDetails `json:"physical_details" validate:"required"`
}

type PhysicalDetails struct {
	Weight float64 `json:"weight" validate:"required"`
	Height int     `json:"height" validate:"required"`
	Width  int     `json:"width" validate:"required"`
}

type CreateUpdateBookResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
	Stock  int    `json:"stock"`
}

type UpdateBookRequest struct {
	BookId uint `json:"book_id" validate:"required"`
	Stock  int  `json:"stock" `
}
