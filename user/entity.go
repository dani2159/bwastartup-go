package user

import "time"

type User struct {
	ID           int
	Name         string
	Occuption    string
	Email        string
	PasswordHash string
	Avatar       string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
