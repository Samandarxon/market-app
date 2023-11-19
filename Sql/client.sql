CREATE TABLE client(
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "first_name" VARCHAR(25),
    "last_name" VARCHAR(25),
    "phone" VARCHAR(20),
    "photo" VARCHAR,
    "date_of_birth" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);