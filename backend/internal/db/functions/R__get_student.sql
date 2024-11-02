CREATE OR REPLACE FUNCTION get_student (
    i_username varchar
)
RETURNS Student
AS $$
DECLARE
    result Student;
BEGIN
    SELECT Student.*
    INTO result
    FROM Student
    INNER JOIN 
        UserAuthentication ON Student.auth_id = UserAuthentication.id
    WHERE UserAuthentication.username = i_username
    AND UserAuthentication.is_active = true;
    
    RETURN result;

    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            RAISE EXCEPTION 'No student found for username %s', i_username;
END;
$$
LANGUAGE plpgsql;
