package pgsql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/errors"
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

func (r *Repository) GetUserByIDFromDB(ctx context.Context, userID int64) (User, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.lib.GetConfig().Database.DefaultTimeout)*time.Second)
	defer cancel()

	var user User
	err := r.db.GetContext(ctxTimeout, &user, queryGetUserByIDFromDB, userID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[GetUserByIDFromDB] r.db.GetContext() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"user_id": userID})
		return User{}, err
	}

	if user.ID == 0 {
		return user, errors.ErrNotFound
	}

	return user, nil
}
