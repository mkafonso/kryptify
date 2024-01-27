package store

import (
	"context"
	"database/sql"
	db "kryptify/db/sqlc"
	"kryptify/entities"
	valueobjects "kryptify/entities/value-objects"
	"kryptify/util"

	"github.com/google/uuid"
)

func (r *PostgresRepository) CreateAccount(ctx context.Context, account *entities.Account) error {
	params := db.CreateAccountParams{
		Name:         account.Name,
		Email:        account.Email,
		PasswordHash: string(account.PasswordHash),
	}

	err := r.Queries.CreateAccount(ctx, params)
	return err
}

func (r *PostgresRepository) UpdateAccount(ctx context.Context, email string, updatedAccount *entities.Account) error {
	params := db.UpdateAccountParams{
		Name:         updatedAccount.Name,
		PasswordHash: string(updatedAccount.PasswordHash),
		AvatarUrl:    sql.NullString{String: updatedAccount.AvatarUrl},
		UpdatedAt:    updatedAccount.UpdatedAt,
		Email:        email,
	}

	err := r.Queries.UpdateAccount(ctx, params)
	return err
}

func (r *PostgresRepository) GetAccountByID(ctx context.Context, accountID string) (*entities.Account, error) {
	account, err := r.Queries.GetAccountByID(ctx, uuid.MustParse(accountID))
	if err != nil {
		return nil, err
	}

	return &entities.Account{
		ID:                account.ID,
		Name:              account.Name,
		Email:             account.Email,
		AvatarUrl:         util.GetStringValue(account.AvatarUrl),
		IsAccountVerified: account.IsAccountVerified.Valid && account.IsAccountVerified.Bool,
		PasswordHash:      valueobjects.Password(account.PasswordHash),
		CreatedAt:         account.CreatedAt,
		UpdatedAt:         account.UpdatedAt,
	}, nil
}

func (r *PostgresRepository) FindAccountByEmail(ctx context.Context, email string) (*entities.Account, error) {
	account, err := r.Queries.FindAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &entities.Account{
		ID:                account.ID,
		Name:              account.Name,
		Email:             account.Email,
		AvatarUrl:         util.GetStringValue(account.AvatarUrl),
		IsAccountVerified: account.IsAccountVerified.Valid && account.IsAccountVerified.Bool,
		PasswordHash:      valueobjects.Password(account.PasswordHash),
		CreatedAt:         account.CreatedAt,
		UpdatedAt:         account.UpdatedAt,
	}, nil
}
