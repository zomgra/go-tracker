CREATE TABLE IF NOT EXISTS shipments (
    Id INT PRIMARY KEY serial,
    Barcode varchar(24) unique NOT NULL
    CreateAt TIMESTAMP NULL
)