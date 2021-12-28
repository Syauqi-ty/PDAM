CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "nama" varchar,
  "phone" varchar,
  "NIK" varchar
);

CREATE TABLE "pointer" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "merge_id" varchar UNIQUE
);

CREATE TABLE "device" (
  "id" SERIAL PRIMARY KEY,
  "merge_id" varchar,
  "valve_status" int,
  "indikator_baterai" int,
  "koordinat" varchar
);

CREATE TABLE "alamat" (
  "id" SERIAL PRIMARY KEY,
  "merge_id" varchar,
  "jalan" varchar,
  "kelurahan" varchar,
  "kecamatan" varchar,
  "kabupaten" varchar,
  "koordinat" varchar UNIQUE
);

CREATE TABLE "finance" (
  "id" SERIAL PRIMARY KEY,
  "merge_id" varchar,
  "method" varchar,
  "status" int,
  "tanggal_payment" timestamp
);

CREATE TABLE "invoice" (
  "id" SERIAL PRIMARY KEY,
  "finance_id" int,
  "tagihan" varchar,
  "paid_at" timestamp
);

CREATE TABLE "volume" (
  "id" SERIAL PRIMARY KEY,
  "volume" varchar,
  "device_id" int,
  "created_at" timestamp 
);

ALTER TABLE "pointer" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "device" ADD FOREIGN KEY ("merge_id") REFERENCES "pointer" ("merge_id");

ALTER TABLE "alamat" ADD FOREIGN KEY ("merge_id") REFERENCES "pointer" ("merge_id");

ALTER TABLE "device" ADD FOREIGN KEY ("koordinat") REFERENCES "alamat" ("koordinat");

ALTER TABLE "finance" ADD FOREIGN KEY ("merge_id") REFERENCES "pointer" ("merge_id");

ALTER TABLE "invoice" ADD FOREIGN KEY ("finance_id") REFERENCES "finance" ("id");

ALTER TABLE "volume" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");
