package data

import "time"

type Wallet struct {
	ID        uint      `db:"id"`
	UserID    uint      `db:"user_id"`
	Balance   float64   `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (w *Wallet) GetWalletByUserID(userID uint) (*Wallet, error) {
	sqlStatement := `SELECT id, user_id, balance, created_at, updated_at FROM wallets WHERE user_id = $1`

	var wallet Wallet
	err := db.Get(&wallet, sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (w *Wallet) GetWalletIdByUserID(userID uint) (uint, error) {
	sqlStatement := `SELECT id FROM wallets WHERE user_id = $1`

	var id uint
	err := db.Get(&id, sqlStatement, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
