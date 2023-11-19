CREATE DATABASE market_app;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE category (
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "title" VARCHAR ,
    "image" VARCHAR,
    "parent_id" UUID REFERENCES "category"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

INSERT INTO "category"("title","image","parent_id","updated_at") VALUES
('MEVA','https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Culinary_fruits_front_view.jpg/500px-Culinary_fruits_front_view.jpg',NULL,NOW())RETURNING "id";




