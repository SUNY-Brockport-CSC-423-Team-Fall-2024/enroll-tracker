CREATE OR REPLACE FUNCTION get_user_role (
    i_username varchar,
    OUT o_role varchar
)
AS $$
DECLARE
    user_id int;
BEGIN
    -- Get the user_id from the UserAuthentication table
    SELECT id 
    INTO user_id
    FROM UserAuthentication
    WHERE username = i_username;

    -- Check if the user exists in the Student table
        SELECT 'student'
        INTO o_role
        FROM Student
        WHERE Student.auth_id = user_id;
        
        IF FOUND THEN
            RETURN;
        END IF;

    -- Check if the user exists in the Teacher table
        SELECT 'teacher'
        INTO o_role
        FROM Teacher
        WHERE Teacher.auth_id = user_id;
        
        IF FOUND THEN
            RETURN;
        END IF;

    -- Check if the user exists in the Administrator table
        SELECT 'admin'
        INTO o_role
        FROM Administrator
        WHERE Administrator.auth_id = user_id;
        
        IF FOUND THEN
            RETURN;
        END IF;
    
    RAISE EXCEPTION 'No role found for user with username %', i_username;
END;
$$
LANGUAGE plpgsql;
