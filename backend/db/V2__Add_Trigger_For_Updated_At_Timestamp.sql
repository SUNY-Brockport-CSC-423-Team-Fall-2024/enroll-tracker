--Create triggers
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp_student
BEFORE UPDATE ON Student
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_timestamp_teacher
BEFORE UPDATE ON Teacher
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_timestamp_admin
BEFORE UPDATE ON Administrator
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
