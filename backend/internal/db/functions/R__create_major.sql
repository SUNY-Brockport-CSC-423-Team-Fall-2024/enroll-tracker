CREATE OR REPLACE FUNCTION create_major (
    i_name varchar,
    i_description varchar
)
RETURNS Major
AS $$
DECLARE
    major Major;
BEGIN
    BEGIN
        INSERT INTO Major (name, description, last_updated, created_at)
        VALUES
        (i_name, i_description, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
        RETURNING * INTO major;

    EXCEPTION
        WHEN integrity_constraint_violation THEN
            RAISE EXCEPTION 'Unable to create major';
    END;
    
    RETURN major;
    
END;
$$
LANGUAGE plpgsql;

