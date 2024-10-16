CREATE OR REPLACE FUNCTION get_user_session(
    i_sess_id text,
    OUT o_id int,
    OUT o_user_id int,
    OUT o_sess_id text,
    OUT o_issued_at timestamp,
    OUT o_expiration_time timestamp,
    OUT o_not_before timestamp
)
AS $$
BEGIN
    SELECT *
    INTO
        o_id,
        o_user_id,
        o_sess_id,
        o_issued_at,
        o_expiration_time,
        o_not_before
    FROM UserSession
    WHERE UserSession.sess_id = i_sess_id;
    
    IF NOT FOUND THEN
        RAISE EXCEPTION 'No session found for session id %s', get_user_session.i_sess_id;
    END IF;
END;
$$
LANGUAGE plpgsql;
