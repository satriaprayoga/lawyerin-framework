DROP TABLE IF EXISTS "public"."firm";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS firm_firm_id_seq;

-- Table Definition
CREATE TABLE "public"."firm" (
    "firm_id" int8 NOT NULL DEFAULT nextval('firm_firm_id_seq'::regclass),
    "firm_name" varchar(100),
    "address" varchar(250),
    "province" varchar(250),
    "city" varchar(250),
    "since" timestamp(0),
    "lat" numeric(10,8),
    "lng" numeric(11,8),
    PRIMARY KEY ("firm_id")
);

INSERT INTO "public"."firm" ("firm_id", "firm_name", "address", "province", "city", "since", "lat", "lng") VALUES
(1, 'Firma Hukum Mutiara Ekuator', 'Menara Tendean (M-Ten Jl. Kapten Tendean No.N0 20 C, Kuningan Bar., Kec. Mampang Prpt., Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12710', 'DKI Jakarta', 'Jakarta', NULL, -6.23930970, 106.82436093);
INSERT INTO "public"."firm" ("firm_id", "firm_name", "address", "province", "city", "since", "lat", "lng") VALUES
(2, 'Firma Hukum CITRALOKA CAHAYA', 'Jl. Andansari No.61Bandung, Sukorejo, LamonganKabupaten, Kabupaten Lamongan, Jawa Timur 62215', 'Jawa Timur', 'Lamongan', NULL, -7.12005542, 112.41292983);
INSERT INTO "public"."firm" ("firm_id", "firm_name", "address", "province", "city", "since", "lat", "lng") VALUES
(3, 'Pengacara BHP & Partners', 'Jl. Tegar Beriman No. 66, Pakansari, Cibinong, Pakansari, Kec. Cibinong, Kabupaten Bogor, Jawa Barat 16915', 'Jawa Barat', 'Kabupaten Bogor', NULL, -6.48017474, 106.84276943);
INSERT INTO "public"."firm" ("firm_id", "firm_name", "address", "province", "city", "since", "lat", "lng") VALUES
(4, 'Jasa Pengacara Ath Thaariq .S.H & Rekan', 'Jl. Komp. Ciriung Cemerlang No.77, RT.03/RW.03, Ciriung, Kec. Cibinong, Kabupaten Bogor, Jawa Barat 16918', 'Jawa Barat', 'Kabupaten Bogor', NULL, -6.47130531, 106.86542873),
(5, 'Kantor Hukum DEDEN SETIAWAN,SH & Rekan', 'Perum Graha Pandak Permai Blok L1 No.1, RT.7/RW.9, Karadenan, Kec. Cibinong, Kabupaten Bogor, Jawa Barat 16913', 'Jawa Barat', 'Kabupaten Bogor', NULL, -6.52534381, 106.81303173),
(6, 'Law Office Muhammad Vicky & Partners', 'Ruko Graha Cibinong, Jl. Raya Jakarta-Bogor No.KM, RW.43, Cirimekar, Kec. Cibinong, Kabupaten Bogor, Jawa Barat 16917', 'Jawa Barat', 'Kabupaten Bogor', NULL, -6.47028798, 106.84983655),
(7, 'Lawfirm MIM & Associates', 'Jl. Cipayung Barat No.25-22, Tengah, Kec. Cibinong, Kabupaten Bogor, Jawa Barat 16914', 'Jawa Barat', 'Kabupaten Bogor', NULL, -6.47310234, 106.82468816);