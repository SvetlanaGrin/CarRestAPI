package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	RequestCarCatalog
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RequestCarCatalog: NewReqPostgres(db),
	}
}
