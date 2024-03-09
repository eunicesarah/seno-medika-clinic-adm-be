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


ALTER TABLE ONLY public.antrian
    ADD CONSTRAINT "PK_Antrian" PRIMARY KEY (antrian_id);

ALTER TABLE ONLY public.pasien
    ADD CONSTRAINT "PK_pasien" PRIMARY KEY (pasien_id, no_erm);

ALTER TABLE ONLY public.pasien
    ADD CONSTRAINT "UNIQUE_Pasien" UNIQUE (pasien_id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (user_id);

ALTER TABLE ONLY public.antrian
    ADD CONSTRAINT "FK_Antrian" FOREIGN KEY (pasien_id) REFERENCES public.pasien(pasien_id) NOT VALID;