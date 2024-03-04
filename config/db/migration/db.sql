CREATE TABLE public."Antrian" (
                                  antrian_id integer NOT NULL,
                                  pasien_id integer NOT NULL,
                                  nomor_antrian integer NOT NULL,
                                  status boolean NOT NULL,
                                  poli character varying(20) NOT NULL,
                                  instalasi character varying(20) NOT NULL,
                                  create_at date NOT NULL
);

CREATE TABLE public."Pasien" (
                                 pasien_id integer NOT NULL,
                                 no_erm integer NOT NULL,
                                 pasien_uuid uuid NOT NULL,
                                 "no_RM_lama" character varying(50) NOT NULL,
                                 "no_dok_RM" character varying(50) NOT NULL,
                                 penjamin character varying(50) NOT NULL,
                                 no_penjamin character varying(50) NOT NULL,
                                 "NIK" character varying(16) NOT NULL,
                                 "no_KK" character varying(16) NOT NULL,
                                 nama character varying(50) NOT NULL,
                                 tempat_lahir character varying(50) NOT NULL,
                                 tanggal_lahir date NOT NULL,
                                 "no_IHS" character varying(50) NOT NULL,
                                 jenis_kelamin character varying(12) NOT NULL,
                                 golongan_darah character varying(2) NOT NULL,
                                 no_telpon character varying(14) NOT NULL,
                                 email character varying(50) NOT NULL,
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
                                 created_at date NOT NULL,
                                 created_by character varying(50) NOT NULL,
                                 updated_at date NOT NULL,
                                 updated_by character varying(50) NOT NULL
);

CREATE TABLE public."Users" (
                                user_id integer NOT NULL,
                                user_uuid character varying(50) NOT NULL,
                                nama character varying(50) NOT NULL,
                                password character varying(50) NOT NULL,
                                email character varying(50) NOT NULL,
                                role character varying(50) NOT NULL
);


ALTER TABLE ONLY public."Antrian"
    ADD CONSTRAINT "PK_Antrian" PRIMARY KEY (antrian_id);

ALTER TABLE ONLY public."Pasien"
    ADD CONSTRAINT "PK_pasien" PRIMARY KEY (pasien_id, no_erm);

ALTER TABLE ONLY public."Pasien"
    ADD CONSTRAINT "UNIQUE_Pasien" UNIQUE (pasien_id);

ALTER TABLE ONLY public."Users"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (user_id);

ALTER TABLE ONLY public."Antrian"
    ADD CONSTRAINT "FK_Antrian" FOREIGN KEY (pasien_id) REFERENCES public."Pasien"(pasien_id) NOT VALID;