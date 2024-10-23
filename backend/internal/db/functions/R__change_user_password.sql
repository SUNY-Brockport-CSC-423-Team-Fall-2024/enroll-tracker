CREATE OR REPLACE FUNCTION change_user_password (
    i_user_auth_id int,
    i_new_password_digest text,
    OUT o_id int,
    OUT o_username varchar,
    OUT o_password_hash text,
    OUT o_last_password_reset timestamp
)
AS $$
BEGIN
    UPDATE UserAuthentication
        SET
            password_hash = i_new_password_digest,
            last_password_reset = CURRENT_TIMESTAMP
    WHERE 
        id = i_user_auth_id
    RETURNING *
    INTO
        o_id,
        o_username,
        o_password_hash,
        o_last_password_reset;
END;
$$
LANGUAGE plpgsql;
