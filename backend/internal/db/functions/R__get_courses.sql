CREATE OR REPLACE FUNCTION get_courses (
    i_limit int,
    i_offset int,
    i_name varchar,
    i_description varchar,
    i_teacher_id int,
    i_max_enrollment int,
    i_min_enrollment int,
    i_max_num_credits int,
    i_min_num_credits int,
    i_status CourseStatus
)
RETURNS TABLE (
    id int,
    name varchar(50),
    description varchar(255),
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
        C.id,
        C.name,
        C.description,
        C.teacher_id,
        COALESCE(COUNT(E.course_id), 0) AS "current_enrollment",
        C.max_enrollment,
        C.num_credits,
        C.status,
        C.last_updated,
        C.created_at
    FROM Course AS C
    LEFT JOIN Enrollments AS E ON E.course_id = C.id AND E.is_enrolled = TRUE
    WHERE
        (i_name IS NULL OR C.name LIKE '%' || i_name || '%')
        AND (i_description IS NULL OR C.description LIKE '%' || i_description || '%')
        AND (i_teacher_id IS NULL OR C.teacher_id = i_teacher_id)
        AND (i_max_enrollment IS NULL OR C.max_enrollment <= i_max_enrollment)
        AND (i_min_enrollment IS NULL OR C.max_enrollment >= i_min_enrollment)
        AND (i_max_num_credits IS NULL OR C.num_credits <= i_max_num_credits)
        AND (i_min_num_credits IS NULL OR C.num_credits >= i_min_num_credits)
        AND (i_status IS NULL OR C.status = i_status)
    GROUP BY C.id, C.name, C.description, C.teacher_id, C.max_enrollment, C.num_credits, C.status, C.last_updated, C.created_at
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
