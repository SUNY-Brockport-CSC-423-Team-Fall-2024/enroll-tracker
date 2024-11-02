CREATE OR REPLACE FUNCTION delete_major (
   i_major_id int
)
RETURNS BOOLEAN
AS $$
DECLARE
    row_cnt int;
BEGIN
    Update Major
    SET
        status = 'inactive'
    WHERE
        id = i_major_id;

    GET DIAGNOSTICS row_cnt = ROW_COUNT;    

    RETURN row_cnt > 0;
END;
$$
LANGUAGE plpgsql;
