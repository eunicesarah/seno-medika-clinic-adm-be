package dokter

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/cashierstation"
)

func PutTindakanById(id string, tindakan cashierstation.Tindakan) error {
	val, err := db.DB.Exec(
		`UPDATE tindakan SET
                    jenis_tindakan = $1,
                    prosedur_tindakan = $2,
                    jumlah = $3,
                    keterangan = $4,
                    tanggal_rencana = $5,
                    harga_tindakan = $6,
                    indikasi_tindakan = $7,
                    tujuan = $8,
                    risiko = $9,
                    komplikasi = $10,
                    alternatif_risiko = $11
                    WHERE tindakan_id = $12`,
		tindakan.JenisTindakan,
		tindakan.ProsedurTindakan,
		tindakan.Jumlah,
		tindakan.Keterangan,
		tindakan.TanggalRencana,
		tindakan.HargaTindakan,
		tindakan.IndikasiTindakan,
		tindakan.Tujuan,
		tindakan.Risiko,
		tindakan.Komplikasi,
		tindakan.AlternatifRisiko,
		id)

	if err != nil {
		return err
	}

	if rows, _ := val.RowsAffected(); rows == 0 {
		return errors.New("id obat not found")
	}

	return nil
}
