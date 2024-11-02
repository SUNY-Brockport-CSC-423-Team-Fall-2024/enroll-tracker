CREATE OR REPLACE FUNCTION create_user_auth (
    input_username varchar,
    input_password_hash text,
    OUT created_id int,
    OUT created_username varchar,
    OUT created_password_hash text,
    OUT created_last_login timestamp,
    OUT created_is_active boolean
)
AS $$
BEGIN
    INSERT INTO UserAuthentication (username, password_hash, last_login)
    VALUES (input_username, input_password_hash, NULL)
    RETURNING id, username, password_hash, last_login, is_active
    INTO created_id, created_username, created_password_hash, created_last_login, created_is_active;
EXCEPTION 
    WHEN unique_violation THEN
        RAISE EXCEPTION 'Username already exists: %', username;
END;
$$
LANGUAGE plpgsql;
