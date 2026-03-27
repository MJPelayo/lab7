-- Create table

CREATE TABLE courses (
    id BIGSERIAL PRIMARY KEY,
    code TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    department TEXT NOT NULL,
    instructor TEXT NOT NULL,
    credits INT NOT NULL CHECK (credits > 0),
    capacity INT NOT NULL CHECK (capacity > 0),
    enrolled INT DEFAULT 0 CHECK (enrolled >= 0)
);