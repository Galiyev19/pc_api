package repository

import (
	"database/sql"
	"errors"
	"user-auth-service/internal/models"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) Login(email string) (*models.User, error) {
	var user models.User

	query := `SELECT FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)
	err := row.Scan(&user.ID, &user.Email, &user.HashPassword, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.IsActive)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
