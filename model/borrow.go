package model

type BorrowRequest struct {
	BookID    uint   `json:"book_id" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type BorrowResponse struct {
	ID         uint    `json:"id"`
	UserID     uint    `json:"user_id"`
	BookID     uint    `json:"book_id"`
	Status     string  `json:"status"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	TotalPrice float64 `json:"total_price"`
	ReturnedAt string  `json:"returned_at"`
}

type ReturnBookRequest struct {
	BorrowID uint `json:"borrow_id" validate:"required"`
}
