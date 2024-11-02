CREATE OR REPLACE FUNCTION update_student (
    i_username varchar,
    i_first_name varchar DEFAULT NULL,
    i_last_name varchar DEFAULT NULL,
    i_phone_number varchar DEFAULT NULL,
    i_email varchar DEFAULT NULL,
    i_major_id int DEFAULT NULL,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_major_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
)
AS $$
DECLARE
    student_id int;
BEGIN
    SELECT get_student.o_id INTO student_id FROM get_student(i_username);

    IF student_id IS NULL THEN
        RAISE EXCEPTION 'No student found with username: %s', i_username;
    END IF;

    UPDATE Student
        SET
            first_name = COALESCE(i_first_name, first_name),
            last_name = COALESCE(i_last_name, last_name),
            phone_number = COALESCE(i_phone_number, phone_number),
            email = COALESCE(i_email, email),
            major_id = COALESCE(i_major_id, major_id)
        WHERE
            id = student_id
        RETURNING *
        INTO 
            o_id, 
            o_first_name,
            o_last_name,
            o_auth_id,
            o_major_id,
            o_phone_number,
            o_email,
            o_created_at,
            o_updated_at;
END;
$$
LANGUAGE plpgsql;
