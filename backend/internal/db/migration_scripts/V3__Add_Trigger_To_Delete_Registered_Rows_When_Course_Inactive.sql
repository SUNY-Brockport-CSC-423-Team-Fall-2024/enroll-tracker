CREATE OR REPLACE FUNCTION deleted_registered_if_inactive()
RETURNS TRIGGER AS $$
BEGIN
    --Check if course is being set to inactive
    IF NEW.status = 'inactive' THEN
        DELETE FROM Registered WHERE course_id = NEW.id;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER course_status_update
BEFORE UPDATE ON Course
FOR EACH ROW
WHEN (OLD.status IS DISTINCT FROM NEW.status)
EXECUTE FUNCTION deleted_registered_if_inactive();
