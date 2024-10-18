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
    BEGIN
        SELECT 1
        INTO o_role
        FROM Student
        WHERE Student.auth_id = user_id;
        o_role := 'student';
        RETURN;
    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            -- Continue to the next role check
            NULL;
    END;

    -- Check if the user exists in the Teacher table
    BEGIN
        SELECT 1
        INTO o_role
        FROM Teacher
        WHERE Teacher.auth_id = user_id;
        
        o_role := 'teacher';
        RETURN;
    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            -- Continue to the next role check
            NULL;
    END;

    -- Check if the user exists in the Administrator table
    BEGIN
        SELECT 1
        INTO o_role
        FROM Administrator
        WHERE Administrator.auth_id = user_id;
        
        o_role := 'admin';
        RETURN;
    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            -- If no role is found, raise an exception
            RAISE EXCEPTION 'No role found for user with username %s', i_username;
    END;

END;
$$
LANGUAGE plpgsql;
