CREATE OR REPLACE FUNCTION delete_course (
   i_course_id int
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    Update Course
    SET
        max_enrollment = 0,
        status = 'inactive'
    WHERE
        id = i_course_id;

    GET DIAGNOSTICS row_cnt = ROW_COUNT;    

    RETURN row_cnt > 0;
END;
$$
LANGUAGE plpgsql;
