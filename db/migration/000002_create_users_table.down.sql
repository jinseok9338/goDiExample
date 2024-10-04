-- Drop the trigger first
DROP TRIGGER IF EXISTS update_users_updated_at ON members_user;

-- Drop the function next
DROP FUNCTION IF EXISTS update_timestamp();

-- Drop the users table
DROP TABLE IF EXISTS members_user;
