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
    hash_password VARCHAR(100) NOT NULL,
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

INSERT INTO users(id,name,country_id,university_id,username,hash_password,created_by, created_at, updated_by, updated_at) VALUES
    ('e647b1ad-2582-4d05-b3fb-98ebe4b27971'::UUID,'Juan',1,1,'juan123','password_test','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

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
    r_user.hash_password := (j_input->>'hash_password')::VARCHAR;
    r_user.created_by := (j_input->>'user_id_creator')::UUID;
    r_user.created_at := NOW();
    r_user.updated_by := (j_input->>'user_id_creator')::UUID;
    r_user.updated_at := NOW();

    INSERT INTO users VALUES(r_user.*);

EXCEPTION
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
BEGIN
    u_user_id := (j_input->>'id')::UUID;
    
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
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_get_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;


CREATE OR REPLACE PROCEDURE p_edit_user(j_input JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    u_user_id UUID;
BEGIN
    u_user_id := (j_input->>'id')::UUID;

    IF j_input->>'name' IS NOT NULL THEN
        UPDATE users SET name = (j_input->>'name')::VARCHAR 
        WHERE id = u_user_id;
    END IF;

    IF j_input->>'country_id' IS NOT NULL THEN
        UPDATE users SET country_id = (j_input->>'country_id')::INTEGER 
        WHERE id = u_user_id;
    END IF;

    IF j_input->>'university_id' IS NOT NULL THEN
        UPDATE users SET university_id = (j_input->>'university_id')::INTEGER 
        WHERE id = u_user_id;
    END IF;  

EXCEPTION
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_edit_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;

CREATE OR REPLACE PROCEDURE p_delete_user(j_input JSONB)
LANGUAGE 'plpgsql'
AS $BODY$
DECLARE 
    u_user_id UUID;
BEGIN
    u_user_id := (j_input->>'id')::UUID;

    DELETE FROM users WHERE users.id = u_user_id;

EXCEPTION
    WHEN OTHERS THEN
        RAISE EXCEPTION 'Error [p_delete_user: uncontrolled error  [%]:  % ]', SQLSTATE, SQLERRM;
END;
$BODY$;