CREATE OR REPLACE FUNCTION revoke_user_session (
    i_refresh_token_id text
)
RETURNS void
AS $$
BEGIN
    UPDATE UserSession
        SET
            revoked = true
        WHERE 
            refresh_token_id = i_refresh_token_id;
END;
$$
LANGUAGE plpgsql;
