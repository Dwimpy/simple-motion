CREATE TABLE users (
       id BIGSERIAL PRIMARY KEY,
       email VARCHAR(256) UNIQUE NOT NULL,
       hashed_password VARCHAR(256) NOT NULL,
       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);