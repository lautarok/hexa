SET statement_timeout = 0;

--bun:split

DROP TRIGGER IF EXISTS update_permissions_updated_at ON permissions;
DROP TABLE IF EXISTS permissions;

--bun:split

SELECT 2
