package person

import "github.com/google/uuid"

type Perawat struct {
	UserID      int         `json:"user_id"`
	UserUUID    uuid.UUID   `json:"user_uuid"`
	Nama        string      `json:"nama"`
	Password    string      `json:"password"`
	Email       string      `json:"email"`
	Role        string      `json:"role"`
	PerawatData PerawatData `json:"perawat_data"`
}

type PerawatData struct {
	PerawatID    int    `json:"perawat_id"`
	NomorLisensi string `json:"nomor_lisensi"`
}
