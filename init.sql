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

INSERT INTO universities(id, name, created_by, created_at, updated_by, updated_at) VALUES
    (1,'Universidad Nacional','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO countries(id,name,created_by, created_at, updated_by, updated_at) VALUES
    (1,'Colombia','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

INSERT INTO users(id,name,country_id,university_id,username,hash_password,created_by, created_at, updated_by, updated_at) VALUES
    ('e647b1ad-2582-4d05-b3fb-98ebe4b27971'::UUID,'Juan',1,1,'juan123','password_test','22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW(),'22cbcc24-f6c4-47d7-8374-76593391c2b2'::UUID,NOW());

