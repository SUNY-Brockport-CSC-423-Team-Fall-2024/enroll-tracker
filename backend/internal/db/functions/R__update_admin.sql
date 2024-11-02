CREATE OR REPLACE FUNCTION update_administrator (
    i_username varchar,
    i_first_name varchar DEFAULT NULL,
    i_last_name varchar DEFAULT NULL,
    i_phone_number varchar DEFAULT NULL,
    i_email varchar DEFAULT NULL,
    i_office varchar DEFAULT NULL,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_office varchar,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
)
AS $$
DECLARE
    admin_id int;
BEGIN
    SELECT get_admin.o_id INTO admin_id FROM get_admin(i_username);

    IF admin_id IS NULL THEN
        RAISE EXCEPTION 'No admin found with username: %s', i_username;
    END IF;

    UPDATE Administrator
        SET
            first_name = COALESCE(i_first_name, first_name),
            last_name = COALESCE(i_last_name, last_name),
            phone_number = COALESCE(i_phone_number, phone_number),
            email = COALESCE(i_email, email),
            office = COALESCE(i_office, email)
        WHERE
            id = admin_id
        RETURNING *
        INTO 
            o_id, 
            o_first_name,
            o_last_name,
            o_auth_id,
            o_phone_number,
            o_email,
            o_office,
            o_created_at,
            o_updated_at;
END;
$$
LANGUAGE plpgsql;
