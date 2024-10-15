CREATE OR REPLACE FUNCTION get_student (
    username varchar,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_major_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_last_login timestamp,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
)
AS $$
BEGIN
    SELECT 
        Student.id,
        Student.first_name,
        Student.last_name,
        Student.auth_id,
        Student.major_id,
        Student.phone_number,
        Student.email,
        Student.last_login,
        Student.created_at,
        Student.updated_at
    INTO 
        o_id,
        o_first_name,
        o_last_name,
        o_auth_id,
        o_major_id,
        o_phone_number,
        o_email,
        o_last_login,
        o_created_at,
        o_updated_at
    FROM Student
    INNER JOIN 
        UserAuthentication ON Student.auth_id = UserAuthentication.id
    WHERE UserAuthentication.username = get_student.username;
    
    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            RAISE EXCEPTION 'No student found for username %s', get_student.username;
END;
$$
LANGUAGE plpgsql;
