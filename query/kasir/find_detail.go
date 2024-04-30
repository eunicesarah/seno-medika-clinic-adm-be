package kasir

import (
	"errors"

	"seno-medika.com/config/db"
	"seno-medika.com/model/cashierstation"
	"seno-medika.com/model/pharmacystation"
)

func FindDetailByResepId(nota_id int) ([]pharmacystation.DetailObat, error) {
	var details []pharmacystation.DetailObat

	rows, err := db.DB.Query("SELECT o.nama_obat,  o.harga, lo.jumlah, lo.dosis, o.obat_id, o.jenis_asuransi, lo.resep_id, lo.obat_id FROM nota n "+
		"INNER JOIN list_obat lo ON n.resep_id = lo.resep_id "+
		"INNER JOIN obat o ON lo.obat_id = o.obat_id "+
		"WHERE n.nota_id = $1", nota_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail cashierstation.DetailNota
		err := rows.Scan(&detail.Obat.NamaObat, &detail.Obat.Harga, &detail.Jumlah, &detail.Dosis, &detail.Obat.ObatID, &detail.Obat.JenisAsuransi, &detail.ListObat.ResepID, &detail.ListObat.ObatID)
		if err != nil {
			return nil, err
		}
		obat := pharmacystation.Obat(detail.Obat)
		listObat := pharmacystation.ListObat(detail.ListObat)
		listObat.Jumlah = detail.Jumlah
		listObat.Dosis = detail.Dosis
		details = append(details, pharmacystation.DetailObat{Obat: obat, ListObat: listObat})
	}

	if len(details) == 0 {
		return nil, errors.New("nota_id not found")
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return details, nil
}

func FindTindakanByNotaId(nota_id int) ([]cashierstation.Tindakan, error) {
	var tindakans []cashierstation.Tindakan

	rows, err := db.DB.Query("SELECT t.tindakan_id, t.jenis_tindakan, t.keterangan, t.harga_tindakan FROM penanganan p "+
		"INNER JOIN list_tindakan lt ON p.list_tindakan_id = lt.list_tindakan_id "+
		"INNER JOIN tindakan t ON p.tindakan_id = t.tindakan_id "+
		"INNER JOIN nota n ON n.list_tindakan_id = lt.list_tindakan_id "+
		"WHERE n.nota_id = $1", nota_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tindakan cashierstation.Tindakan
		err := rows.Scan(&tindakan.TindakanID, &tindakan.JenisTindakan, &tindakan.Keterangan, &tindakan.HargaTindakan)
		if err != nil {
			return nil, err
		}
		tindakans = append(tindakans, tindakan)
	}
	if len(tindakans) == 0 {
		return nil, errors.New("nota_id not found")
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tindakans, nil
}
