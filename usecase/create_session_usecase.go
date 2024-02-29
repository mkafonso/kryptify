package usecase

import (
	"context"
	"kryptify/entity"
	"kryptify/repository"
	"kryptify/token"

	appError "kryptify/usecase/error"
	"time"
)

type CreateSessionRequest struct {
	Email, Password     string
	UserAgent, ClientIP string
}

type CreateSessionResponse struct {
	Account               *entity.Account
	SessionID             string
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type CreateSession struct {
	accountRepo repository.AccountsRepositoryInterface
	sessionRepo repository.SessionsRepositoryInterface
}

func NewCreateSession(
	accountRepo repository.AccountsRepositoryInterface,
	sessionRepo repository.SessionsRepositoryInterface,
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
	accessToken, accessTokenPayload, err := token.GenerateAccessToken(account.ID)
	if err != nil {
		return nil, err
	}

	// generate refresh token
	refreshToken, refreshTokenPayload, err := token.GenerateRefreshToken(account.ID)
	if err != nil {
		return nil, err
	}

	// save the new session
	_, err = c.sessionRepo.CreateSession(ctx, &entity.Session{
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
