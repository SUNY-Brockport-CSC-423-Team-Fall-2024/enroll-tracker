CREATE OR REPLACE FUNCTION add_student_to_major (
    i_major_id int,
    i_student_id int
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
    cur_student_major int;
BEGIN
    SELECT S.major_id
    INTO cur_student_major
    FROM Student AS S
    WHERE S.id = i_student_id;

    IF NOT cur_student_major IS NULL THEN
        RAISE EXCEPTION 'Student %s is already apart of a major %', i_student_id, cur_student_major;
    END IF;

    UPDATE Student
    SET
        major_id = i_major_id
    WHERE
        id = i_student_id;

    GET DIAGNOSTICS row_cnt = ROW_COUNT;

    RETURN row_cnt > 0;
    
END;
$$
LANGUAGE plpgsql;
