package user

import (
	"context"
	"log"
	"time"
)

func (uc *UseCase) CreateUser(ctx context.Context, name string, password string) error {
	err := uc.user.CreateUser(ctx, name, uc.auth.EncryptPassword(password))
	if err != nil {
		log.Printf("[CreateUser] uc.user.CreateUser() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"name": name})
		return err
	}

	return nil
}

func (uc *UseCase) GetDelinquentsUsers(ctx context.Context, date time.Time) ([]User, error) {
	currentTime := uc.lib.GetTimeGMT7()
	if !date.IsZero() {
		currentTime = date
	}

	user, err := uc.user.GetDelinquentUsers(ctx, currentTime)
	if err != nil {
		log.Printf("[GetDelinquentsUsers] uc.user.GetDelinquentUsers() got error: %v\n", err)
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
