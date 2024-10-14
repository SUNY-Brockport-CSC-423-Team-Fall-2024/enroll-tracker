CREATE OR REPLACE FUNCTION create_student(
    first_name varchar,
    last_name varchar,
    phone_number varchar,
    email varchar,
    auth_id int
) 
RETURNS void
AS $$
BEGIN
    INSERT INTO Student (first_name, last_name, auth_id, major_id, phone_number, email, last_login)
    VALUES (first_name, last_name, auth_id, NULL, phone_number, email, NULL);
END;
$$
LANGUAGE plpgsql;
