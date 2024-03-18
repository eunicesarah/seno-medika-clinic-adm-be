package person

import "github.com/google/uuid"

type Dokter struct {
	UserID     int        `json:"user_id" binding:"required"`
	UserUUID   uuid.UUID  `json:"user_uuid" binding:"required"`
	Nama       string     `json:"nama" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	Email      string     `json:"email" binding:"required"`
	Role       string     `json:"role" binding:"required"`
	DokterData DokterData `json:"dokter_data" binding:"required"`
}

type DokterData struct {
	DokterID         int                `json:"dokter_id"`
	JagaPoliMana     string             `json:"jaga_poli_mana"`
	JadwalJaga       string             `json:"jadwal_jaga"`
	NomorLisensi     string             `json:"nomor_lisensi"`
	ListJadwalDokter []ListJadwalDokter `json:"list_jadwal_dokter"`
}

type ListJadwalDokter struct {
	DokterID int    `json:"dokter_id"`
	Hari     string `json:"hari"`
	Shift    int    `json:"shift"`
}
