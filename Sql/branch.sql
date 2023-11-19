CREATE TABLE branch(
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" VARCHAR,
    "phone" VARCHAR,
    "photo" VARCHAR,
    "work_start_hour" VARCHAR,
    "work_end_hour" VARCHAR,
    "address" VARCHAR,
    "delivery_price" INT DEFAULT 100000,
    "status" VARCHAR NOT NULL CHECK ("status"=ANY ('{ACTIVE,NO_ACTIVE}'::VARCHAR[])) ,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
)
;


INSERT INTO branch (name,phone,photo,work_start_hour,work_end_hour,address,status) VALUES 
('Samandarxon','+99894-673-22-77','https://img.freepik.com/premium-photo/young-handsome-man-with-beard-isolated-keeping-arms-crossed-frontal-position_1368-132662.jpg','08:00:00','19:30:00', 'Qashqadaryo','sam');

