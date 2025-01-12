package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
)

func InitDBConn(cfg configuration.DatabaseConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DatabaseName)

	db, err := sqlx.Open(cfg.Driver, psqlInfo)
	if err != nil {
		log.Printf("[InitDBConn] sql.Open() got error: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("[InitDBConn] db.Ping() got error: %v\n", err)
		return nil, err
	}

	return db, nil
}
