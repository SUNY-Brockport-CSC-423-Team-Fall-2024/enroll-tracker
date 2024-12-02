CREATE OR REPLACE FUNCTION update_student (
    i_username varchar,
    i_first_name varchar DEFAULT NULL,
    i_last_name varchar DEFAULT NULL,
    i_phone_number varchar DEFAULT NULL,
    i_email varchar DEFAULT NULL
)
RETURNS TABLE (
    username varchar(60),
    id int,
    first_name varchar(50),
    last_name varchar(50),
    auth_id int,
    major_id int,
    phone_number varchar(20),
    email varchar(50),
    created_at timestamp,
    updated_at timestamp
)
AS $$
DECLARE
    student_id int;
    student_username varchar;
BEGIN
    SELECT get_student.username, get_student.id INTO student_username, student_id FROM get_student(i_username);

    IF student_id IS NULL THEN
        RAISE EXCEPTION 'No student found with username: %s', i_username;
    END IF;
    
    RETURN QUERY
    UPDATE Student
        SET
            first_name = COALESCE(i_first_name, Student.first_name),
            last_name = COALESCE(i_last_name, Student.last_name),
            phone_number = COALESCE(i_phone_number, Student.phone_number),
            email = COALESCE(i_email, Student.email)
        WHERE
            Student.id = student_id
        RETURNING student_username, Student.id, Student.first_name, Student.last_name, Student.auth_id, Student.major_id, Student.phone_number, Student.email, Student.created_at, Student.updated_at;
END;
$$
LANGUAGE plpgsql;
