package dokter

import (
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/doctorstation"
	"time"
)

func PatchPemeriksaanFisik(antrianId int, pemeriksaanFisik doctorstation.PemeriksaanFisik) error {
	var (
		pemeriksaanDokterId int
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanDokterId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("UPDATE pemeriksaan_fisik SET terapi_yg_sdh_dilakukan = $1, rencana_tindakan = $2, tindakan_keperawatan = $3, observasi = $4, merokok = $5, konsumsi_alkohol = $6, kurang_sayur = $7 WHERE pemeriksaan_dokter_id = $8",
		pemeriksaanFisik.TerapiYgSdhDilakukan, pemeriksaanFisik.RencanaTindakan, pemeriksaanFisik.TindakanKeperawatan, pemeriksaanFisik.Observasi, pemeriksaanFisik.Merokok, pemeriksaanFisik.KonsumsiAlkohol, pemeriksaanFisik.KurangSayur, pemeriksaanDokterId); err != nil {
		return err
	}

	return nil
}

func PatchRiwayatPemeriksaan(antrianId int, riwayatPemeriksaan doctorstation.RiwayatPemeriksaan) error {
	var (
		pemeriksaanDokterId int
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanDokterId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("UPDATE riwayat_pemeriksaan SET tanggal = $1, pemeriksaan = $2, keterangan = $3 WHERE pemeriksaan_dokter_id = $4",
		time.Now().Format("2006-01-02"), riwayatPemeriksaan.Pemeriksaan, riwayatPemeriksaan.Keterangan, pemeriksaanDokterId); err != nil {
		return err
	}

	return nil
}

func PatchKeadaanFisik(antrianId int, keadaanFisik doctorstation.KeadaanFisik) error {
	var (
		pemeriksaanDokterId int
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanDokterId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("UPDATE keadaan_fisik SET pemeriksaan_kulit = $1, pemeriksaan_kuku = $2, pemeriksaan_kepala = $3, pemeriksaan_mata = $4, pemeriksaan_telinga = $5, pemeriksaan_hidung_sinus = $6, pemeriksaan_mulut_bibir = $7, pemeriksaan_leher = $8, pemeriksaan_dada_punggung = $9, pemeriksaan_kardiovaskuler = $10, pemeriksaan_abdomen_perut = $11, pemeriksaan_ekstremitas_atas = $12, pemeriksaan_ekstremitas_bawah = $13, pemeriksaan_genitalia_pria = $14 WHERE pemeriksaan_dokter_id = $15",
		keadaanFisik.PemeriksaanKulit, keadaanFisik.PemeriksaanKuku, keadaanFisik.PemeriksaanKepala, keadaanFisik.PemeriksaanMata, keadaanFisik.PemeriksaanTelinga, keadaanFisik.PemeriksaanHidungSinus, keadaanFisik.PemeriksaanMulutBibir, keadaanFisik.PemeriksaanLeher, keadaanFisik.PemeriksaanDadaPunggung, keadaanFisik.PemeriksaanKardiovaskuler, keadaanFisik.PemeriksaanAbdomenPerut, keadaanFisik.PemeriksaanEkstremitasAtas, keadaanFisik.PemeriksaanEkstremitasBawah, keadaanFisik.PemeriksaanGenitaliaPria, pemeriksaanDokterId); err != nil {
		return err
	}

	return nil
}

func PatchDiagnosa(antrianId int, diagnosa doctorstation.Diagnosa) error {
	var (
		pemeriksaanDokterId int
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanDokterId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("UPDATE diagnosa SET diagnosa = $1, jenis = $2, kasus = $3, status_diagnosis = $4 WHERE pemeriksaan_dokter_id = $5",
		diagnosa.Diagnosa, diagnosa.Jenis, diagnosa.Kasus, diagnosa.StatusDiagnosis, pemeriksaanDokterId); err != nil {
		return err
	}

	return nil

}

func PatchAnatomi(antrianId int, anatomi doctorstation.Anatomi) error {
	var (
		pemeriksaanDokterId int
	)

	if err := db.DB.QueryRow("SELECT pemeriksaan_dokter_id FROM pemeriksaan_dokter WHERE antrian_id = $1", antrianId).Scan(&pemeriksaanDokterId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("UPDATE anatomi SET bagian_tubuh = $1, keterangan = $2 WHERE pemeriksaan_dokter_id = $3",
		anatomi.BagianTubuh, anatomi.Keterangan, pemeriksaanDokterId); err != nil {
		return err
	}

	return nil
}
