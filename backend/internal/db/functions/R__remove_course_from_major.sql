CREATE OR REPLACE FUNCTION remove_course_from_major (
    i_major_id int,
    i_course_id int
)
RETURNS void
AS $$
DECLARE
    row_cnt int;
BEGIN
    DELETE FROM Course_Major
    WHERE
        major_id = i_major_id
        AND course_id = i_course_id;
    
    GET DIAGNOSTICS row_cnt = ROW_COUNT;
    IF row_cnt = 0 THEN
        RAISE EXCEPTION 'No matching record found for major ID % and course ID %', i_major_id, i_course_id;
    END IF;

    EXCEPTION
        WHEN foreign_key_violation THEN
            RAISE EXCEPTION 'Cannot delete due to foreign key constraint violation for major ID % and course ID %', i_major_id, i_course_id;
        WHEN data_exception THEN
            RAISE EXCEPTION 'Invalid data input for major ID % or course ID %', i_major_id, i_course_id;
END;
$$
LANGUAGE plpgsql;
