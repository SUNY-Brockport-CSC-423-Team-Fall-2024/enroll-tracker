CREATE OR REPLACE FUNCTION get_majors (
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
    SELECT *
    FROM Major AS M
    WHERE
        (i_name IS NULL OR M.name LIKE '%' || i_name || '%')
        AND (i_description IS NULL OR M.description LIKE '%' || i_description || '%')
        AND (i_status IS NULL OR M.status = i_status)
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
