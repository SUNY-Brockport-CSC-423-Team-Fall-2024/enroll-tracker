CREATE OR REPLACE FUNCTION get_user_id (
    i_username varchar,
    OUT o_id varchar
)
AS $$
DECLARE
    a_id int;
BEGIN
    -- Get the user_id from the UserAuthentication table
    SELECT id 
    INTO a_id
    FROM UserAuthentication
    WHERE username = i_username;

    -- Check if the user exists in the Student table
        SELECT Student.id
        INTO o_id
        FROM Student
        WHERE Student.auth_id = a_id;
        
        IF FOUND THEN
            RETURN;
        END IF;

    -- Check if the user exists in the Teacher table
        SELECT Teacher.id
        INTO o_id
        FROM Teacher
        WHERE Teacher.auth_id = a_id;
        
        IF FOUND THEN
            RETURN;
        END IF;

    -- Check if the user exists in the Administrator table
        SELECT Administrator.id
        INTO o_id
        FROM Administrator
        WHERE Administrator.auth_id = a_id;
        
        IF FOUND THEN
            RETURN;
        END IF;
    
    RAISE EXCEPTION 'No id found for user with username %', i_username;
END;
$$
LANGUAGE plpgsql;
