CREATE OR REPLACE FUNCTION delete_course_from_majors (
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
            DELETE FROM Course_Major
            WHERE
                major_id = i_major_id
                AND course_id = i_course_id;
            
            EXCEPTION
                WHEN foreign_key_violation THEN
                    RAISE EXCEPTION 'Cannot delete due to foreign key constraint violation for major ID % and course ID %', i_major_id, i_course_id;
                WHEN data_exception THEN
                    RAISE EXCEPTION 'Invalid data input for major ID % or course ID %', i_major_id, i_course_id;
        END;
    END LOOP;
END;        
$$
LANGUAGE plpgsql;
