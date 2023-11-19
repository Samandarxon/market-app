CREATE TABLE product (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "category_id" UUID NOT NULL REFERENCES "category"("id"),
    "title" VARCHAR ,
    "description" VARCHAR,
    "photo" VARCHAR,
    "price" NUMERIC,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);