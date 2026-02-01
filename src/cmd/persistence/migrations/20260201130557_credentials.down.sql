SET statement_timeout = 0;

--bun:split

DROP TRIGGER IF EXISTS update_credentials_updated_at ON credentials;
DROP TABLE IF EXISTS credentials;

--bun:split

SELECT 2
