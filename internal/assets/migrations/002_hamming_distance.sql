-- +migrate Up

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION hamming_distance(val1 BYTEA, val2 BYTEA) RETURNS INTEGER AS $$
DECLARE
    distance INTEGER := 0;
    i INTEGER;
BEGIN
    IF LENGTH(val1) <> LENGTH(val2) THEN
        RAISE EXCEPTION 'Inputs must have the same length for Hamming distance calculation';
    END IF;

    FOR i IN 1..LENGTH(val1) LOOP
            distance := distance + (GET_BYTE(val1, i-1) # GET_BYTE(val2, i-1));
    END LOOP;

    RETURN distance;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down

DROP FUNCTION IF EXISTS hamming_distance;