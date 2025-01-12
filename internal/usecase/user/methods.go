package user

import (
	"context"
	"log"
)

func (uc *UseCase) CreateUser(ctx context.Context, name string, password string) error {
	err := uc.user.CreateUser(ctx, name, uc.auth.EncryptPassword(password))
	if err != nil {
		log.Printf("[CreateUser] uc.user.CreateUser() got error: %v\nMetadata: %v\n", err, map[string]interface{}{"name": name})
		return err
	}

	return nil
}
