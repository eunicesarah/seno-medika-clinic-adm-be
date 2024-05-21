package ttv

import (
	"strconv"
	"sync"

	"seno-medika.com/config/db"
	"seno-medika.com/model/doctorstation"
	"seno-medika.com/model/nursestation"
	doctorstation2 "seno-medika.com/model/station/doctorstation"
	"seno-medika.com/model/station/nursestation"
	"strconv"
	"sync"
)

func FindSkriningAwalById(
	id string, errorChan chan error,
	wg *sync.WaitGroup, skriningAwalRes *nursestation.SkriningAwal,
) {
	defer wg.Done()
	skrinAwalId := db.DB.QueryRow("SELECT skrin_awal_id FROM anamnesis WHERE pasien_id = $1", id)

	if err := db.DB.QueryRow("SELECT * FROM skrining_awal WHERE skrin_awal_id = $1", skrinAwalId).Scan(
		&skriningAwalRes.SkriningAwalID,
		&skriningAwalRes.Disabilitas,
		&skriningAwalRes.Ambulansi,
		&skriningAwalRes.HambatanKomunikasi,
		&skriningAwalRes.JalanTidakSeimbang,
		&skriningAwalRes.JalanAlatBantu,
		&skriningAwalRes.MenopangSaatDuduk,
		&skriningAwalRes.HasilCaraJalan,
		&skriningAwalRes.SkalaNyeri,
		&skriningAwalRes.NyeriBerulang,
		&skriningAwalRes.SifatNyeri,
	); err != nil {
		errorChan <- err
		return
	}
}

func FindSkriningGiziById(
	id string, errorChan chan error,
	wg *sync.WaitGroup, skriningGiziRes *nursestation.SkriningGizi,
) {
	defer wg.Done()
	skrinGiziId := db.DB.QueryRow("SELECT skrin_gizi_id FROM anamnesis WHERE pasien_id = $1", id)

	if err := db.DB.QueryRow("SELECT * FROM skrining_gizi WHERE skrin_gizi_id = $1", skrinGiziId).Scan(
		&skriningGiziRes.SkriningGiziID,
		&skriningGiziRes.PenurunanBB,
		&skriningGiziRes.TdkNafsuMakan,
		&skriningGiziRes.DiagnosisKhusus,
		&skriningGiziRes.NamaPenyakit,
	); err != nil {
		errorChan <- err
		return
	}

	return
}

func FindTTVById(
	id string, errorChan chan error,
	wg *sync.WaitGroup, ttvRes *nursestation.TTV,
) {
	defer wg.Done()
	ttvId := db.DB.QueryRow("SELECT ttv_id FROM anamnesis WHERE pasien_id = $1", id)

	if err := db.DB.QueryRow("SELECT * FROM ttv WHERE ttv_id = $1", ttvId).Scan(
		&ttvRes.TTVID,
		&ttvRes.Kesadaran,
		&ttvRes.Sistole,
		&ttvRes.Diastole,
		&ttvRes.TinggiBadan,
		&ttvRes.CaraUkurTB,
		&ttvRes.BeratBadan,
		&ttvRes.LingkarPerut,
		&ttvRes.DetakNadi,
		&ttvRes.Nafas,
		&ttvRes.Saturasi,
		&ttvRes.Suhu,
		&ttvRes.DetakJantung,
		&ttvRes.Triage,
		&ttvRes.PsikolososialSpirit,
		&ttvRes.Keterangan,
	); err != nil {
		errorChan <- err
		return
	}
}

func FindRiwayatPenyakitById(
	id string, errorChan chan error,
	wg *sync.WaitGroup, riwayatPenyakitRes *nursestation.RiwayatPenyakit,
) {
	defer wg.Done()
	riwayatPenyakitId := db.DB.QueryRow("SELECT riwayat_penyakit_id FROM anamnesis WHERE pasien_id = $1", id)

	if err := db.DB.QueryRow("SELECT * FROM riwayat_penyakit WHERE riwayat_penyakit_id = $1", riwayatPenyakitId).Scan(
		&riwayatPenyakitRes.RiwayatPenyakitID,
		&riwayatPenyakitRes.RPS,
		&riwayatPenyakitRes.RPD,
		&riwayatPenyakitRes.RPK,
	); err != nil {
		errorChan <- err
		return
	}
}

func FindAnamnesisById(
	id string, errorChan *error,
	riwayatPenyakitRes *doctorstation.Anamnesis,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	if err := db.DB.QueryRow("SELECT * FROM anamnesis WHERE anamnesis_id = $1", id).Scan(
		&riwayatPenyakitRes.AnamnesisID,
		&riwayatPenyakitRes.PasienID,
		&riwayatPenyakitRes.SkrinAwalID,
		&riwayatPenyakitRes.SkrinGiziID,
		&riwayatPenyakitRes.TTVID,
		&riwayatPenyakitRes.RiwayatPenyakitID,
		&riwayatPenyakitRes.AlergiID,
		&riwayatPenyakitRes.DokterID,
		&riwayatPenyakitRes.PerawatID,
		&riwayatPenyakitRes.KeluhanUtama,
		&riwayatPenyakitRes.KeluhanTambahan,
		&riwayatPenyakitRes.LamaSakit,
	); err != nil {
		*errorChan = err
		return
	}

	return
}

func FindAnamnesisByPasienId(
	id string, errorChan *error,
	riwayatPenyakitRes *doctorstation.Anamnesis,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	if err := db.DB.QueryRow("SELECT * FROM anamnesis WHERE pasien_id = $1", id).Scan(
		&riwayatPenyakitRes.AnamnesisID,
		&riwayatPenyakitRes.PasienID,
		&riwayatPenyakitRes.SkrinAwalID,
		&riwayatPenyakitRes.SkrinGiziID,
		&riwayatPenyakitRes.TTVID,
		&riwayatPenyakitRes.RiwayatPenyakitID,
		&riwayatPenyakitRes.AlergiID,
		&riwayatPenyakitRes.DokterID,
		&riwayatPenyakitRes.PerawatID,
		&riwayatPenyakitRes.KeluhanUtama,
		&riwayatPenyakitRes.KeluhanTambahan,
		&riwayatPenyakitRes.LamaSakit,
	); err != nil {
		*errorChan = err
		return
	}

	return
}

func FindAlergiById(
	id string, errorChan chan error,
	wg *sync.WaitGroup, alergiRes *doctorstation2.Alergi,
) {
	defer wg.Done()

	if err := db.DB.QueryRow("SELECT * FROM alergi WHERE alergi_id = $1", id).Scan(
		&alergiRes.AlergiID,
		&alergiRes.Obat,
		&alergiRes.Makanan,
		&alergiRes.Lainnya,
	); err != nil {
		errorChan <- err
		return
	}
}

func FindAllNurseStation() (nurseStationRes []nursestation.NurseStation, err error) {
	var anamnesisRes []doctorstation2.Anamnesis
	var wg sync.WaitGroup
	errChan := make(chan error, 1)

	wg.Add(1)

	go func() {
		defer wg.Done()
		rows, err := db.DB.Query("SELECT * FROM anamnesis")
		if err != nil {
			return
		}
		// Close the rows after the function ends
		defer rows.Close()

		for rows.Next() {
			var anamnesis doctorstation2.Anamnesis
			if err := rows.Scan(
				&anamnesis.AnamnesisID,
				&anamnesis.PasienID,
				&anamnesis.SkrinAwalID,
				&anamnesis.SkrinGiziID,
				&anamnesis.TTVID,
				&anamnesis.RiwayatPenyakitID,
				&anamnesis.AlergiID,
				&anamnesis.DokterID,
				&anamnesis.PerawatID,
				&anamnesis.KeluhanUtama,
				&anamnesis.KeluhanTambahan,
				&anamnesis.LamaSakit,
			); err != nil {
				errChan <- err
				return
			}
			anamnesisRes = append(anamnesisRes, anamnesis)
		}
	}()

	wg.Wait()
	close(errChan)
	if errChan != nil {
		err = <-errChan
		return nil, err

	}

	for _, anamnesis := range anamnesisRes {
		var nurseStation nursestation.NurseStation
		nurseStation.Anamnesis = anamnesis

		wg.Add(5)
		errChan := make(chan error, 5)

		go FindAlergiById(strconv.Itoa(anamnesis.AlergiID), errChan, &wg, &nurseStation.Alergi)
		go FindSkriningAwalById(strconv.Itoa(anamnesis.SkrinAwalID), errChan, &wg, &nurseStation.SkriningAwal)
		go FindSkriningGiziById(strconv.Itoa(anamnesis.SkrinGiziID), errChan, &wg, &nurseStation.SkriningGizi)
		go FindTTVById(strconv.Itoa(anamnesis.TTVID), errChan, &wg, &nurseStation.TTV)
		go FindRiwayatPenyakitById(strconv.Itoa(anamnesis.RiwayatPenyakitID), errChan, &wg, &nurseStation.RiwayatPenyakit)

		wg.Wait()
		close(errChan)
		if errChan != nil {
			err = <-errChan
			return nil, err
		}

		nurseStationRes = append(nurseStationRes, nurseStation)
	}

	return nurseStationRes, nil
}

func FindNurseStationById(id string) (nurseStationRes nursestation.NurseStation, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go FindAnamnesisById(id, &err, &nurseStationRes.Anamnesis, &wg)
	wg.Wait()

	if err != nil {
		return nursestation.NurseStation{}, err
	}

	errorChan := make(chan error, 5)
	wg.Add(5)
	go FindAlergiById(id, errorChan, &wg, &nurseStationRes.Alergi)
	go FindSkriningAwalById(id, errorChan, &wg, &nurseStationRes.SkriningAwal)
	go FindSkriningGiziById(id, errorChan, &wg, &nurseStationRes.SkriningGizi)
	go FindTTVById(id, errorChan, &wg, &nurseStationRes.TTV)
	go FindRiwayatPenyakitById(id, errorChan, &wg, &nurseStationRes.RiwayatPenyakit)

	wg.Wait()
	close(errorChan)

	for err := range errorChan {
		if err != nil {
			return nursestation.NurseStation{}, err
		}
	}

	return nurseStationRes, nil
}
func FindNurseStationByPasienId(id string) (nurseStationRes nursestation.NurseStation, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	var anamnesisID string;
	go FindAnamnesisByPasienId(id, &err, &nurseStationRes.Anamnesis, &wg)
	err = db.DB.QueryRow("SELECT anamnesis_id FROM anamnesis WHERE pasien_id = $1", id).Scan(&anamnesisID)

	wg.Wait()

	if err != nil {
		return nursestation.NurseStation{}, err
	}

	errorChan := make(chan error, 5)
	wg.Add(5)
	go FindAlergiById(anamnesisID, errorChan, &wg, &nurseStationRes.Alergi)
	go FindSkriningAwalById(anamnesisID, errorChan, &wg, &nurseStationRes.SkriningAwal)
	go FindSkriningGiziById(anamnesisID, errorChan, &wg, &nurseStationRes.SkriningGizi)
	go FindTTVById(anamnesisID, errorChan, &wg, &nurseStationRes.TTV)
	go FindRiwayatPenyakitById(anamnesisID, errorChan, &wg, &nurseStationRes.RiwayatPenyakit)

	wg.Wait()
	close(errorChan)

	for err := range errorChan {
		if err != nil {
			return nursestation.NurseStation{}, err
		}
	}

	return nurseStationRes, nil
}
