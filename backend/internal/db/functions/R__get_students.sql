CREATE OR REPLACE FUNCTION get_students (
    i_limit int,
    i_offset int,
    i_first_name varchar,
    i_last_name varchar,
    i_username varchar,
    i_majors varchar[],
    i_email varchar,
    i_phone_number varchar
)
RETURNS SETOF Student
AS $$
BEGIN
    RETURN QUERY
    SELECT *
    FROM Student AS S
    WHERE 
        (i_first_name IS NULL OR S.first_name LIKE '%' || i_first_name || '%')
        AND 
        (i_last_name IS NULL OR S.last_name LIKE '%' || i_last_name || '%')
        AND
        (i_email IS NULL OR S.email LIKE '%' || i_email || '%')
        AND
        (i_phone_number IS NULL OR S.phone_number LIKE '%' || i_phone_number || '%')
        AND (i_majors IS NULL OR EXISTS ( -- If majors provided. For each major, grab records where the students major id matches
            SELECT 1
            FROM UNNEST(i_majors) AS major_name
            WHERE EXISTS (
                SELECT 1
                FROM Major AS M
                WHERE M.name LIKE '%' || major_name || '%' 
                AND M.id = S.major_id
            )
        ))
        AND (i_username IS NULL OR EXISTS ( -- If username provided. Grab entries where the users username contains the parameter
            SELECT 1
            FROM UserAuthentication AS UA
            WHERE UA.username LIKE '%' || i_username || '%'
            AND UA.id = S.auth_id
        ))
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
