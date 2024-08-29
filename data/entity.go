package data

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func New(dbPool *sqlx.DB) *Models {
	db = dbPool

	return &Models{
		User: &User{},
	}
}

type Models struct {
	User UserInterfaces
}
