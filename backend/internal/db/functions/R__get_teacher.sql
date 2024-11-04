CREATE OR REPLACE FUNCTION get_teacher (
    i_username varchar
)
RETURNS Teacher
AS $$
DECLARE
    result Teacher;
BEGIN
    SELECT *
    INTO result
    FROM Teacher
    INNER JOIN
        UserAuthentication ON Teacher.auth_id = UserAuthentication.id
    WHERE UserAuthentication.username = i_username
    AND UserAuthentication.is_active = true;
    
    RETURN result;

    EXCEPTION
        WHEN NO_DATA_FOUND THEN
            RAISE EXCEPTION 'No teacher found for username %s', i_username;
END;
$$
LANGUAGE plpgsql;
