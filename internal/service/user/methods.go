package user

import (
	"context"
	"log"

	"github.com/arifinhermawan/amartha-billing-engine/internal/repository/pgsql"
)

func (svc *Service) CreateUser(ctx context.Context, name string, hashedPassword string) error {
	timeNow := svc.lib.GetTimeGMT7()
	err := svc.db.CreateUserInDB(ctx, pgsql.CreateUserReq{
		Name:      name,
		Password:  hashedPassword,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	})
	if err != nil {
		log.Printf("[CreateUser] svc.db.CreateUserInDB() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"name": name})
		return err
	}

	return nil
}

func (svc *Service) GetUserByID(ctx context.Context, userID int64) (User, error) {
	res, err := svc.db.GetUserByIDFromDB(ctx, userID)
	if err != nil {
		log.Printf("[GetUserByID] svc.db.GetUserByIDFromDB() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"user_id": userID})
		return User{}, err
	}

	return User(res), nil
}
