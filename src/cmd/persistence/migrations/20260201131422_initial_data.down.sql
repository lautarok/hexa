SET statement_timeout = 0;

--bun:split

TRUNCATE TABLE credentials;
TRUNCATE TABLE users;
TRUNCATE TABLE permissions;
TRUNCATE TABLE role_permissions;
TRUNCATE TABLE roles;

--bun:split

SELECT 2
