CREATE OR REPLACE FUNCTION create_user_session(
    i_user_id   int,
    i_sess_id   text,
    i_iat timestamp,
    i_expiration_time timestamp,
    i_not_before timestamp,
    OUT o_id int,
    OUT o_user_id int,
    OUT o_sess_id text,
    OUT o_issued_at timestamp,
    OUT o_expiration_time timestamp,
    OUT o_not_before timestamp
)
AS $$
BEGIN
    INSERT INTO UserSession (user_id, sess_id, issued_at, expiration_time, not_before)
    VALUES (i_user_id, i_sess_id, i_iat, i_expiration_time, i_not_before)
    RETURNING *
    INTO
        o_id,
        o_user_id,
        o_sess_id,
        o_issued_at,
        o_expiration_time,
        o_not_before;
END;
$$
LANGUAGE plpgsql;
