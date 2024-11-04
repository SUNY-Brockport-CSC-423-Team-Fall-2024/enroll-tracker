CREATE OR REPLACE FUNCTION delete_user_auth (
    i_username varchar
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    UPDATE UserAuthentication
    SET
        is_active = false
    WHERE 
        username = i_username;

    GET DIAGNOSTICS row_cnt = ROW_COUNT;

    RETURN row_cnt > 0;
END;
$$
LANGUAGE plpgsql;
