CREATE OR REPLACE FUNCTION get_major (
    i_id int
)
RETURNS Major
AS $$
DECLARE
    major Major;
BEGIN
    SELECT *
    INTO major
    FROM Major
    WHERE id = i_id;

    RETURN major;
END;
$$
LANGUAGE plpgsql;
