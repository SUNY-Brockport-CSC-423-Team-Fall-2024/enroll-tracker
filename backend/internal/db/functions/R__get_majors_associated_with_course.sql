CREATE OR REPLACE FUNCTION get_majors_associated_with_course (
    i_course_id int,
    i_limit int,
    i_offset int,
    i_name varchar,
    i_description varchar,
    i_status MajorStatus
)
RETURNS SETOF Major
AS $$
BEGIN
    RETURN QUERY
    SELECT
        M.id,
        M.name,
        M.description,
        M.status,
        M.last_updated,
        M.created_at
    FROM Course_Major AS CM
    JOIN Major AS M ON M.id = CM.major_id
    WHERE
        CM.course_id = i_course_id
        AND (i_name IS NULL OR M.name LIKE '%' || i_name || '%')
        AND (i_description IS NULL OR M.description LIKE '%' || i_description || '%')
        AND (i_status IS NULL OR M.status = i_status)
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
