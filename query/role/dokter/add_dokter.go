package dokter

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/cashierstation"
)

func AddTindakan(tindakan cashierstation.Tindakan) error {
	if _, err := db.DB.Exec(`INSERT INTO tindakan(jenis_tindakan, prosedur_tindakan, jumlah, keterangan, tanggal_rencana, harga_tindakan, indikasi_tindakan, tujuan, risiko, komplikasi, alternatif_risiko) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, tindakan.JenisTindakan, tindakan.ProsedurTindakan, tindakan.Jumlah, tindakan.Keterangan, tindakan.TanggalRencana, tindakan.HargaTindakan, tindakan.IndikasiTindakan, tindakan.Tujuan, tindakan.Risiko, tindakan.Komplikasi, tindakan.AlternatifRisiko); err != nil {
		return err
	}

	return nil
}
