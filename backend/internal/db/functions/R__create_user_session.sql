CREATE OR REPLACE FUNCTION create_user_session(
    i_user_id   int,
    i_username varchar,
    i_refresh_token   text,
    i_refresh_token_id text,
    i_issued_at timestamp,
    i_expires_at timestamp,
    OUT o_id int,
    OUT o_user_id int,
    OUT o_username varchar,
    OUT o_refresh_token text,
    OUT o_refresh_token_id text,
    OUT o_issued_at timestamp,
    OUT o_expires_at timestamp,
    OUT o_revoked boolean
)
AS $$
BEGIN
    INSERT INTO UserSession (user_id, username, refresh_token, refresh_token_id, issued_at, expires_at)
    VALUES (i_user_id, i_username, i_refresh_token, i_refresh_token_id, i_issued_at, i_expires_at)
    RETURNING *
    INTO
        o_id,
        o_user_id,
        o_username,
        o_refresh_token,
        o_refresh_token_id,
        o_issued_at,
        o_expires_at,
        o_revoked;
END;
$$
LANGUAGE plpgsql;
