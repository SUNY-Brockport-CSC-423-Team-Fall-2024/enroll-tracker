CREATE OR REPLACE FUNCTION get_administrator (
    i_username varchar,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_office varchar,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
)
AS $$
BEGIN
    SELECT *
    INTO
        o_id,
        o_first_name,
        o_last_name,
        o_auth_id,
        o_phone_number,
        o_email,
        o_office,
        o_created_at,
        o_updated_at
    FROM Administrator
    INNER JOIN
        UserAuthentication ON Administrator.auth_id = UserAuthentication.id
    WHERE UserAuthentication.username = i_username
    AND UserAuthentication.is_active = true;

    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            RAISE EXCEPTION 'No admin found for username %s', i_username;
END;
$$
LANGUAGE plpgsql;
