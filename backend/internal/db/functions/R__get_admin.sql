CREATE OR REPLACE FUNCTION get_admin (
    i_username varchar
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
    SELECT UA.username, A.*
    FROM Administrator AS A
    INNER JOIN
        UserAuthentication AS UA ON A.auth_id = UA.id
    WHERE UA.username = i_username
    AND UA.is_active = true;

    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            RAISE EXCEPTION 'No administrator found for username %s', i_username;
END;
$$
LANGUAGE plpgsql;
