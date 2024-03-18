package person

import "github.com/google/uuid"

type Apoteker struct {
	UserID       int          `json:"user_id"`
	UserUUID     uuid.UUID    `json:"user_uuid"`
	Nama         string       `json:"nama"`
	Password     string       `json:"password"`
	Email        string       `json:"email"`
	Role         string       `json:"role"`
	ApotekerData ApotekerData `json:"apoteker_data"`
}

type ApotekerData struct {
	ApotekerID   int    `json:"apoteker_id"`
	NomorLisensi string `json:"nomor_lisensi"`
}
