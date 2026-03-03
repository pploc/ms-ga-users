CREATE TABLE user_profiles (
    user_id                   UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    department                VARCHAR(100),
    hire_date                 DATE,
    emergency_contact_name    VARCHAR(200),
    emergency_contact_phone   VARCHAR(20),
    notes                     TEXT,
    updated_at                TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
