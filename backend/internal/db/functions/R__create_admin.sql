CREATE OR REPLACE FUNCTION create_administrator (
    i_first_name varchar,
    i_last_name varchar,
    i_auth_id int,
    i_phone_number varchar,
    i_email varchar,
    i_office varchar,
    OUT o_id int,
    OUT o_first_name varchar,
    OUT o_last_name varchar,
    OUT o_auth_id int,
    OUT o_phone_number varchar,
    OUT o_email varchar,
    OUT o_office varchar,
    OUT o_created_at timestamp,
    OUT o_updated_at timestamp
)
AS $$
BEGIN
    INSERT INTO Administrator (first_name, last_name, auth_id, phone_number, email, office)
    VALUES (i_first_name, i_last_name, i_auth_id, i_phone_number, i_email, i_office)
    RETURNING 
        id,
        first_name,
        last_name,
        auth_id,
        phone_number,
        email,
        office,
        created_at,
        updated_at
    INTO
        o_id,
        o_first_name,
        o_last_name,
        o_auth_id,
        o_phone_number,
        o_email,
        o_office,
        o_created_at,
        o_updated_at;
END;
$$
LANGUAGE plpgsql;
