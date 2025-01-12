package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
)

func InitDBConn(cfg configuration.DatabaseConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DatabaseName)

	db, err := sql.Open(cfg.Driver, psqlInfo)
	if err != nil {
		log.Fatalf("[InitDBConn] sql.Open() got error: %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[InitDBConn] db.Ping() got error: %v\n", err)
		return nil, err
	}

	return db, nil
}
