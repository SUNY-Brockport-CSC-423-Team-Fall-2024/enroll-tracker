CREATE OR REPLACE FUNCTION update_major (
    i_major_id int,
    i_description varchar,
    i_status MajorStatus
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    UPDATE Major
    SET
       description = COALESCE(i_description, description),
       status = COALESCE(i_status, status)
    WHERE
        id = i_major_id;
    
    GET DIAGNOSTICS row_cnt = ROW_COUNT;

    RETURN row_cnt > 0;
END;
$$
LANGUAGE plpgsql;
