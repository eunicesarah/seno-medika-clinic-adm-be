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
                                role character varying(225) NOT NULL
);

CREATE TABLE public.alergi (
                                alergi_id SERIAL NOT NULL,
                                obat character varying(50),
                                makanan character varying(50),
                                lainnya character varying(100)
);


CREATE TABLE public.anamnesis (
                                anamnesis_id SERIAL NOT NULL,
                                pasien_id SERIAL NOT NULL,
                                skrin_awal_id SERIAL NOT NULL,
                                skrin_gizi_id SERIAL NOT NULL,
                                ttv_id SERIAL NOT NULL,
                                riwayat_penyakit_id SERIAL NOT NULL,
                                alergi_id SERIAL NOT NULL,
                                dokter_id SERIAL NOT NULL,
                                perawat_id SERIAL NOT NULL,
                                keluhan_utama character varying(255) NOT NULL,
                                keluhan_tambahan character varying(255),
                                lama_sakit integer NOT NULL
);

CREATE TABLE public.apoteker (
                                apoteker_id SERIAL NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);


CREATE TABLE public.dokter (
                                dokter_id SERIAL NOT NULL,
                                jaga_poli_mana character varying(50) NOT NULL,
                                jadwal_jaga character varying(50) NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);

CREATE TABLE public.list_jadwal_dokter (
                                dokter_id SERIAL NOT NULL,
                                hari character varying(20) NOT NULL,
                                shift integer NOT NULL
);

CREATE TABLE public.list_obat (
                                obat_id SERIAL NOT NULL,
                                resep_id SERIAL NOT NULL,
                                jumlah integer NOT NULL,
                                dosis character varying(50) NOT NULL,
                                aturan_pakai character varying(50) NOT NULL,
                                keterangan character varying(50) NOT NULL,
);

CREATE TABLE public.nota (
                                nota_id SERIAL NOT NULL,
                                pasien_id SERIAL NOT NULL,
                                dokter_id SERIAL NOT NULL,
                                resep_id SERIAL NOT NULL,
                                list_tindakan_id integer NOT NULL,
                                total_biaya bigint NOT NULL,
                                metode_pembayaran character varying(50) NOT NULL
);

CREATE TABLE public.obat (
                                obat_id SERIAL NOT NULL,
                                nama_obat character varying(50) NOT NULL,
                                jenis_asuransi character varying(50) NOT NULL,
                                harga bigint NOT NULL,
                                stock integer NOT NULL,
                                satuan character varying(20) NOT NULL
);

CREATE TABLE public.perawat (
                                perawat_id SERIAL NOT NULL,
                                nomor_lisensi character varying(50) NOT NULL
);

CREATE TABLE public.rekam_medis (
                                rekam_medis_id SERIAL NOT NULL,
                                pasien_id SERIAL NOT NULL,
                                no_erm character varying(255) NOT NULL,
                                resep_id SERIAL NOT NULL,
                                resume character varying(50) NOT NULL,
                                created_at character varying(225) NOT NULL,
                                update_at character varying(225) NOT NULL
);

CREATE TABLE public.resep (
                                resep_id SERIAL NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                deskripsi character varying(500) NOT NULL,
                                ruang_tujuan character varying(500) NOT NULL,
                                status_obat character varying(500) NOT NULL
);

CREATE TABLE public.riwayat_penyakit (
                                riwayat_penyakit_id SERIAL NOT NULL,
                                rps character varying(20),
                                rpd character varying(20),
                                rpk character varying(20)
);

CREATE TABLE public.skrining_awal (
                                skrin_awal_id SERIAL NOT NULL,
                                disabilitas boolean NOT NULL,
                                ambulansi character varying(100) NOT NULL,
                                hambatan_komunikasi boolean NOT NULL,
                                jalan_tidak_seimbang boolean NOT NULL,
                                jalan_alat_bantu boolean NOT NULL,
                                menopang_saat_duduk boolean NOT NULL,
                                hasil_cara_jalan character varying(50) NOT NULL,
                                skala_nyeri integer NOT NULL,
                                nyeri_berulang character varying(100) NOT NULL,
                                sifat_nyeri character varying(100) NOT NULL
);

CREATE TABLE public.skrining_gizi (
                                skrin_gizi_id SERIAL NOT NULL,
                                penurunan_bb integer NOT NULL,
                                tdk_nafsu_makan boolean NOT NULL,
                                diagnosis_khusus boolean NOT NULL,
                                nama_penyakit character varying(100)
);

CREATE TABLE public.ttv (
                                ttv_id SERIAL NOT NULL,
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

CREATE TABLE public.list_cppt (
                                pasien_id integer NOT NULL,
                                no_erm character varying(255) NOT NULL,
                                cppt_id integer NOT NULL
);

CREATE TABLE public.cppt (
                                cppt_id SERIAL NOT NULL,
                                unit character varying(100) NOT NULL,
                                tanggal date NOT NULL,
                                objektif character varying(255) NOT NULL,
                                assessment character varying(255) NOT NULL,
                                penatalaksanaan character varying(255) NOT NULL
);

CREATE TABLE public.pemeriksaan_dokter (
                                pemeriksaan_dokter_id SERIAL NOT NULL,
                                pasien_id integer NOT NULL,
                                dokter_id integer NOT NULL,
                                perawat_id integer NOT NULL
);

CREATE TABLE public.pemeriksaan_fisik (
                                pemeriksaan_fisik_id SERIAL NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                terapi_yg_sdh_dilakukan character varying(255),
                                rencana_tindakan character varying(255),
                                tindakan_keperawatan character varying(255),
                                observasi character varying(255),
                                merokok boolean,
                                konsumsi_alkohol boolean,
                                kurang_sayur boolean
);

CREATE TABLE public.riwayat_pemeriksaan (
                                riwayat_pemeriksaan_id SERIAL NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                pasien_id integer NOT NULL,
                                tanggal date NOT NULL,
                                pemeriksaan character varying(255),
                                keterangan character varying(255)
);

CREATE TABLE public.keadaan_fisik (
                                keadaan_fisik_id SERIAL NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                pemeriksaan_kulit boolean,
                                pemeriksaan_kuku boolean,
                                pemeriksaan_kepala boolean,
                                pemeriksaan_mata boolean,
                                pemeriksaan_telinga boolean,
                                pemeriksaan_hidung_sinus boolean,
                                pemeriksaan_mulut_bibir boolean,
                                pemeriksaan_leher boolean,
                                pemeriksaan_dada_punggung boolean,
                                pemeriksaan_kardiovaskuler boolean,
                                pemeriksaan_abdomen_perut boolean,
                                pemeriksaan_ekstremitas_atas boolean,
                                pemeriksaan_ekstremitas_bawah boolean,
                                pemeriksaan_genitalia_pria boolean
);

CREATE TABLE public.diagnosa (
                                diagnosa_id SERIAL NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                diagnosa character varying(255) NOT NULL,
                                jenis character varying(255) NOT NULL,
                                kasus character varying(255) NOT NULL,
                                status_diagnosis character varying(255) NOT NULL
);

CREATE TABLE public.list_tindakan (
                                list_tindakan_id SERIAL NOT NULL
);

CREATE TABLE public.tindakan (
                                tindakan_id SERIAL NOT NULL,
                                jenis_tindakan character varying(500) NOT NULL,
                                prosedur_tindakan character varying(500),
                                jumlah integer NOT NULL,
                                keterangan character varying(500),
                                tanggal_rencana date,
                                harga_tindakan bigint NOT NULL,
                                indikasi_tindakan character varying(500),
                                tujuan character varying(500),
                                risiko character varying(500),
                                komplikasi character varying(500),
                                alternatif_risiko character varying(500)
);

CREATE TABLE public.penanganan (
                                tindakan_id integer NOT NULL,
                                list_tindakan_id integer NOT NULL
);

CREATE TABLE public.anatomi (
                                anatomi_id SERIAL NOT NULL,
                                pasien_id integer NOT NULL,
                                pemeriksaan_dokter_id integer NOT NULL,
                                bagian_tubuh character varying(255) NOT NULL,
                                keterangan character varying(255) NOT NULL
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

ALTER TABLE ONLY public.cppt
    ADD CONSTRAINT cppt_cppt_id_pkey PRIMARY KEY (cppt_id);

ALTER TABLE ONLY public.pemeriksaan_fisik
    ADD CONSTRAINT pemeriksaan_fisik_pemeriksaan_fisik_id_pkey PRIMARY KEY (pemeriksaan_fisik_id);

ALTER TABLE ONLY public.pemeriksaan_dokter
    ADD CONSTRAINT pemeriksaan_dokter_pemeriksaan_dokter_id_pkey PRIMARY KEY (pemeriksaan_dokter_id);

ALTER TABLE ONLY public.riwayat_pemeriksaan
    ADD CONSTRAINT riwayat_pemeriksaan_riwayat_pemeriksaan_id_pkey PRIMARY KEY (riwayat_pemeriksaan_id);

ALTER TABLE ONLY public.keadaan_fisik
    ADD CONSTRAINT keadaan_fisik_keadaan_fisik_id_pkey PRIMARY KEY (keadaan_fisik_id);

ALTER TABLE ONLY public.diagnosa
    ADD CONSTRAINT diagnosa_diagnosa_id_pkey PRIMARY KEY (diagnosa_id);

ALTER TABLE ONLY public.tindakan
    ADD CONSTRAINT tindakan_tindakan_id_pkey PRIMARY KEY (tindakan_id);

ALTER TABLE ONLY public.list_tindakan
    ADD CONSTRAINT list_tindakan_list_tindakan_id_pkey PRIMARY KEY (list_tindakan_id);

ALTER TABLE ONLY public.anatomi
    ADD CONSTRAINT anatomi_anatomi_id_pkey PRIMARY KEY (anatomi_id);

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_alergi_id_fkey FOREIGN KEY (alergi_id) REFERENCES public.alergi(alergi_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_perawat_id_fkey FOREIGN KEY (perawat_id) REFERENCES public.perawat(perawat_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_riwayat_penyakit_id_fkey FOREIGN KEY (riwayat_penyakit_id) REFERENCES public.riwayat_penyakit(riwayat_penyakit_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_skrin_awal_id_fkey FOREIGN KEY (skrin_awal_id) REFERENCES public.skrining_awal(skrin_awal_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_skrin_gizi_id_fkey FOREIGN KEY (skrin_gizi_id) REFERENCES public.skrining_gizi(skrin_gizi_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anamnesis
    ADD CONSTRAINT anamnesis_ttv_id_fkey FOREIGN KEY (ttv_id) REFERENCES public.ttv(ttv_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.apoteker
    ADD CONSTRAINT apoteker_apoteker_id_fkey FOREIGN KEY (apoteker_id) REFERENCES public.users(user_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.dokter
    ADD CONSTRAINT dokter_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.users(user_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_jadwal_dokter
    ADD CONSTRAINT list_jadwal_dokter_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_obat
    ADD CONSTRAINT list_obat_obat_id_fkey FOREIGN KEY (obat_id) REFERENCES public.obat(obat_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_obat
    ADD CONSTRAINT list_obat_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.nota
    ADD CONSTRAINT nota_list_tindakan_id_fkey FOREIGN KEY (list_tindakan_id) REFERENCES public.list_tindakan(list_tindakan_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.perawat
    ADD CONSTRAINT perawat_perawat_id_fkey FOREIGN KEY (perawat_id) REFERENCES public.users(user_id)ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_no_erm_fkey FOREIGN KEY (no_erm) REFERENCES public.pasien(no_erm) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.rekam_medis
    ADD CONSTRAINT rekam_medis_resep_id_fkey FOREIGN KEY (resep_id) REFERENCES public.resep(resep_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_cppt
    ADD CONSTRAINT list_cppt_cppt_id_fkey FOREIGN KEY (cppt_id) REFERENCES public.cppt(cppt_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_cppt
    ADD CONSTRAINT list_cppt_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.list_cppt
    ADD CONSTRAINT list_cppt_no_erm_fkey FOREIGN KEY (no_erm) REFERENCES public.pasien(no_erm) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.pemeriksaan_dokter
    ADD CONSTRAINT pemeriksaan_dokter_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.pemeriksaan_dokter
    ADD CONSTRAINT pemeriksaan_dokter_dokter_id_fkey FOREIGN KEY (dokter_id) REFERENCES public.dokter(dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.pemeriksaan_dokter
    ADD CONSTRAINT pemeriksaan_dokter_perawat_id_fkey FOREIGN KEY (perawat_id) REFERENCES public.perawat(perawat_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.riwayat_pemeriksaan
    ADD CONSTRAINT riwayat_pemeriksaan_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.riwayat_pemeriksaan
    ADD CONSTRAINT riwayat_pemeriksaan_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.penanganan
    ADD CONSTRAINT penanganan_tindakan_id_fkey FOREIGN KEY (tindakan_id) REFERENCES public.tindakan(tindakan_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.penanganan
    ADD CONSTRAINT penanganan_list_tindakan_id_fkey FOREIGN KEY (list_tindakan_id) REFERENCES public.list_tindakan(list_tindakan_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anatomi
    ADD CONSTRAINT anatomi_pasien_id_fkey FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.anatomi
    ADD CONSTRAINT anatomi_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.resep
    ADD CONSTRAINT resep_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.pemeriksaan_fisik
    ADD CONSTRAINT pemeriksaan_fisik_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.keadaan_fisik
    ADD CONSTRAINT keadaan_fisik_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY public.diagnosa
    ADD CONSTRAINT diagnosa_pemeriksaan_dokter_id_fkey FOREIGN KEY (pemeriksaan_dokter_id) REFERENCES public.pemeriksaan_dokter(pemeriksaan_dokter_id) ON UPDATE CASCADE ON DELETE CASCADE;