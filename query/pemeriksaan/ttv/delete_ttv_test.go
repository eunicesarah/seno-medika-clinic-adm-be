package ttv

import (
	"testing"

	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
)

func TestDeleteSkriningAwalById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO skrining_awal (
		skrin_awal_id,
		disabilitas,
		ambulansi,
		hambatan_komunikasi,
		jalan_tidak_seimbang,
		jalan_alat_bantu,
		menopang_saat_duduk,
		hasil_cara_jalan,
		skala_nyeri,
		nyeri_berulang,
		sifat_nyeri
	) VALUES (
		8911
		true,
		'Tidak Ada',
		false,
		true,
		true,
		false,
		'Tidak Normal',
		5,
		'Tidak Ada',
		'Tidak Ada'
	);
	
	`)
	err := DeleteSkriningAwalById("8990")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteSkriningAwalById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteSkriningAwalById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteSkriningAwalById("8990")
	require.Error(t, err)
}

func TestDeleteSkriningGiziById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO skrining_gizi (
		skrin_gizi_id,
		penurunan_bb,
		tdk_nafsu_makan,
		diagnosis_khusus,
		nama_penyakit
	) VALUES (
		8911
		2,
		true,
		false,
		'Tidak Ada'
	);
	
	`)
	err := DeleteSkriningGiziById("8911")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteSkriningGiziById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteSkriningGiziById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteSkriningGiziById("8990")
	require.Error(t, err)
}

func TestDeleteSkriningPenyakitById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO  (
		skrin_awal_id,
		disabilitas,
		ambulansi,
		hambatan_komunikasi,
		jalan_tidak_seimbang,
		jalan_alat_bantu,
		menopang_saat_duduk,
		hasil_cara_jalan,
		skala_nyeri,
		nyeri_berulang,
		sifat_nyeri
	) VALUES (
		8911
		true,
		'Tidak Ada',
		false,
		true,
		true,
		false,
		'Tidak Normal',
		5,
		'Tidak Ada',
		'Tidak Ada'
	);
	
	`)
	err := DeleteSkriningPenyakitById("8990")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}



func TestDeleteTTVById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`"INSERT INTO ttv(
		ttv_id,
		kesadaran, 
		sistole,
		diastole, 
		tinggi_badan, 
		cara_ukur_tb, 
		berat_badan, 
		lingkar_perut, 
		detak_nadi, 
		nafas, 
		saturasi, 
		suhu, 
		detak_jantung, 
		triage, 
		psikolososial_spirit, 
		keterangan
		) VALUES (
			8990, 
    		'Ada', 
    		120, 
    		80, 
    		170, 
    		'Berdiri', 
    		65, 
    		90, 
    		80,
    		18, 
    		98, 
    		37, 
    		70, 
    		'Tidak Gawat Darurat', 
    		'Normal', 
    		'Tidak ada' 
	);
	
	`)
	err := DeleteTTVById("8990")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteTTVById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteTTVById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteTTVById("8990")
	require.Error(t, err)
}

func TestDeleteAlergiById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO alergi (
		alergi_id,
		obat,
		makanan, 
		lainnya
	) VALUES (
		8911,
		'Tidak Ada',
		'Tidak Ada',
		'Tidak Ada'
	);
	
	`)
	err := DeleteAlergiById("8990")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteAlergiById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteAlergiById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteAlergiById("8990")
	require.Error(t, err)
}

func TestDeleteAnamnesisById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO anamnesis (
		anamnesis_id,
		pasien_id,
		skrin_awal_id,
		skrin_gizi_id,
		ttv_id,
		riwayat_penyakit_id,
		alergi_id,
		dokter_id,
		perawat_id,
		keluhan_utama,
		keluhan_tambahan,
		lama_sakit
	) VALUES (
		5134,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		'Sakit kepala',
		'Mual dan pusing',
		8
	);
	
	`)
	err := DeleteAnamnesisById("5134")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteAnamnesisById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteAnamnesisById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteAnamnesisById("8990")
	require.Error(t, err)
}

func TestDeleteRiwayatPenyakitById_Success(t *testing.T) {
	_db := db.DB

	defer func() {
		db.DB = _db
	}()

	_db.Query(`INSERT INTO riwayat_penyakit (
		riwayat_penyakit_id,
		rps,
		rpd,
		rpk
	) VALUES (
		8990
		'Tidak Ada',
		'Tidak Ada',
		'Tidak Ada'
	);
	
	`)
	err := DeleteRiwayatPenyakitById("8990")
	if err != nil {
		t.Errorf("This should not be error, but have %v", err)
	}

}

func TestDeleteRiwayatPenyakitById_Failed(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := DeleteRiwayatPenyakitById("8990")
	require.Error(t, err)

	_db.Close()
	err = DeleteRiwayatPenyakitById("8990")
	require.Error(t, err)
}