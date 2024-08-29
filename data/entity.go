package data

import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func New(dbPool *sqlx.DB) *Models {
	db = dbPool

	return &Models{
		User: &User{},
		Book: &Book{},
	}
}

type Models struct {
	User UserInterfaces
	Book BookInterfaces
}
