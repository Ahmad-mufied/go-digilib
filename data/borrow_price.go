package data

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"time"
)

type BorrowPrice struct {
	ID              uint                   `db:"id"`
	BookID          uint                   `db:"book_id"`
	DurationType    constants.DurationType `db:"duration_type"`
	PriceMultiplier float64                `db:"price_multiplier"`
	CreatedAt       time.Time              `db:"created_at"`
	UpdatedAt       time.Time              `db:"updated_at"`
}

func (bp *BorrowPrice) CreateBorrowPrice(borrowPrice *BorrowPrice) (int, error) {
	tx := db.MustBegin()

	sqlStatement := `INSERT INTO borrow_prices (book_id, duration_type, price_multiplier) VALUES ($1, $2, $3) RETURNING id`

	var lastInsertID int
	err := tx.QueryRow(sqlStatement, borrowPrice.BookID, borrowPrice.DurationType, borrowPrice.PriceMultiplier).Scan(&lastInsertID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (bp *BorrowPrice) GetBorrowPrice(bookID uint, durationType constants.DurationType) (*BorrowPrice, error) {

	sqlStatement := `SELECT id, book_id, duration_type, price_multiplier, created_at, updated_at FROM borrow_prices WHERE book_id = $1 AND duration_type = $2`

	var borrowPrice BorrowPrice
	err := db.Get(&borrowPrice, sqlStatement, bookID, durationType)
	if err != nil {
		return &BorrowPrice{}, err
	}

	return &borrowPrice, nil
}

func (bp *BorrowPrice) UpdateBorrowPrice(borrowPrice *BorrowPrice) error {
	tx := db.MustBegin()

	sqlStatement := `UPDATE borrow_prices SET price_multiplier = $1 WHERE book_id = $2`

	_, err := tx.Exec(sqlStatement, borrowPrice.PriceMultiplier, borrowPrice.BookID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
