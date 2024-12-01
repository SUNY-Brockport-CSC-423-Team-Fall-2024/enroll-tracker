CREATE OR REPLACE FUNCTION unenroll_students_after_course_dropped_from_major()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM Enrollments
    WHERE course_id = OLD.course_id
    AND student_id IN (
        SELECT S.id
        FROM Student AS S
        WHERE OLD.major_id = S.major_id
    );
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
