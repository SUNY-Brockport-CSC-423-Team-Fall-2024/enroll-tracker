CREATE OR REPLACE FUNCTION revoke_user_sessions_with_username (
    i_username text
)
RETURNS void
AS $$
BEGIN
    UPDATE UserSession
        SET
            revoked = true
        WHERE 
            username = i_username;
END;
$$
LANGUAGE plpgsql;
