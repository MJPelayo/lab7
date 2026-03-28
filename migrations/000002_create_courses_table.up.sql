-- 1. Create table

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

-- 2. Index (for faster searching)
CREATE INDEX idx_courses_code ON courses (code);

-- 3. Seed data (initial test data)
INSERT INTO courses (code, title, department, instructor, credits, capacity) VALUES
('CS101', 'Introduction to Computer Science', 'Computer Science', 'Dr. Smith', 3, 30),
('IS201', 'Information Systems Analysis', 'Information Systems', 'Dr. Johnson', 4, 25),
('CS202', 'Data Structures', 'Computer Science', 'Dr. Lee', 3, 30);     
