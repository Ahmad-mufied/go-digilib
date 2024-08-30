package data

import (
	"database/sql"
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/pkg/errors"
	"time"
)

type Borrow struct {
	ID         uint                     `db:"id"`
	UserID     uint                     `db:"user_id"`
	BookID     uint                     `db:"book_id"`
	Status     constants.BookStatusEnum `db:"status"`
	StartDate  time.Time                `db:"start_date"`
	EndDate    time.Time                `db:"end_date"`
	TotalPrice float64                  `db:"total_price"`
	ReturnedAt sql.NullTime             `db:"returned_at"`
}

func (b *Borrow) CreateBorrow(borrow *Borrow) (uint, error) {
	tx := db.MustBegin()

	sqlStatement := `INSERT INTO borrows (user_id, book_id, status, start_date, end_date, total_price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var lastInsertID uint
	err := tx.QueryRow(sqlStatement, borrow.UserID, borrow.BookID, borrow.Status, borrow.StartDate, borrow.EndDate, borrow.TotalPrice).Scan(&lastInsertID)
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

func (b *Borrow) GetBorrowById(borrowID uint) (*Borrow, error) {
	sqlStatement := `SELECT id, user_id, book_id, status, start_date, end_date, total_price, returned_at FROM borrows WHERE id = $1`

	var borrow Borrow
	err := db.Get(&borrow, sqlStatement, borrowID)
	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (b *Borrow) GetAllBorrowsByUserID(userID uint) ([]Borrow, error) {
	sqlStatement := `SELECT id, user_id, book_id, status, start_date, end_date, total_price, returned_at FROM borrows WHERE user_id = $1`

	var borrows []Borrow
	err := db.Select(&borrows, sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	return borrows, nil
}

func (b *Borrow) UpdateBorrowReturnedAt(borrowID uint) error {
	tx := db.MustBegin()

	// update returned_at and status to returned
	sqlStatement := `UPDATE borrows SET returned_at = $1, status = $2 WHERE id = $3`

	result, err := tx.Exec(sqlStatement, time.Now(), constants.BorrowStatusReturned, borrowID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		tx.Rollback()
		return errors.New("borrow not found")
	}

	// update book stock
	sqlStatement = `UPDATE books SET stock = stock + 1 WHERE id = (SELECT book_id FROM borrows WHERE id = $1)`
	_, err = tx.Exec(sqlStatement, borrowID)
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

func (b *Borrow) UpdateBorrowStatusConfirm(borrowID uint, status constants.BookStatusEnum) error {
	tx := db.MustBegin()

	sqlStatement := `UPDATE borrows SET status = $1 WHERE id = $2`

	_, err := tx.Exec(sqlStatement, status, borrowID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// deduct user wallet balance
	sqlStatement = `UPDATE wallets SET balance = balance - (SELECT total_price FROM borrows WHERE id = $1) WHERE user_id = (SELECT user_id FROM borrows WHERE id = $1)`
	result, err := tx.Exec(sqlStatement, borrowID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	// update book stock
	sqlStatement = `UPDATE books SET stock = stock - 1 WHERE id = (SELECT book_id FROM borrows WHERE id = $1)`
	_, err = tx.Exec(sqlStatement, borrowID)
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
