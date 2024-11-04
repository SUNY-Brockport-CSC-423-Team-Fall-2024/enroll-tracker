CREATE OR REPLACE FUNCTION update_last_login (
    i_id int
)
RETURNS void
AS $$
BEGIN
    Update UserAuthentication
    Set
        last_login = CURRENT_TIMESTAMP
    WHERE
        id = i_id;
END;
$$
LANGUAGE plpgsql;
