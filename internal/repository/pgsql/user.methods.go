package pgsql

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func (r *Repository) CreateUserInDB(ctx context.Context, req CreateUserReq) error {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	metadata := map[string]interface{}{
		"name": req.Name,
	}

	namedQuery, args, err := sqlx.Named(queryCreateUserInDB, req)
	if err != nil {
		log.Printf("[CreateUserInDB] sqlxNamed() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	_, err = r.db.ExecContext(ctxTimeout, r.db.Rebind(namedQuery), args...)
	if err != nil {
		log.Printf("[CreateUserInDB] r.db.ExecContext() got error: %v\nMetadata: %v\n", err, metadata)
		return err
	}

	return nil
}
