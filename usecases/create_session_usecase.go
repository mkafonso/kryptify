package usecases

import (
	"context"
	"kryptify/entities"
	"kryptify/repositories"
	"kryptify/utils"

	appError "kryptify/usecases/errors"
	"time"
)

type CreateSessionRequest struct {
	Email, Password     string
	UserAgent, ClientIP string
}

type CreateSessionResponse struct {
	Account               *entities.Account
	SessionID             string
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type CreateSession struct {
	accountRepo repositories.AccountsRepositoryInterface
	sessionRepo repositories.SessionsRepositoryInterface
}

func NewCreateSession(
	accountRepo repositories.AccountsRepositoryInterface,
	sessionRepo repositories.SessionsRepositoryInterface,
) *CreateSession {
	return &CreateSession{
		accountRepo: accountRepo,
		sessionRepo: sessionRepo,
	}
}

func (c *CreateSession) Execute(ctx context.Context, data *CreateSessionRequest) (*CreateSessionResponse, error) {
	// check if email exists
	account, err := c.accountRepo.FindAccountByEmail(ctx, data.Email)
	if err != nil {
		return nil, appError.NewErrorInvalidCredentials()
	}

	// compare the provided password with the stored hash
	if !account.PasswordHash.IsPasswordValid(data.Password) {
		return nil, appError.NewErrorInvalidCredentials()
	}

	// generate access token
	accessToken, accessTokenPayload, err := utils.GenerateAccessToken(account.ID)
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, refreshTokenPayload, err := utils.GenerateRefreshToken(account.ID)
	if err != nil {
		return nil, err
	}

	// save the new session
	_, err = c.sessionRepo.CreateSession(ctx, &entities.Session{
		ID:           refreshTokenPayload.ID,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
		AccountID:    account.ID.String(),
		RefreshToken: refreshToken,
		IsBlocked:    false,
		UserAgent:    data.UserAgent,
		ClientIP:     data.ClientIP,
		CreatedAt:    time.Now().UTC(),
	})

	if err != nil {
		return nil, err
	}

	response := &CreateSessionResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
		SessionID:             refreshTokenPayload.ID.String(),
		Account:               account,
	}

	return response, nil
}
