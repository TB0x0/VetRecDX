CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    species TEXT NOT NULL,
    breed TEXT,
    gender TEXT CHECK (gender IN ('Male', 'Female', 'Unknown')),
    date_of_birth DATE,
    flags TEXT,
    owner_name TEXT NOT NULL,
    owner_contact_phone_hash TEXT,
    owner_contact_email_hash TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);