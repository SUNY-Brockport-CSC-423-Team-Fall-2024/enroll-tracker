CREATE OR REPLACE FUNCTION enroll_student (
    i_course_id int,
    i_student_id int
)
RETURNS void
AS $$
DECLARE
    major_course_match_exists BOOLEAN;
BEGIN
    SELECT EXISTS (
        SELECT 1
        FROM Student AS S
        JOIN Course_Major AS CM ON CM.major_id = S.major_id
        WHERE S.id = i_student_id 
        AND CM.course_id = i_course_id
    )
    INTO major_course_match_exists;
    
    IF NOT major_course_match_exists THEN
        RAISE EXCEPTION 'Enrollment failed: Course % is not part of students major', i_course_id;
    END IF;

    INSERT INTO Enrollments (course_id, student_id)
    VALUES
    (i_course_id, i_student_id);

        EXCEPTION
        WHEN unique_violation THEN
            RAISE EXCEPTION 'Student % enrolled in course %.', i_student_id, i_course_id;
        WHEN foreign_key_violation THEN
            RAISE EXCEPTION 'Foreign key violation: student % or course % does not exist.', i_student_id, i_course_id;
        WHEN not_null_violation THEN
            RAISE EXCEPTION 'Null value violation: all fields must be non-null.';
        WHEN check_violation THEN
            RAISE EXCEPTION 'Check constraint violation for student % or course %.', i_student_id, i_course_id;
        WHEN data_exception THEN
        RAISE EXCEPTION 'Data exception: invalid data type or format for student % or course %.', i_student_id, i_course_id;
END;
$$
LANGUAGE plpgsql;
