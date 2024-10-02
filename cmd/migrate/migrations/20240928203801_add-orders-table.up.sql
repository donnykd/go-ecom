CREATE TABLE IF NOT EXISTS orders(
    "id" SERIAL PRIMARY KEY,
    "userId" INTEGER NOT NULL REFERENCES users("id") ON DELETE CASCADE,
    "total" DECIMAL(10,2) NOT NULL,
    "status" VARCHAR(50) NOT NULL CHECK ("status" IN ('pending', 'completed', 'cancelled')),
    "address" TEXT NOT NULL,
    "quantity" INTEGER NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);