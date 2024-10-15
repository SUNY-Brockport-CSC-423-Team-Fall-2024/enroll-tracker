CREATE OR REPLACE FUNCTION create_user_auth (
    input_username varchar,
    input_password_hash text,
    OUT created_id int,
    OUT created_username varchar,
    OUT created_password_hash text
)
AS $$
BEGIN
    INSERT INTO UserAuthentication (username, password_hash)
    VALUES (input_username, input_password_hash)
    RETURNING id, username, password_hash
    INTO created_id, created_username, created_password_hash;
EXCEPTION 
    WHEN unique_violation THEN
        RAISE EXCEPTION 'Username already exists: %', username;
END;
$$
LANGUAGE plpgsql;
