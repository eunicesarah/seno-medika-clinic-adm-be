package kasir

import (
	"errors"
	"seno-medika.com/config/db"
	"seno-medika.com/model/station/cashierstation"
)

func FindNotaById(id int) (cashierstation.Nota, error) {
	var notaVar cashierstation.Nota

	err := db.DB.QueryRow("SELECT * FROM nota WHERE nota_id = $1", id).Scan(id)
	if err != nil {
		return cashierstation.Nota{}, err
	}

	err = db.DB.QueryRow("SELECT * FROM nota WHERE nota_id = $1", id).Scan(
		&notaVar.NotaID,
		&notaVar.PasienID,
		&notaVar.DokterID,
		&notaVar.ResepID,
		&notaVar.ListTindakanID,
		&notaVar.TotalBiaya,
		&notaVar.MetodePembayaran,
	)

	if err != nil {
		return cashierstation.Nota{}, err
	}

	return notaVar, nil
}

func FindNotaAll() ([]cashierstation.Nota, error) {
	var notaVar []cashierstation.Nota

	val, err := db.DB.Query("SELECT * FROM nota")
	if err != nil {
		return nil, err
	}
	for val.Next() {
		var eachNota cashierstation.Nota
		err := val.Scan(
			&eachNota.NotaID,
			&eachNota.PasienID,
			&eachNota.DokterID,
			&eachNota.ResepID,
			&eachNota.ListTindakanID,
			&eachNota.TotalBiaya,
			&eachNota.MetodePembayaran,
		)

		if err != nil {
			return nil, err
		}
		notaVar = append(notaVar, eachNota)
	}

	return notaVar, nil
}

func FindNotaByPasienID(id int) ([]cashierstation.Nota, error) {
	var notaVar []cashierstation.Nota

	val, err := db.DB.Query("SELECT * FROM nota WHERE pasien_id = $1", id)
	if err != nil {
		return nil, err
	}
	for val.Next() {
		var eachNota cashierstation.Nota
		err := val.Scan(
			&eachNota.NotaID,
			&eachNota.PasienID,
			&eachNota.DokterID,
			&eachNota.ResepID,
			&eachNota.ListTindakanID,
			&eachNota.TotalBiaya,
			&eachNota.MetodePembayaran,
		)

		if err != nil {
			return nil, err
		}
		notaVar = append(notaVar, eachNota)
	}

	if len(notaVar) == 0 {
		return nil, errors.New("pasien_id not found")
	}

	return notaVar, nil
}

func FindNotaByResepId(id int) ([]cashierstation.Nota, error) {
	var notaVar []cashierstation.Nota

	val, err := db.DB.Query("SELECT * FROM nota WHERE resep_id = $1", id)
	if err != nil {
		return nil, err
	}
	for val.Next() {
		var eachNota cashierstation.Nota
		err := val.Scan(
			&eachNota.NotaID,
			&eachNota.PasienID,
			&eachNota.DokterID,
			&eachNota.ResepID,
			&eachNota.ListTindakanID,
			&eachNota.TotalBiaya,
			&eachNota.MetodePembayaran,
		)

		if err != nil {
			return nil, err
		}
		notaVar = append(notaVar, eachNota)
	}
	if len(notaVar) == 0 {
		return nil, errors.New("resep_id not found")
	}

	return notaVar, nil
}

func FindNotaByMetodePembayaran(metode_pembayaran string) ([]cashierstation.Nota, error) {
	var notaVar []cashierstation.Nota

	val, err := db.DB.Query("SELECT * FROM nota WHERE metode_pembayaran = $1", metode_pembayaran)
	if err != nil {
		return nil, err
	}
	for val.Next() {
		var eachNota cashierstation.Nota
		err := val.Scan(
			&eachNota.NotaID,
			&eachNota.PasienID,
			&eachNota.DokterID,
			&eachNota.ResepID,
			&eachNota.ListTindakanID,
			&eachNota.TotalBiaya,
			&eachNota.MetodePembayaran,
		)

		if err != nil {
			return nil, err
		}
		notaVar = append(notaVar, eachNota)
	}
	if len(notaVar) == 0 {
		return nil, errors.New("metode_pembayaran not found")
	}

	return notaVar, nil
}
