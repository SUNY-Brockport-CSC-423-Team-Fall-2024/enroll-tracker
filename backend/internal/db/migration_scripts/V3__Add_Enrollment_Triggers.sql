-- Course deletion
CREATE OR REPLACE FUNCTION unenroll_students_if_course_inactive()
RETURNS TRIGGER AS $$
BEGIN
    --Check if course is being set to inactive
    IF NEW.status = 'inactive' THEN
        UPDATE Enrollments 
        SET
            is_enrolled = false,
            unenrolled_date = CURRENT_TIMESTAMP
        WHERE 
            course_id = NEW.id;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER course_status_update
BEFORE UPDATE ON Course
FOR EACH ROW
WHEN (OLD.status IS DISTINCT FROM NEW.status)
EXECUTE FUNCTION unenroll_students_if_course_inactive();

-- Student deletion
CREATE OR REPLACE FUNCTION unenroll_student_if_inactive()
RETURNS TRIGGER AS $$
DECLARE
    found_student_id int;
BEGIN
    SELECT Student.id
    INTO found_student_id
    FROM Student
    WHERE Student.auth_id = NEW.id;

    --If a student is found, update their enrollments
    IF found_student_id IS NOT NULL AND NEW.is_active = false THEN
        UPDATE Enrollments 
        SET
            is_enrolled = false,
            unenrolled_date = CURRENT_TIMESTAMP
        WHERE 
            student_id = found_student_id;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER unenroll_inactive_student
BEFORE UPDATE ON UserAuthentication
FOR EACH ROW
WHEN 
    (OLD.is_active IS DISTINCT FROM NEW.is_active)
EXECUTE FUNCTION unenroll_student_if_inactive();

-- Teacher deletion
CREATE OR REPLACE FUNCTION inactivate_course_if_teacher_inactive()
RETURNS TRIGGER AS $$
DECLARE
    found_teacher_id int;
BEGIN
    SELECT Teacher.id
    INTO found_teacher_id
    FROM Teacher
    WHERE Teacher.auth_id = NEW.id;

    --If a teacher is found, update their enrollments
    IF found_teacher_id IS NOT NULL AND NEW.is_active = false THEN
        UPDATE Course 
        SET
            status = 'inactive',
            max_enrollment = 0,
            last_updated = CURRENT_TIMESTAMP
        WHERE 
            teacher_id = found_teacher_id;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER inactivate_course_when_teacher_inactive
BEFORE UPDATE ON UserAuthentication
FOR EACH ROW
WHEN 
    (OLD.is_active IS DISTINCT FROM NEW.is_active)
EXECUTE FUNCTION inactivate_course_if_teacher_inactive();
