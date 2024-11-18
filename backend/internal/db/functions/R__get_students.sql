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
RETURNS TABLE (
    username varchar(60),
    id int,
    first_name varchar(50),
    last_name varchar(50),
    auth_id int,
    major_id int,
    phone_number varchar(20),
    email varchar(50),
    created_at timestamp,
    updated_at timestamp
)
AS $$
BEGIN
    RETURN QUERY
    SELECT UA.username, S.id, S.first_name, S.last_name, S.auth_id, S.major_id, S.phone_number, S.email, S.created_at, S.updated_at
    FROM Student AS S
    JOIN UserAuthentication AS UA ON UA.id = S.auth_id
    WHERE 
        UA.is_active = true
        AND
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
        AND (i_username IS NULL OR UA.username LIKE '%' || i_username || '%')
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
