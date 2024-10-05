CREATE TABLE IF NOT EXISTS orders(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL REFERENCES users("id") ON DELETE CASCADE,
    "total" DECIMAL(10,2) NOT NULL,
    "status" VARCHAR(50) NOT NULL CHECK ("status" IN ('pending', 'completed', 'cancelled')),
    "address" TEXT NOT NULL,
    "quantity" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);