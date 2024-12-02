CREATE OR REPLACE FUNCTION update_teacher (
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
    teacher_id int;
    teacher_username varchar;
BEGIN
    SELECT get_teacher.username, get_teacher.id INTO teacher_username, teacher_id FROM get_teacher(i_username);

    IF teacher_id IS NULL THEN
        RAISE EXCEPTION 'No teacher found with username: %s', i_username;
    END IF;

    RETURN QUERY
    UPDATE Teacher
        SET
            first_name = COALESCE(i_first_name, Teacher.first_name),
            last_name = COALESCE(i_last_name, Teacher.last_name),
            phone_number = COALESCE(i_phone_number, Teacher.phone_number),
            email = COALESCE(i_email, Teacher.email),
            office = COALESCE(i_office, Teacher.office)
        WHERE
            Teacher.id = teacher_id
        RETURNING teacher_username, Teacher.id, Teacher.first_name, Teacher.last_name, Teacher.auth_id, Teacher.phone_number, Teacher.email, Teacher.office, Teacher.created_at, Teacher.updated_at;
END;
$$
LANGUAGE plpgsql;
