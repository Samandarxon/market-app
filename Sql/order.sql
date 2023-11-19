
CREATE TABLE "order" (
    "id" UUID NOT NULL PRIMARY KEY,
    "order_id"  VARCHAR(255) NOT NULL UNIQUE ,
    "client_id" uuid not null REFERENCES "client"("id"),
    "branch_id" uuid NOT NULL REFERENCES "branches"("id"), 
    "address" VARCHAR(255),
    "delivery_price" NUMERIC, 
    "total_count" INT,
    "total_price" NUMERIC,
    "status" VARCHAR(20) NOT NULL DEFAULT 'new',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE "order_products" (
    "order_product_id" UUID NOT NULL PRIMARY KEY,
    "order_id" UUID NOT NULL REFERENCES "order"("id"),
    "product_id" UUID NOT NULL REFERENCES "product"("id"),
    "discount_type" VARCHAR(20) ,
    "discount_amount" NUMERIC ,
    "quantity" NUMERIC NOT NULL,
    "price" NUMERIC NOT NULL,
    "sum" NUMERIC NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);