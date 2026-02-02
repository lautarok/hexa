SET statement_timeout = 0;

--bun:split

DELETE FROM credentials;
DELETE FROM users;
DELETE FROM role_permissions;
DELETE FROM roles;
DELETE FROM permissions;

--bun:split

SELECT 2
