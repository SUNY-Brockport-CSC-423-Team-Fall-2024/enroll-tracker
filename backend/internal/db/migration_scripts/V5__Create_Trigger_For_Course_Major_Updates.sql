CREATE OR REPLACE FUNCTION update_last_updated_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_updated = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp_course
BEFORE UPDATE ON Course
FOR EACH ROW
EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_timestamp_course
BEFORE UPDATE ON Major
FOR EACH ROW
EXECUTE FUNCTION update_last_updated_column();
