CREATE OR REPLACE FUNCTION get_courses_students (
    i_course_id int,
    i_enrolled boolean DEFAULT NULL
)
RETURNS TABLE (
    student_id int,
    first_name varchar(50),
    last_name varchar(50),
    auth_id int,
    major_id int,
    phone_number varchar(20),
    email varchar(50),
    created_at timestamp,
    updated_at timestamp,
    is_enrolled boolean,
    enrolled_date timestamp,
    unenrolled_date timestamp
)
AS $$
BEGIN
    RETURN QUERY
    SELECT 
        S.id AS student_id,
        S.first_name,
        S.last_name,
        S.auth_id,
        S.major_id,
        S.phone_number,
        S.email,
        S.created_at,
        S.updated_at,
        E.is_enrolled,
        E.enrolled_date,
        E.unenrolled_date
    FROM Enrollments AS E
    JOIN Student AS S ON S.id = E.student_id
    WHERE 
        E.course_id = i_course_id
        AND (i_enrolled IS NULL OR E.is_enrolled = i_enrolled);
END;
$$
LANGUAGE plpgsql;
