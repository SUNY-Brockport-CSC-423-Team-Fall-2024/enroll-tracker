CREATE OR REPLACE FUNCTION create_user_auth (
    username varchar,
    password_hash text
)
RETURNS VOID
AS $$
BEGIN
    INSERT INTO UserAuthentication (username, password_hash)
    VALUES (username, password_hash)
    RETURNING id INTO auth_id;
EXCEPTION 
    WHEN unique_violation THEN
        RAISE EXCEPTION 'Username already exists: %', username;
END;
$$
LANGUAGE plpgsql;
