package person

import "github.com/google/uuid"

type User struct {
	UserID   int       `json:"user_id"`
	UserUUID uuid.UUID `json:"user_uuid"`
	Nama     string    `json:"nama"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type UserWithoutPassword struct {
	UserID   int       `json:"user_id"`
	UserUUID uuid.UUID `json:"user_uuid"`
	Nama     string    `json:"nama"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}
