CREATE OR REPLACE FUNCTION get_course (
    i_id int
)
RETURNS Course
AS $$
DECLARE
    course Course;
BEGIN
    SELECT *
    INTO course
    FROM Course
    WHERE id = i_id;

    RETURN course;
END;
$$
LANGUAGE plpgsql;
