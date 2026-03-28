-- 1. Create table

CREATE TABLE enrollments (
    id BIGSERIAL PRIMARY KEY,
    student_id BIGINT REFERENCES students(id) ON DELETE CASCADE,
    course_id BIGINT REFERENCES courses(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMPTZ DEFAULT NOW()
);

-- 2. Index (for faster searching)
CREATE INDEX idx_enrollments_student_id ON enrollments (student_id);
CREATE INDEX idx_enrollments_course_id ON enrollments (course_id);

-- 3. Seed data (initial test data)
INSERT INTO enrollments (student_id, course_id) VALUES
(1, 1), -- John Doe enrolls in CS101
(1, 2), -- John Doe enrolls in IS201
(2, 1), -- Alice Smith enrolls in CS101
(3, 3); -- Bob Lee enrolls in CS202


