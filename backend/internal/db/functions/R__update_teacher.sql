CREATE OR REPLACE FUNCTION update_teacher (
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
    teacher_id int;
BEGIN
    SELECT get_teacher.o_id INTO teacher_id FROM get_teacher(i_username);

    IF teacher_id IS NULL THEN
        RAISE EXCEPTION 'No teacher found with username: %s', i_username;
    END IF;

    UPDATE Teacher
        SET
            first_name = COALESCE(i_first_name, first_name),
            last_name = COALESCE(i_last_name, last_name),
            phone_number = COALESCE(i_phone_number, phone_number),
            email = COALESCE(i_email, email),
            office = COALESCE(i_office, email)
        WHERE
            id = teacher_id
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
