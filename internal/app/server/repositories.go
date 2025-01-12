package server

import (
	"github.com/arifinhermawan/amartha-billing-engine/internal/lib"
	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	DB *pgsql.Repository
}

func NewRepositories(lib *lib.Lib, db *sqlx.DB) *Repositories {
	return &Repositories{
		DB: pgsql.NewRepository(lib, db),
	}
}
