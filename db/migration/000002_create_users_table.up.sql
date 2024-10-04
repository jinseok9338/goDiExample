CREATE TABLE IF NOT EXISTS members_user (
    id SERIAL PRIMARY KEY,            -- Auto-incrementing primary key
    username VARCHAR(50) UNIQUE NOT NULL,  -- Username with a unique constraint
    email VARCHAR(100) UNIQUE NOT NULL,    -- Email with a unique constraint
    password VARCHAR(255) NOT NULL,   -- Password (hashed)
    first_name VARCHAR(50),           -- First name of the user
    last_name VARCHAR(50),            -- Last name of the user
    is_active BOOLEAN DEFAULT TRUE,   -- Boolean flag for active/inactive status
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Update timestamp
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON members_user
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();