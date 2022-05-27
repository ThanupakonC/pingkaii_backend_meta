CREATE TABLE "region" (
  "region_id" int PRIMARY KEY NOT NULL,
  "region_thai" varchar NOT NULL,
  "region_eng" varchar NOT NULL
);

CREATE TABLE "province" (
  "province_id" int PRIMARY KEY NOT NULL,
  "province_thai" varchar NOT NULL,
  "province_eng" varchar NOT NULL,
  "region_id" int
);

CREATE TABLE "district" (
  "district_id" int PRIMARY KEY NOT NULL,
  "district_thai" varchar NOT NULL,
  "district_eng" varchar NOT NULL,
  "district_thai_short" varchar NOT NULL,
  "district_eng_short" varchar NOT NULL,
  "province_id" int NOT NULL,
  "region_id" int NOT NULL
);

CREATE TABLE "tambon" (
  "tambon_id" int PRIMARY KEY NOT NULL,
  "tambon_thai" varchar NOT NULL,
  "tambon_eng" varchar NOT NULL,
  "tambon_thai_short" varchar NOT NULL,
  "tambon_thai_eng" varchar NOT NULL,
  "district_id" int NOT NULL,
  "province_id" int NOT NULL,
  "region_id" int NOT NULL
);

CREATE TABLE "postcode" (
  "postcode" int NOT NULL,
  "tambon_id" int NOT NULL,
  PRIMARY KEY ("postcode", "tambon_id")
);

ALTER TABLE "province" ADD FOREIGN KEY ("region_id") REFERENCES "region" ("region_id");

ALTER TABLE "district" ADD FOREIGN KEY ("province_id") REFERENCES "province" ("province_id");

ALTER TABLE "district" ADD FOREIGN KEY ("region_id") REFERENCES "region" ("region_id");

ALTER TABLE "tambon" ADD FOREIGN KEY ("district_id") REFERENCES "district" ("district_id");

ALTER TABLE "tambon" ADD FOREIGN KEY ("province_id") REFERENCES "province" ("province_id");

ALTER TABLE "tambon" ADD FOREIGN KEY ("region_id") REFERENCES "region" ("region_id");

ALTER TABLE "postcode" ADD FOREIGN KEY ("tambon_id") REFERENCES "tambon" ("tambon_id");
