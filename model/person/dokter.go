package person

import "github.com/google/uuid"

type Dokter struct {
	UserID     int        `json:"user_id"`
	UserUUID   uuid.UUID  `json:"user_uuid"`
	Nama       string     `json:"nama"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Role       string     `json:"role"`
	DokterData DokterData `json:"dokter_data"`
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
