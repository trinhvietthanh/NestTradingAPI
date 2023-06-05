package repository

import (
	"github.com/jmoiron/sqlx"
)

type pGUserRepository struct {
	DB *sqlx.DB
}
