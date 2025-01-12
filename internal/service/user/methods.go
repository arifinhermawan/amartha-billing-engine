package user

import (
	"context"
	"log"
	"time"

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

func (svc *Service) GetDelinquentUsers(ctx context.Context, date time.Time) ([]User, error) {
	user, err := svc.db.GetDelinquentsUsersFromDB(ctx, date)
	if err != nil {
		log.Printf("[GetDelinquentUsers] svc.db.GetDelinquentsUsersFromDB() got error: %v\n", err)
		return nil, err
	}

	result := make([]User, len(user))
	for idx, val := range user {
		result[idx] = User{
			ID:   val.ID,
			Name: val.Name,
		}
	}

	return result, nil
}

func (svc *Service) GetUserByID(ctx context.Context, userID int64) (User, error) {
	res, err := svc.db.GetUserByIDFromDB(ctx, userID)
	if err != nil {
		log.Printf("[GetUserByID] svc.db.GetUserByIDFromDB() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"user_id": userID})
		return User{}, err
	}

	return User(res), nil
}
