CREATE OR REPLACE FUNCTION get_teachers_courses (
    i_teacher_id int
)
RETURNS TABLE (
    course_id int,
    course_name varchar,
    course_description varchar,
    teacher_id int,
    current_enrollment bigint,
    max_enrollment int,
    num_credits int,
    status CourseStatus,
    last_updated timestamp,
    created_at timestamp
)
AS $$
BEGIN
    RETURN QUERY
    SELECT
        C.id AS course_id,
        C.name AS course_name,
        C.description AS course_description,
        C.teacher_id,
        COALESCE(COUNT(E.course_id), 0) AS current_enrollment,
        C.max_enrollment,
        C.num_credits,
        C.status,
        C.last_updated,
        C.created_at
    FROM Course AS C
    LEFT JOIN Enrollments AS E ON E.course_id = C.id AND E.is_enrolled = TRUE
    WHERE 
        C.teacher_id = i_teacher_id
    GROUP BY
        C.id, C.name, C.description, C.teacher_id, C.max_enrollment, C.num_credits, C.status, C.last_updated, C.created_at;
END;
$$
LANGUAGE plpgsql;
