CREATE OR REPLACE FUNCTION create_student (
    i_first_name varchar,
    i_last_name varchar,
    i_auth_id int,
    i_phone_number varchar,
    i_email varchar
)
RETURNS TABLE(
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
BEGIN
    SELECT UA.username
    INTO username
    FROM UserAuthentication AS UA
    WHERE UA.id = i_auth_id;
    
    RETURN QUERY
    INSERT INTO Student (first_name, last_name, auth_id, major_id, phone_number, email)
    VALUES (i_first_name, i_last_name, i_auth_id, NULL, i_phone_number, i_email)
    RETURNING username, Student.id, Student.first_name, Student.last_name, Student.auth_id, Student.major_id, Student.phone_number, Student.email, Student.created_at, Student.updated_at;
END;
$$
LANGUAGE plpgsql;
