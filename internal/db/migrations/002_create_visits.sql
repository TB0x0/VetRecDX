CREATE TABLE IF NOT EXISTS visits (
    id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    visit_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    notes TEXT,
    diagnosis TEXT,
    weight_lbs NUMERIC(5,2),
    temperature_f NUMERIC(4,1),
    follow_up_date DATE
);