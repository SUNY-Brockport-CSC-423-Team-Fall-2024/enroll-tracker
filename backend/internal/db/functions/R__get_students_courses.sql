CREATE OR REPLACE FUNCTION get_students_courses (
    i_student_id int,
    i_enrolled boolean DEFAULT NULL
)
RETURNS TABLE (
    course_id int,
    course_name varchar,
    course_description varchar,
    teacher_id int,
    max_enrollment int,
    num_credits int,
    status CourseStatus,
    last_updated timestamp,
    created_at timestamp,
    is_enrolled boolean,
    unenrolled_date timestamp,
    enrolled_date timestamp
)
AS $$
BEGIN
    RETURN QUERY
    SELECT
        C.id AS course_id,
        C.name AS course_name,
        C.description AS course_description,
        C.teacher_id,
        C.max_enrollment,
        C.num_credits,
        C.status,
        C.last_updated,
        C.created_at,
        E.is_enrolled,
        E.unenrolled_date,
        E.enrolled_date
    FROM Enrollments AS E
    JOIN Course AS C ON C.id = E.course_id
    WHERE 
        E.student_id = i_student_id
        AND (i_enrolled IS NULL OR E.is_enrolled = i_enrolled);
END;
$$
LANGUAGE plpgsql;
