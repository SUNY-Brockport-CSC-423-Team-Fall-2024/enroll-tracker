CREATE OR REPLACE FUNCTION get_user_session(
    i_refresh_token_id text,
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
    SELECT *
    INTO
        o_id,
        o_user_id,
        o_username,
        o_refresh_token,
        o_refresh_token_id,
        o_issued_at,
        o_expires_at,
        o_revoked
    FROM UserSession
    WHERE UserSession.refresh_token_id = i_refresh_token_id;
    
    IF NOT FOUND THEN
        RAISE EXCEPTION 'No session found for refresh token id %s', get_user_session.i_refresh_token_id;
    END IF;
END;
$$
LANGUAGE plpgsql;
