CREATE OR REPLACE FUNCTION update_administrator (
    i_username varchar,
    i_first_name varchar DEFAULT NULL,
    i_last_name varchar DEFAULT NULL,
    i_phone_number varchar DEFAULT NULL,
    i_email varchar DEFAULT NULL,
    i_office varchar DEFAULT NULL
)
RETURNS TABLE (
    username varchar(60),
    id int,
    first_name varchar(50),
    last_name varchar(50),
    auth_id int,
    phone_number varchar(20),
    email varchar(50),
    office varchar(60),
    created_at timestamp,
    updated_at timestamp
)
AS $$
DECLARE
    admin_id int;
    admin_username varchar;
BEGIN
    SELECT get_admin.username, get_admin.id INTO admin_username, admin_id FROM get_admin(i_username);

    IF admin_id IS NULL THEN
        RAISE EXCEPTION 'No admin found with username: %s', i_username;
    END IF;
    
    RETURN QUERY
    UPDATE Administrator
        SET
            first_name = COALESCE(i_first_name, Administrator.first_name),
            last_name = COALESCE(i_last_name, Administrator.last_name),
            phone_number = COALESCE(i_phone_number, Administrator.phone_number),
            email = COALESCE(i_email, Administrator.email),
            office = COALESCE(i_office, Administrator.office)
        WHERE
            Administrator.id = admin_id
        RETURNING admin_username, Administrator.id, Administrator.first_name, Administrator.last_name, Administrator.auth_id, Administrator.phone_number, Administrator.email, Administrator.office, Administrator.created_at, Administrator.updated_at;
END;
$$
LANGUAGE plpgsql;
