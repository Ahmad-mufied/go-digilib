package data

import (
	"github.com/Ahmad-mufied/go-digilib/constants"
	"time"
)

// Deposit represents the deposits table in the database.
type Deposit struct {
	ID         uint                        `db:"id"`
	Amount     float64                     `db:"amount"`
	WalletID   uint                        `db:"wallet_id"`
	InvoiceURL string                      `db:"invoice_url"`
	Status     constants.PaymentStatusEnum `db:"status"`
	CreatedAt  time.Time                   `db:"created_at"`
	UpdatedAt  time.Time                   `db:"updated_at"`
}

func (d *Deposit) CreateDeposit(deposit *Deposit) (uint, error) {

	tx := db.MustBegin()

	sqlStatement := `INSERT INTO deposits (amount, wallet_id, invoice_url, status) VALUES ($1, $2, $3, $4) RETURNING id`

	// Insert deposit data to database
	var lastInsertID uint
	err := tx.QueryRow(sqlStatement, deposit.Amount, deposit.WalletID, deposit.InvoiceURL, deposit.Status).Scan(&lastInsertID)
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

func (d *Deposit) GetAllDepositsByWalletID(walletID uint) ([]Deposit, error) {
	sqlStatement := `SELECT id, amount, wallet_id, invoice_url, status, created_at, updated_at FROM deposits WHERE wallet_id = $1`

	var deposits []Deposit
	err := db.Select(&deposits, sqlStatement, walletID)
	if err != nil {
		return nil, err
	}

	return deposits, nil
}

func (d *Deposit) GetDepositById(depositID uint) (*Deposit, error) {
	sqlStatement := `SELECT id, amount, wallet_id, invoice_url, status, created_at, updated_at FROM deposits WHERE id = $1`

	var deposit Deposit
	err := db.Get(&deposit, sqlStatement, depositID)
	if err != nil {
		return nil, err
	}

	return &deposit, nil
}

func (d *Deposit) UpdateDepositStatus(invoiceUrl string, status constants.PaymentStatusEnum) error {

	tx := db.MustBegin()

	sqlStatement := `SELECT id, amount, wallet_id, invoice_url, status, created_at, updated_at FROM deposits WHERE invoice_url = $1`

	var deposit Deposit
	err := tx.Get(&deposit, sqlStatement, invoiceUrl)
	if err != nil {
		tx.Rollback()
		return err
	}

	if status == constants.PaymentStatusConfirmed {
		// Update wallet balance
		result, err := tx.Exec(`UPDATE wallets SET balance = balance + $1 WHERE id = $2`, deposit.Amount, deposit.WalletID)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Check if wallet balance updated
		if affected, _ := result.RowsAffected(); affected == 0 {
			tx.Rollback()
			return err
		}
	}

	sqlStatement = `UPDATE deposits SET status = $1 WHERE invoice_url = $2`

	result, err := tx.Exec(sqlStatement, status, invoiceUrl)
	if err != nil {
		return err
	}

	// Check if deposit status updated
	if affected, _ := result.RowsAffected(); affected == 0 {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return nil
}

func (d *Deposit) UpdateDepositInvoiceURL(depositID uint, invoiceURL string) error {
	sqlStatement := `UPDATE deposits SET invoice_url = $1 WHERE id = $2`

	_, err := db.Exec(sqlStatement, invoiceURL, depositID)
	if err != nil {
		return err
	}

	return nil
}

func (d *Deposit) GetDepositByInvoiceURL(invoiceURL string) (*Deposit, error) {
	sqlStatement := `SELECT id, amount, wallet_id, invoice_url, status, created_at, updated_at FROM deposits WHERE invoice_url = $1`

	var deposit Deposit
	err := db.Get(&deposit, sqlStatement, invoiceURL)
	if err != nil {
		return nil, err
	}

	return &deposit, nil
}
