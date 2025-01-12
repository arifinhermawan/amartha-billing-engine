package user

import (
	"context"
	"time"

	"github.com/arifinhermawan/amartha-billing-engine/internal/lib/configuration"
	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	CreateUserInDB(ctx context.Context, req pgsql.CreateUserReq) error
	GetDelinquentsUsersFromDB(ctx context.Context, date time.Time) ([]pgsql.User, error)
	GetUserByIDFromDB(ctx context.Context, userID int64) (pgsql.User, error)
}

type Service struct {
	lib libProvider
	db  dbProvider
}

func NewService(lib libProvider, db dbProvider) *Service {
	return &Service{
		lib: lib,
		db:  db,
	}
}
