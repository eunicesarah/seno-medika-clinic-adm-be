package kasir

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"seno-medika.com/config/db"
	"testing"
	"time"
	// "seno-medika.com/model/cashierstation"
)

func TestUpdateMetodePembayaran_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()
	var obatId int
	var pasienId int
	var dokterId int
	var ResepId int
	var ListTindakanId int
	var NotaId int
	var PerawatId int
	var PemeriksaanDokterId int

	if err := _db.QueryRow(`INSERT INTO obat (
		nama_obat,
		harga,
		jenis_asuransi,
		stock,
		satuan

	) VALUES (
		'nama_obat',
		10000,
		'jenis_asuransi',
		100,
		'sirup'
	) RETURNING obat_id`).Scan(&obatId); err != nil {
		t.Errorf("error: %v", err)
	}

	if err := _db.QueryRow(`INSERT INTO pasien (
		no_erm,
		pasien_uuid,
		no_rm_lama,
		no_dok_rm,
		penjamin,
		no_penjamin,
		nik,
		no_kk,
		nama,
		tempat_lahir,
		tanggal_lahir,
		no_ihs,
		jenis_kelamin,
		golongan_darah,
		no_telpon,
		email,
		provinsi,
		kabupaten_kota,
		kecamatan,
		kelurahan,
		alamat,
		nama_kontak_darurat,
		nomor_kontak_darurat,
		pekerjaan,
		agama,
		warga_negara,
		pendidikan,
		status_perkawinan,
		created_at,
		created_by,
		updated_at,
		updated_by
	   ) VALUES ( '1232222', $1, '123', '123', 'BPJS', '123', '330122202', '124212',
		'testNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '031286123',
		'12223@tesdst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO users (
		user_uuid,
		nama,
		password,
		email,
		role
	   ) VALUES ( $1, 'dokter1', 'password', 'email@email.com', 'dokter'
	   ) RETURNING user_id
	   `, uuid.New()).Scan(&dokterId); err != nil {
		t.Error(err)
		return
	}
	if err := _db.QueryRow(`INSERT INTO users (
		user_uuid,
		nama,
		password,
		email,
		role
	   ) VALUES ( $1, 'dokter1', 'password', 'email@email.com', 'perawat'
	   ) RETURNING user_id
	   `, uuid.New()).Scan(&PerawatId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO dokter (
		dokter_id,
		jaga_poli_mana,
		jadwal_jaga,
		nomor_lisensi
	) VALUES ($1, 'umum', 'senin', '1234567890') RETURNING dokter_id
	`, dokterId).Scan(&dokterId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO perawat (
		perawat_id,
		nomor_lisensi
	) VALUES ($1, '123123123') RETURNING perawat_id
	`, PerawatId).Scan(&PerawatId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO pemeriksaan_dokter (
		pasien_id,
		dokter_id,
		perawat_id
	) VALUES ($1, $2, $3) RETURNING pemeriksaan_dokter_id
	`, pasienId, dokterId, PerawatId).Scan(&PemeriksaanDokterId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO resep (
		deskripsi,
		pemeriksaan_dokter_id,
		ruang_tujuan,
		status_obat
	) VALUES ('deskripsi', $1,'surga', 'selesai') RETURNING resep_id
	`, PemeriksaanDokterId).Scan(&ResepId); err != nil {
		t.Error(err)
		return
	}

	_db.Exec(`INSERT INTO list_obat (
		obat_id,
		resep_id,
		jumlah,
		dosis,
		keterangan,
		tanggal_kadaluarsa
	) VALUES ($1, $2, 100, '1x1', 'wajib malam', '2021-01-01')
	`, obatId, ResepId)

	if err := _db.QueryRow(`INSERT INTO list_tindakan (
		list_tindakan_id
	) VALUES (9990) RETURNING list_tindakan_id
	`).Scan(&ListTindakanId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO nota (
		pasien_id,
		dokter_id,
		resep_id,
		list_tindakan_id,
		total_biaya,
		metode_pembayaran
	) VALUES ($1, $2, $3, $4, 10000, 'qris') RETURNING nota_id
	`, pasienId, dokterId, ResepId, ListTindakanId).Scan(&NotaId); err != nil {
		t.Error(err)
		return
	}

	err := UpdateMetodePembayaran(NotaId, "ovo")
	require.NoError(t, err)

	_db.Exec("DELETE FROM nota WHERE nota_id = $1", NotaId)
	_db.Exec("DELETE FROM list_tindakan WHERE list_tindakan_id = $1", ListTindakanId)
	_db.Exec("DELETE FROM resep WHERE resep_id = $1", ResepId)
	_db.Exec("DELETE FROM dokter WHERE dokter_id = $1", dokterId)
	_db.Exec("DELETE FROM users WHERE user_id = $1", dokterId)
	_db.Exec("DELETE FROM pasien WHERE pasien_id = $1", pasienId)
	_db.Exec("DELETE FROM obat WHERE obat_id = $1", obatId)
	_db.Exec("DELETE FROM pemeriksaan_dokter WHERE pemeriksaan_dokter_id = $1", PemeriksaanDokterId)
	_db.Close()

}
func TestUpdateMetodePembayaran_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := UpdateMetodePembayaran(998, "ovo")
	require.Error(t, err)

	_db.Close()
	err = UpdateMetodePembayaran(998, "ovo")
	require.Error(t, err)
}
func TestUpdateTotalBiaya_Success(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	var obatId int
	var pasienId int
	var dokterId int
	var ResepId int
	var ListTindakanId int
	var NotaId int
	var PerawatId int
	var PemeriksaanDokterId int

	if err := _db.QueryRow(`INSERT INTO obat (
		nama_obat,
		harga,
		jenis_asuransi,
		stock,
		satuan

	) VALUES (
		'nama_obat',
		10000,
		'jenis_asuransi',
		100,
		'sirup'
	) RETURNING obat_id`).Scan(&obatId); err != nil {
		t.Errorf("error: %v", err)
	}

	if err := _db.QueryRow(`INSERT INTO pasien (
		no_erm,
		pasien_uuid,
		no_rm_lama,
		no_dok_rm,
		penjamin,
		no_penjamin,
		nik,
		no_kk,
		nama,
		tempat_lahir,
		tanggal_lahir,
		no_ihs,
		jenis_kelamin,
		golongan_darah,
		no_telpon,
		email,
		provinsi,
		kabupaten_kota,
		kecamatan,
		kelurahan,
		alamat,
		nama_kontak_darurat,
		nomor_kontak_darurat,
		pekerjaan,
		agama,
		warga_negara,
		pendidikan,
		status_perkawinan,
		created_at,
		created_by,
		updated_at,
		updated_by
	   ) VALUES ( '1231321222', $1, '123', '123', 'BPJS', '123', '3301222202', '124212',
		'testNamadfb123', 'sumedfbdfbdang', $2, '123dfbdfb45', 'perempuan', 'E', '0312486123',
		'122123@te3sdst.go', 'jawa bali', 'jatinandfbdfbgor', 'saydfbdfbang', 'cikedfbdfbruh', 'jalan-jaldfbdfban no 12', 'jokoTesdfbdfbt', '0852345237123',
		'dokter','islam	', 'WNA', 'SMA', 'belum-dfbdfbkawin', $3, '123test', $4, '123test'
	   ) RETURNING pasien_id
	   `, uuid.New(), time.Now().Local().Format("2006-01-02"), time.Now().Local().String(), time.Now().Local().String()).Scan(&pasienId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO users (
		user_uuid,
		nama,
		password,
		email,
		role
	   ) VALUES ( $1, 'dokter1', 'password', 'email@email.com', 'dokter'
	   ) RETURNING user_id
	   `, uuid.New()).Scan(&dokterId); err != nil {
		t.Error(err)
		return
	}
	if err := _db.QueryRow(`INSERT INTO users (
		user_uuid,
		nama,
		password,
		email,
		role
	   ) VALUES ( $1, 'dokter1', 'password', 'email@email.com', 'perawat'
	   ) RETURNING user_id
	   `, uuid.New()).Scan(&PerawatId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO dokter (
		dokter_id,
		jaga_poli_mana,
		jadwal_jaga,
		nomor_lisensi
	) VALUES ($1, 'umum', 'senin', '1234567890') RETURNING dokter_id
	`, dokterId).Scan(&dokterId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO perawat (
		perawat_id,
		nomor_lisensi
	) VALUES ($1, '123123123') RETURNING perawat_id
	`, PerawatId).Scan(&PerawatId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO pemeriksaan_dokter (
		pasien_id,
		dokter_id,
		perawat_id
	) VALUES ($1, $2, $3) RETURNING pemeriksaan_dokter_id
	`, pasienId, dokterId, PerawatId).Scan(&PemeriksaanDokterId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO resep (
		deskripsi,
		pemeriksaan_dokter_id,
		ruang_tujuan,
		status_obat
	) VALUES ('deskripsi', $1,'surga', 'selesai') RETURNING resep_id
	`, PemeriksaanDokterId).Scan(&ResepId); err != nil {
		t.Error(err)
		return
	}

	_db.Exec(`INSERT INTO list_obat (
		obat_id,
		resep_id,
		jumlah,
		dosis,
		keterangan,
		tanggal_kadaluarsa
	) VALUES ($1, $2, 100, '1x1', 'wajib malam', '2021-01-01')
	`, obatId, ResepId)

	if err := _db.QueryRow(`INSERT INTO list_tindakan (
		list_tindakan_id
	) VALUES (9991) RETURNING list_tindakan_id
	`).Scan(&ListTindakanId); err != nil {
		t.Error(err)
		return
	}

	if err := _db.QueryRow(`INSERT INTO nota (
		pasien_id,
		dokter_id,
		resep_id,
		list_tindakan_id,
		total_biaya,
		metode_pembayaran
	) VALUES ($1, $2, $3, $4, 10000, 'qris') RETURNING nota_id
	`, pasienId, dokterId, ResepId, ListTindakanId).Scan(&NotaId); err != nil {
		t.Error(err)
		return
	}

	err := UpdateTotalBiaya(NotaId, 10000000)
	require.NoError(t, err)

	_db.Exec("DELETE FROM nota WHERE nota_id = $1", NotaId)
	_db.Exec("DELETE FROM list_tindakan WHERE list_tindakan_id = $1", ListTindakanId)
	_db.Exec("DELETE FROM resep WHERE resep_id = $1", ResepId)
	_db.Exec("DELETE FROM dokter WHERE dokter_id = $1", dokterId)
	_db.Exec("DELETE FROM users WHERE user_id = $1", dokterId)
	_db.Exec("DELETE FROM pasien WHERE pasien_id = $1", pasienId)
	_db.Exec("DELETE FROM obat WHERE obat_id = $1", obatId)
	_db.Exec("DELETE FROM pemeriksaan_dokter WHERE pemeriksaan_dokter_id = $1", PemeriksaanDokterId)
	_db.Close()
}
func TestUpdateTotalBiaya_Fail(t *testing.T) {
	_db := db.DB
	defer func() {
		_db = db.Conn()
		db.DB = _db
	}()

	err := UpdateTotalBiaya(998, 1000)
	require.Error(t, err)

	_db.Close()
	err = UpdateTotalBiaya(998, 1000)
	require.Error(t, err)
}
