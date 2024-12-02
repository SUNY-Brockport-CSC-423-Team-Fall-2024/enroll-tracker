CREATE OR REPLACE FUNCTION add_course_to_major (
    i_major_ids int[],
    i_course_id int
)
RETURNS void
AS $$
DECLARE
    i_major_id int;
BEGIN
    FOREACH i_major_id IN ARRAY i_major_ids LOOP
        BEGIN
            INSERT INTO Course_Major (major_id, course_id)
            VALUES
            (i_major_id, i_course_id);
            EXCEPTION 
            WHEN unique_violation THEN
                RAISE EXCEPTION 'Association between major % and course % already exists.', i_major_id, i_course_id;
            WHEN foreign_key_violation THEN
                RAISE EXCEPTION 'Foreign key violation: major % or course % does not exist.', i_major_id, i_course_id;
            WHEN not_null_violation THEN
                RAISE EXCEPTION 'Null value violation: all fields must be non-null.';
            WHEN check_violation THEN
                RAISE EXCEPTION 'Check constraint violation for major % or course %.', i_major_id, i_course_id;
            WHEN data_exception THEN
            RAISE EXCEPTION 'Data exception: invalid data type or format for major % or course %.', i_major_id, i_course_id;
        END;
    END LOOP;
END;
$$
LANGUAGE plpgsql;
