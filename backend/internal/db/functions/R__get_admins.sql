CREATE OR REPLACE FUNCTION get_admins (
    i_limit int,
    i_offset int,
    i_first_name varchar,
    i_last_name varchar,
    i_username varchar,
    i_email varchar,
    i_phone_number varchar,
    i_office varchar
)
RETURNS SETOF Administrator
AS $$
BEGIN
    RETURN QUERY
    SELECT A.id, A.first_name, A.last_name, A.auth_id, A.phone_number, A.email, A.office, A.created_at, A.updated_at
    FROM Administrator AS A
    JOIN UserAuthentication AS UA ON UA.id = A.auth_id
    WHERE 
        UA.is_active = true
        AND
        (i_first_name IS NULL OR A.first_name LIKE '%' || i_first_name || '%')
        AND 
        (i_last_name IS NULL OR A.last_name LIKE '%' || i_last_name || '%')
        AND
        (i_email IS NULL OR A.email LIKE '%' || i_email || '%')
        AND
        (i_phone_number IS NULL OR A.phone_number LIKE '%' || i_phone_number || '%')
        AND
        (i_office IS NULL OR A.office LIKE '%' || i_office || '%')
        AND (i_username IS NULL OR UA.username LIKE '%' || i_username || '%')
    LIMIT COALESCE(i_limit, 10)
    OFFSET COALESCE(i_offset, 0);
END;
$$
LANGUAGE plpgsql;
