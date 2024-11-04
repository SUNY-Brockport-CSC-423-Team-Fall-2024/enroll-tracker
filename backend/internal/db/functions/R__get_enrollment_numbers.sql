CREATE OR REPLACE FUNCTION get_enrollment_numbers (
    i_course_id int,
    i_student_id int
)
RETURNS TABLE(
    enrolled_count int,
    max_enrollment int,
    students_enrolled_credits int,
    course_credits int
)
AS $$
DECLARE
    enrolled_count int;
    max_enrollment int;
    students_enrolled_credits int;
    course_credits int;
BEGIN
    SELECT COUNT(*)
    INTO enrolled_count
    FROM Enrollments AS E
    WHERE
        E.course_id = i_course_id
        AND E.is_enrolled = true;
    
    SELECT C.max_enrollment
    INTO max_enrollment
    FROM Course AS C
    WHERE C.id = i_course_id;

    SELECT COALESCE(SUM(C.num_credits), 0)
    INTO students_enrolled_credits
    FROM Enrollments AS E
    JOIN Course AS C ON C.id = E.course_id
    WHERE E.student_id = i_student_id
    AND E.is_enrolled = true;

    SELECT C.num_credits
    INTO course_credits
    FROM Course AS C
    WHERE C.id = i_course_id;
    
    RETURN QUERY
    SELECT enrolled_count, max_enrollment, students_enrolled_credits, course_credits;
END;
$$
LANGUAGE plpgsql;
