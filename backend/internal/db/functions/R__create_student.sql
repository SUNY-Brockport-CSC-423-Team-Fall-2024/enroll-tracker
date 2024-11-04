CREATE OR REPLACE FUNCTION create_student(
    i_first_name varchar,
    i_last_name varchar,
    i_auth_id int,
    i_phone_number varchar,
    i_email varchar,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
) 
AS $$
BEGIN
    INSERT INTO Student (first_name, last_name, auth_id, major_id, phone_number, email)
    VALUES (i_first_name, i_last_name, i_auth_id, NULL, i_phone_number, i_email)
    RETURNING id, first_name, last_name, auth_id, phone_number, email, created_at, updated_at
    INTO o_id, o_first_name, o_last_name, o_auth_id, o_phone_number, o_email, o_created_at, o_updated_at;
END;
$$
LANGUAGE plpgsql;
