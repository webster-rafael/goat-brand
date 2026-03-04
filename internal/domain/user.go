package domain

import "time"

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Code          string    `json:"code"`
	PaymentStatus string    `json:"paymentStatus"`
	CreatedAt     time.Time `json:"createdAt"`
}

type UserRepository interface {
	FindAll() ([]User, error)
}
