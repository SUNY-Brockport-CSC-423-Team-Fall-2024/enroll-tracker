CREATE OR REPLACE FUNCTION create_teacher (
    i_first_name varchar,
    i_last_name varchar,
    i_auth_id int,
    i_phone_number varchar,
    i_email varchar,
    i_office varchar
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
BEGIN
    SELECT UA.username
    INTO username
    FROM UserAuthentication AS UA
    WHERE UA.id = i_auth_id;
    
    RETURN QUERY
    INSERT INTO Teacher (first_name, last_name, auth_id, phone_number, email, office)
    VALUES (i_first_name, i_last_name, i_auth_id, i_phone_number, i_email, i_office)
    RETURNING username, Teacher.id, Teacher.first_name, Teacher.last_name, Teacher.auth_id, Teacher.phone_number, Teacher.email, Teacher.office, Teacher.created_at, Teacher.updated_at;
END;
$$
LANGUAGE plpgsql;
