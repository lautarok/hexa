SET statement_timeout = 0;

--bun:split

DROP TRIGGER IF EXISTS update_role_permissions_updated_at ON role_permissions;
DROP TABLE IF EXISTS role_permissions;

--bun:split

SELECT 2
