CREATE OR REPLACE FUNCTION update_last_login (
    i_id int,
    OUT o_id int,
    OUT o_username varchar,
    OUT o_password_hash text,
    OUT o_last_login timestamp,
    OUT o_last_password_reset timestamp
)
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
