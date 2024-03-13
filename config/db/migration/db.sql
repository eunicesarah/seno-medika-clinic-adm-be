CREATE TABLE public.antrian (
                                antrian_id SERIAL NOT NULL,
                                pasien_id SERIAL NOT NULL,
                                nomor_antrian integer NOT NULL,
                                status boolean NOT NULL,
                                poli character varying(20) NOT NULL,
                                instalasi character varying(20) NOT NULL,
                                created_at character varying(225) NOT NULL
);

CREATE TABLE public.pasien (
                                pasien_id SERIAL NOT NULL,
                                no_erm character varying(225) NOT NULL,
                                pasien_uuid uuid NOT NULL,
                                no_rm_lama character varying(50),
                                no_dok_rm character varying(50),
                                penjamin character varying(50) NOT NULL,
                                no_penjamin character varying(50),
                                nik character varying(16) NOT NULL,
                                no_kk character varying(16),
                                nama character varying(50) NOT NULL,
                                tempat_lahir character varying(50) NOT NULL,
                                tanggal_lahir date NOT NULL,
                                no_ihs character varying(50),
                                jenis_kelamin character varying(12) NOT NULL,
                                golongan_darah character varying(2) NOT NULL,
                                no_telpon character varying(14) NOT NULL,
                                email character varying(50),
                                provinsi character varying(50) NOT NULL,
                                kabupaten_kota character varying(50) NOT NULL,
                                kecamatan character varying(50) NOT NULL,
                                kelurahan character varying(50) NOT NULL,
                                alamat character varying(50) NOT NULL,
                                nama_kontak_darurat character varying(50) NOT NULL,
                                nomor_kontak_darurat character varying(15) NOT NULL,
                                pekerjaan character varying(50) NOT NULL,
                                agama character varying(20) NOT NULL,
                                warga_negara character varying(20) NOT NULL,
                                pendidikan character varying(5) NOT NULL,
                                status_perkawinan character varying(20) NOT NULL,
                                created_at character varying(225) NOT NULL,
                                created_by character varying(50) NOT NULL,
                                updated_at character varying(225) NOT NULL,
                                updated_by character varying(50) NOT NULL
);

CREATE TABLE public.users (
                                user_id SERIAL NOT NULL,
                                user_uuid uuid NOT NULL,
                                nama character varying(225) NOT NULL,
                                password character varying(225) NOT NULL,
                                email character varying(225) NOT NULL,
                                role character varying(10) NOT NULL
);

CREATE TABLE public.alergi (
                                alergi_id integer NOT NULL,
                                obat character varying(50),
                                makanan character varying(50),
                                lainnya character varying(100)
);


CREATE TABLE public.anamnesis (
                                anamnesis_id integer NOT NULL,
                                pasien_id integer NOT NULL,
                                skrin_awal_id integer NOT NULL,
                                skrin_gizi_id integer NOT NULL,
                                ttv_id integer NOT NULL,
                                riwayat_penyakit_id integer NOT NULL,
                                alergi_id integer NOT NULL,
                                dokter_id integer NOT NULL,
                                perawat_id integer NOT NULL,
                                keluhan_utama character varying(255) NOT NULL,
                                keluhan_tambahan character varying(255),
                                lama_sakit integer NOT NULL
);

CREATE TABLE public.apoteker (
                                apoteker_id integer NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);


CREATE TABLE public.dokter (
                                dokter_id integer NOT NULL,
                                jaga_poli_mana character varying(50) NOT NULL,
                                jadwal_jaga character varying(50) NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);

CREATE TABLE public.list_jadwal_dokter (
                                dokter_id integer NOT NULL,
                                hari character varying(20) NOT NULL,
                                shift integer NOT NULL
);

CREATE TABLE public.list_obat (
                                obat_id integer NOT NULL,
                                resep_id integer NOT NULL,
                                jumlah integer NOT NULL,
                                dosis character varying(50) NOT NULL
);

CREATE TABLE public.nota (
                                nota_id integer NOT NULL,
                                pasien_id integer NOT NULL,
                                dokter_id integer NOT NULL,
                                resep_id integer NOT NULL,
                                total_biaya integer NOT NULL
);

CREATE TABLE public.obat (
                                obat_id integer NOT NULL,
                                nama_obat character varying(50) NOT NULL,
                                jenis_asuransi character varying(50) NOT NULL,
                                harga integer NOT NULL
);

CREATE TABLE public.perawat (
                                perawat_id integer NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);

CREATE TABLE public.rekam_medis (
                                rekam_medis_id integer NOT NULL,
                                pasien_id integer NOT NULL,
                                no_erm character varying(50) NOT NULL,
                                resep_id integer NOT NULL,
                                resume character varying(50) NOT NULL,
                                created_at character varying(225) NOT NULL,
                                update_at character varying(225) NOT NULL
);

CREATE TABLE public.resep (
                                resep_id integer NOT NULL,
                                deskripsi character varying(500) NOT NULL
);

CREATE TABLE public.riwayat_penyakit (
                                riwayat_penyakit_id integer NOT NULL,
                                rps character varying(20),
                                rpd character varying(20),
                                rpk character varying(20)
);

CREATE TABLE public.skrining_awal (
                                skrin_awal_id integer NOT NULL,
                                disabilitas boolean NOT NULL,
                                ambulansi boolean NOT NULL,
                                hambatan _komunikasi boolean NOT NULL,
                                jalan_tidak_seimbang boolean NOT NULL,
                                jalan_alat_bantu boolean NOT NULL,
                                menopang_saat_duduk boolean NOT NULL,
                                hasil_cara_jalan character varying(50) NOT NULL,
                                skala_nyeri integer NOT NULL,
                                nyeri_berulang character varying(100) NOT NULL,
                                sifat_nyeri character varying(100) NOT NULL
);

CREATE TABLE public.skrining_gizi (
                                skrin_gizi_id integer NOT NULL,
                                penurunan_bb integer NOT NULL,
                                tdk_nafsu_makan boolean NOT NULL,
                                diagnosis_khusus boolean NOT NULL,
                                nama_penyakit character varying(100),
                                skala_nyeri integer NOT NULL,
                                nyeri_berulang character varying(100) NOT NULL,
                                sifat_nyeri character varying(100) NOT NULL
);

CREATE TABLE public.ttv (
                                ttv_id integer NOT NULL,
                                kesadaran character varying(20) NOT NULL,
                                sistole integer NOT NULL,
                                diastole integer NOT NULL,
                                tinggi_badan integer NOT NULL,
                                cara_ukur_tb character varying(10) NOT NULL,
                                berat_badan integer NOT NULL,
                                lingkar_perut integer NOT NULL,
                                detak_nadi integer NOT NULL,
                                nafas integer NOT NULL,
                                saturasi integer,
                                suhu integer,
                                detak_jantung boolean,
                                triage character varying(50) NOT NULL,
                                psikolososial_spirit character varying(255),
                                keterangan character varying(255)
);


ALTER TABLE ONLY public.antrian
    ADD CONSTRAINT "PK_Antrian" PRIMARY KEY (antrian_id);

ALTER TABLE ONLY public.pasien
    ADD CONSTRAINT "PK_pasien" PRIMARY KEY (pasien_id, no_erm);

ALTER TABLE ONLY public.pasien
    ADD CONSTRAINT "UNIQUE_Pasien" UNIQUE (pasien_id);

ALTER TABLE ONLY public.pasien
    ADD CONSTRAINT "UNIQUE_Pasien2" UNIQUE (no_erm);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (user_id);

ALTER TABLE ONLY public.antrian
    ADD CONSTRAINT "FK_Antrian" FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) NOT VALID;

ALTER TABLE ONLY public.resep
    ADD CONSTRAINT resep_pkey PRIMARY KEY (resep_id);

ALTER TABLE ONLY public.alergi
    ADD CONSTRAINT alergi_pkey PRIMARY KEY (alergi_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_pkey PRIMARY KEY (anamnesis_id);

ALTER TABLE ONLY public.apoteker
    ADD CONSTRAINT apoteker_pkey PRIMARY KEY (apoteker_id);

ALTER TABLE ONLY public.dokter
    ADD CONSTRAINT dokter_pkey PRIMARY KEY (dokter_id);

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_pkey PRIMARY KEY (nota_id);

ALTER TABLE ONLY public.obat
    ADD CONSTRAINT obat_pkey PRIMARY KEY (obat_id);

ALTER TABLE ONLY public.perawat
    ADD CONSTRAINT perawat_pkey PRIMARY KEY (perawat_id);

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_pkey PRIMARY KEY (rekam_medis_id);

ALTER TABLE ONLY public.riwayat_penyakit
    ADD CONSTRAINT riwayat_penyakit_pkey PRIMARY KEY (riwayat_penyakit_id);

ALTER TABLE ONLY public.skrining_awal
    ADD CONSTRAINT skrining_awal_pkey PRIMARY KEY (skrin_awal_id);

ALTER TABLE ONLY public.skrining_gizi
    ADD CONSTRAINT skrining_gizi_pkey PRIMARY KEY (skrin_gizi_id);

ALTER TABLE ONLY public.ttv
    ADD CONSTRAINT ttv_pkey PRIMARY KEY (ttv_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_alergi_id_fkey FOREIGN KEY (alergi_id) REFERENCES public.alergi(alergi_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_perawat_id_fkey FOREIGN KEY (perawat_id) REFERENCES public.perawat(perawat_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_riwayat_penyakit_id_fkey FOREIGN KEY (riwayat_penyakit_id) REFERENCES public.riwayat_penyakit(riwayat_penyakit_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_skrin_awal_id_fkey FOREIGN KEY (skrin_awal_id) REFERENCES public.skrining_awal(skrin_awal_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_skrin_gizi_id_fkey FOREIGN KEY (skrin_gizi_id) REFERENCES public.skrining_gizi(skrin_gizi_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_ttv_id_fkey FOREIGN KEY (ttv_id) REFERENCES public.ttv(ttv_id);

ALTER TABLE ONLY public.apoteker
    ADD CONSTRAINT apoteker_apoteker_id_fkey FOREIGN KEY (apoteker_id) REFERENCES public.users(user_id);

ALTER TABLE ONLY public.dokter
    ADD CONSTRAINT dokter_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.users(user_id);

ALTER TABLE ONLY public.list_jadwal_dokter
    ADD CONSTRAINT list_jadwal_dokter_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id);

ALTER TABLE ONLY public.list_obat
    ADD CONSTRAINT list_obat_obat_id_fkey FOREIGN KEY (obat_id) REFERENCES public.obat(obat_id);

ALTER TABLE ONLY public.list_obat
    ADD CONSTRAINT list_obat_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id);

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id);

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id);

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id);

ALTER TABLE ONLY public.perawat
    ADD CONSTRAINT perawat_perawat_id_fkey FOREIGN KEY (perawat_id) REFERENCES public.users(user_id);

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_no_erm_fkey FOREIGN KEY (no_erm) REFERENCES public.pasien(no_erm) NOT VALID;

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) NOT VALID;

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id) NOT VALID;
