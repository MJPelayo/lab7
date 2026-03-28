-- 1. Create table
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    programme TEXT NOT NULL,
    year SMALLINT NOT NULL CHECK (year BETWEEN 1 AND 4),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. Index (for faster searching)
CREATE INDEX idx_students_name ON students (name);

-- 3. Seed data (initial test data)
INSERT INTO students (name, programme, year) VALUES
('John Doe','Computer Science',2),
('Alice Smith','Information Systems',1),
('Bob Lee','Computer Science',3);