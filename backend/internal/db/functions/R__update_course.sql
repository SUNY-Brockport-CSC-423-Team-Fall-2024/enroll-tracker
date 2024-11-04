CREATE OR REPLACE FUNCTION update_course (
    i_course_id int,
    i_description varchar,
    i_teacher_id int,
    i_max_enrollment int,
    i_num_credits int,
    i_status CourseStatus
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    UPDATE Course
    SET
       description = COALESCE(i_description, description),
       teacher_id = COALESCE(i_teacher_id, teacher_id),
       max_enrollment = COALESCE(i_max_enrollment, max_enrollment),
       num_credits = COALESCE(i_num_credits, num_credits),
       status = COALESCE(i_status, status)
    WHERE
        id = i_course_id;
    
    GET DIAGNOSTICS row_cnt = ROW_COUNT;

    RETURN row_cnt > 0;
END;
$$
LANGUAGE plpgsql;
