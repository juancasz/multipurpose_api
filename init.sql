CREATE TABLE IF NOT EXISTS auth_users (
    id          UUID  NOT NULL,
    username    VARCHAR NOT NULL,
    hash_password   VARCHAR NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT auth_user_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS countries (
    id          SERIAL  NOT NULL,
    name        VARCHAR NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT countries_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id          UUID  NOT NULL,
    name        VARCHAR NOT NULL,
    country_id  INTEGER NOT NULL,
    auth_user_id UUID NOT NULL,
    created_by  UUID NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_by  UUID NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT fk_countries FOREIGN KEY (country_id)
        REFERENCES countries(id),
    CONSTRAINT fk_auth_users FOREIGN KEY (auth_user_id)
        REFERENCES auth_users(id)
);

INSERT INTO auth_users(id,username,hash_password,created_by, created_at, updated_by, updated_at) VALUES
    ('22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,'juan123','test-hash-pashword','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO countries(id,name,created_by, created_at, updated_by, updated_at) VALUES
    (1,'Colombia','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO users(id,name,country_id,auth_user_id,created_by, created_at, updated_by, updated_at) VALUES
    ('e647b1ad-2582-4d05-b3fb-98ebe4b27971'::UUID,'Juan',1,'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

