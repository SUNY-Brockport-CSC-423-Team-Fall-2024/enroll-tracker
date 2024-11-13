CREATE OR REPLACE FUNCTION unenroll_student (
    i_course_id int,
    i_student_id int
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    UPDATE Enrollments
    SET
        is_enrolled = false,
        unenrolled_date = CURRENT_TIMESTAMP
    WHERE
        course_id = i_course_id
        AND student_id = i_student_id;
    
    GET DIAGNOSTICS row_cnt = ROW_COUNT;
    
    RETURN row_cnt > 0;
    
    EXCEPTION
        WHEN foreign_key_violation THEN
            RAISE EXCEPTION 'Foreign key violation: student % or course % does not exist.', i_student_id, i_course_id;
        WHEN not_null_violation THEN
            RAISE EXCEPTION 'Null value violation: all fields must be non-null.';
        WHEN check_violation THEN
            RAISE EXCEPTION 'Check constraint violation for student % or course %.', i_student_id, i_course_id;
        WHEN data_exception THEN
            RAISE EXCEPTION 'Data exception: invalid data type or format for student % or course %.', i_student_id, i_course_id;
END;
$$
LANGUAGE plpgsql;

