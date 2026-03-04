package repository

import (
	"context"
	"database/sql"
	"goatbrand-backend/internal/domain"
)

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) domain.UserRepository {
	return &userPostgres{db: db}
}

func (r *userPostgres) FindAll() ([]domain.User, error) {
	query := `SELECT id, name, email, phone, code, payment_status, created_at FROM users`
	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.Code, &u.PaymentStatus, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
