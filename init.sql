--TABLES
---------
---------
CREATE TABLE IF NOT EXISTS universities (
    id          SERIAL  NOT NULL,
    name    VARCHAR(100) NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT universities_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS countries (
    id          SERIAL  NOT NULL,
    name        VARCHAR(100) NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT countries_pkey PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS users (
    id          UUID  NOT NULL,
    name        VARCHAR(100) NOT NULL,
    country_id  INTEGER NOT NULL,
    university_id INTEGER NOT NULL,
    username VARCHAR(100) NOT NULL,
    hash_token VARCHAR(100) NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT fk_countries FOREIGN KEY (country_id)
        REFERENCES countries(id),
    CONSTRAINT fk_universities FOREIGN KEY (university_id)
        REFERENCES universities(id)
);


--INSERTS
---------
---------
INSERT INTO universities(id, name, created_by, created_at, updated_by, updated_at) VALUES
    (1,'Universidad Nacional','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO countries(id,name,created_by, created_at, updated_by, updated_at) VALUES
    (1,'Colombia','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO users(id,name,country_id,university_id,username,hash_token,created_by, created_at, updated_by, updated_at) VALUES
    ('e647b1ad-2582-4d05-b3fb-98ebe4b27971'::UUID,'Juan',1,1,'juan123','anVhbjEyMzpwYXNzd29yZF90ZXN0','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW()),
    ('0c299523-0aeb-42d8-ac44-e2f1d5f04c05'::UUID,'Luis',1,1,'luis123','bHVpczEyMzpwYXNzd29yZF90ZXN0','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());


--PROCEDURES
------------
------------
CREATE OR REPLACE PROCEDURE p_add_user(j_input JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    r_user USERS%ROWTYPE;
BEGIN
    r_user.id := (j_input->>'id')::UUID;
    r_user.name := (j_input->>'name')::VARCHAR;
    r_user.country_id := (j_input->>'country_id')::INTEGER;
    r_user.university_id := (j_input->>'university_id')::INTEGER;
    r_user.username := (j_input->>'username')::VARCHAR;
    r_user.hash_token := (j_input->>'hash_token')::VARCHAR;
    r_user.created_by := (j_input->>'user_id_creator')::UUID;
    r_user.created_at := NOW();
    r_user.updated_by := (j_input->>'user_id_creator')::UUID;
    r_user.updated_at := NOW();

    INSERT INTO users VALUES(r_user.*);

EXCEPTION
    WHEN SQLSTATE '23503' THEN
        RAISE EXCEPTION 'Error [p_add_user: not found country or not found university  [%]:  % ]', SQLSTATE, SQLERRM USING ERRCODE = SQLSTATE;
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_add_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;


CREATE OR REPLACE PROCEDURE p_get_user(
    j_input JSONB,
    INOUT j_user JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    u_user_id UUID;
    r_user RECORD;
    exist BOOLEAN;
BEGIN
    u_user_id := (j_input->>'id')::UUID;

    SELECT EXISTS(SELECT 1 FROM users u WHERE u.id = u_user_id) INTO exist;

    IF exist IS FALSE THEN
        RAISE EXCEPTION SQLSTATE 'P0002';
    END IF;
    
    SELECT 
        u.id,
        u.name AS name,
        c.name AS country,
        un.name AS university      
    INTO r_user
    FROM users u
    JOIN countries c ON u.country_id = c.id
    JOIN universities un ON u.university_id = un.id
    WHERE u.id = u_user_id;

    j_user := JSONB_BUILD_OBJECT(
        'id', r_user.id,
        'name', r_user.name,
        'country',r_user.country,
        'university', r_user.university
    );

EXCEPTION
    WHEN SQLSTATE 'P0002' THEN
        RAISE EXCEPTION 'Error [p_get_user: not found user]' USING ERRCODE = SQLSTATE;
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_get_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM ;
END;
$BODY$;


CREATE OR REPLACE PROCEDURE p_edit_user(j_input JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    u_user_id UUID;
    u_user_id_creator UUID;
    exist BOOLEAN;
BEGIN
    u_user_id := (j_input->>'id')::UUID;
    u_user_id_creator := (j_input->>'user_id_creator')::UUID;

    SELECT EXISTS(SELECT 1 FROM users u WHERE u.id = u_user_id) INTO exist;

    IF exist IS FALSE THEN
        RAISE EXCEPTION SQLSTATE 'P0002';
    END IF;

    IF j_input->>'name' IS NOT NULL THEN
        UPDATE users 
            SET name = (j_input->>'name')::VARCHAR,
                updated_at = NOW(),
                updated_by = u_user_id_creator
        WHERE id = u_user_id;
    END IF;

    IF j_input->>'country_id' IS NOT NULL THEN
        UPDATE users 
            SET country_id = (j_input->>'country_id')::INTEGER,
                updated_at = NOW(),
                updated_by = u_user_id_creator
        WHERE id = u_user_id;
    END IF;

    IF j_input->>'university_id' IS NOT NULL THEN
        UPDATE users 
            SET university_id = (j_input->>'university_id')::INTEGER,
                updated_at = NOW(),
                updated_by = u_user_id_creator
        WHERE id = u_user_id;
    END IF;  

EXCEPTION
    WHEN SQLSTATE 'P0002' THEN
        RAISE EXCEPTION 'Error [p_edit_user: not found user]' USING ERRCODE = SQLSTATE;
    WHEN SQLSTATE '23503' THEN
        RAISE EXCEPTION 'Error [p_edit_user: not found country or not found university  [%]:  % ]', SQLSTATE, SQLERRM USING ERRCODE = SQLSTATE;
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_edit_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;


CREATE OR REPLACE PROCEDURE p_delete_user(j_input JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    u_user_id UUID;
    exist BOOLEAN;
BEGIN
    u_user_id := (j_input->>'id')::UUID;

    SELECT EXISTS(SELECT 1 FROM users u WHERE u.id = u_user_id) INTO exist;

    IF exist IS FALSE THEN
        RAISE EXCEPTION SQLSTATE 'P0002';
    END IF;

    DELETE FROM users WHERE users.id = u_user_id;

EXCEPTION
    WHEN SQLSTATE 'P0002' THEN
        RAISE EXCEPTION 'Error [p_delete_user: not found user]' USING ERRCODE = SQLSTATE;
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_delete_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;


CREATE OR REPLACE FUNCTION f_get_user_credentials(
    username_input VARCHAR(100)
) RETURNS JSONB
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE
    r_user RECORD;
    j_output JSONB;
    exist BOOLEAN;
BEGIN
    SELECT EXISTS(SELECT 1 FROM users u WHERE u.username = username_input) INTO exist;

    IF exist IS FALSE THEN
        RAISE EXCEPTION SQLSTATE 'P0002';
    END IF;

    SELECT 
        u.id,
        u.hash_token 
    INTO r_user
    FROM users u
    WHERE u.username = username_input;

    j_output := JSONB_BUILD_OBJECT(
        'id', r_user.id,
        'hash_token', r_user.hash_token
    );

    RETURN j_output;
    
EXCEPTION
    WHEN SQLSTATE 'P0002' THEN
        RAISE EXCEPTION 'Error [f_get_user_credentials: not found user]' USING ERRCODE = SQLSTATE;
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [f_get_user_credentials: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;