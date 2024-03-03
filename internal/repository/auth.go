package repository

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

type AuthRepository interface {
	VerifyHeader(authHeader string, ctx context.Context) (*auth.Token, error)
	GetLoginInformation(uid string, ctx context.Context) (*auth.UserRecord, error)
}

func NewAuthRepository(repository *Repository) AuthRepository {

	auth, err := repository.firebaseApp.Auth(context.Background())

	if err != nil {
		repository.logger.Sugar().Panicf("could not create auth repository: %v", err)
	}

	return &authRepository{
		Repository: repository,
		Client:     auth,
	}
}

type authRepository struct {
	*Repository
	Client *auth.Client
}

func (r *authRepository) VerifyHeader(
	authHeader string,
	ctx context.Context,
) (*auth.Token, error) {
	token, err := r.Client.VerifyIDToken(ctx, authHeader)
	if err != nil {
		return nil, fmt.Errorf("could not verify the token properly: %v", err)
	}
	return token, err
}

func (r *authRepository) GetLoginInformation(
	uid string,
	ctx context.Context,
) (*auth.UserRecord, error) {

	userData, err := r.Client.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("could not grab the user from firestore db: %v", err)
	}
	return userData, nil
}
