CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL, -- @sensitive
    role VARCHAR(20) NOT NULL DEFAULT 'operator' CHECK (role IN ('operator', 'viewer')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
