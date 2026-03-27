-- Create table
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    programme TEXT NOT NULL,
    year SMALLINT NOT NULL CHECK (year BETWEEN 1 AND 4),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. Index (for faster searching)
CREATE INDEX idx_students_name ON students (name);

-- 4. Seed data (initial test data)
INSERT INTO students (name, programme, year) VALUES
    ('Eve Castillo',   'BSc Computer Science',    2),
    ('Marco Tillett',  'BSc Computer Science',    3),
    ('Aisha Gentle',   'BSc Information Systems', 1),
    ('Raj Palacio',    'BSc Computer Science',    4);