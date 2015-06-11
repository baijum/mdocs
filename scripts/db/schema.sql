DROP TABLE IF EXISTS "member" CASCADE;

CREATE TABLE "member" (
       id serial PRIMARY KEY,
       username TEXT UNIQUE NOT NULL,
       email TEXT UNIQUE NOT NULL,
       password BYTEA NOT NULL,
       salt BYTEA NOT NULL,
       fullname TEXT NOT NULL,
       verified BOOLEAN DEFAULT false,
       active BOOLEAN DEFAULT false,
       created_at TIMESTAMP DEFAULT now(),
       update_at TIMESTAMP DEFAULT now()
);

DROP TABLE IF EXISTS "doc" CASCADE;

CREATE TABLE "doc" (
       id serial PRIMARY KEY,
       link TEXT NOT NULL,
       source TEXT NOT NULL,
       created_at TIMESTAMP DEFAULT now(),
       update_at TIMESTAMP DEFAULT now(),
       member_id BIGINT REFERENCES member (id)
);
