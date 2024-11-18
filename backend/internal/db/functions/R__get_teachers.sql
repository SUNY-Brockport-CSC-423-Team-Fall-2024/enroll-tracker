CREATE OR REPLACE FUNCTION get_teachers (
    i_limit int,
    i_offset int,
    i_first_name varchar,
    i_last_name varchar,
    i_username varchar,
    i_email varchar,
    i_phone_number varchar,
    i_office varchar
)
RETURNS TABLE (
    username varchar(60),
    id int,
    first_name varchar(50),
    last_name varchar(50),
    auth_id int,
    phone_number varchar(20),
    email varchar(50),
    office varchar(60),
    created_at timestamp,
    updated_at timestamp
)
AS $$
BEGIN
    RETURN QUERY
    SELECT UA.username, T.id, T.first_name, T.last_name, T.auth_id, T.phone_number, T.email, T.office, T.created_at, T.updated_at
    FROM Teacher AS T
    JOIN UserAuthentication AS UA ON UA.id = T.auth_id
    WHERE 
        UA.is_active = true
        AND
        (i_first_name IS NULL OR T.first_name LIKE '%' || i_first_name || '%')
        AND 
        (i_last_name IS NULL OR T.last_name LIKE '%' || i_last_name || '%')
        AND
        (i_email IS NULL OR T.email LIKE '%' || i_email || '%')
        AND
        (i_phone_number IS NULL OR T.phone_number LIKE '%' || i_phone_number || '%')
        AND
        (i_office IS NULL OR T.office LIKE '%' || i_office || '%')
        AND (i_username IS NULL OR UA.username LIKE '%' || i_username || '%')
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
