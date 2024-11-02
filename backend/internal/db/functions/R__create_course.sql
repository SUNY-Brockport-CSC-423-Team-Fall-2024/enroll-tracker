CREATE OR REPLACE FUNCTION create_course (
    i_name varchar,
    i_description varchar,
    i_teacher_id int,
    i_max_enrollment int,
    i_num_credits int
)
RETURNS Course
AS $$
DECLARE
    course Course;
BEGIN
    BEGIN
        INSERT INTO Course (name, description, teacher_id, max_enrollment, num_credits, last_updated, created_at)
        VALUES
        (i_name, i_description, i_teacher_id, i_max_enrollment, i_num_credits, CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)
        RETURNING * INTO course;

    EXCEPTION
        WHEN integrity_constraint_violation THEN
            RAISE EXCEPTION 'Unable to create course';
    END;
    
    RETURN course;
    
END;
$$
LANGUAGE plpgsql;

