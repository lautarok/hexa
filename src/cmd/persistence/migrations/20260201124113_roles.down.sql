SET statement_timeout = 0;

--bun:split

DROP TRIGGER IF EXISTS update_roles_updated_at ON roles;
DROP TABLE IF EXISTS roles;

--bun:split

SELECT 2
