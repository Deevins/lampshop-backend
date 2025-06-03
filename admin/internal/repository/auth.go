// internal/repository/auth_repository.go

package repository

import (
	"context"
	"errors"
	"github.com/Deevins/lampshop-backend/admin/internal/repository/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository interface {
	VerifyUser(ctx context.Context, username, password string) error
}

type authRepoPG struct {
	queries *sql.Queries
}

func NewAuthRepository(pool *pgxpool.Pool) AuthRepository {
	return &authRepoPG{
		queries: sql.New(pool),
	}
}

func (r *authRepoPG) VerifyUser(ctx context.Context, username, password string) error {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return errors.New("invalid username or password")
	}
	return nil
}
