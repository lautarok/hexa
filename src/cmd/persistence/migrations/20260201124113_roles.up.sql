SET statement_timeout = 0;

--bun:split

CREATE TABLE IF NOT EXISTS roles (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name_en VARCHAR(30) UNIQUE,
    name_es VARCHAR(30) UNIQUE,
    name_fr VARCHAR(30) UNIQUE,
    name_pt VARCHAR(30) UNIQUE,
    name_nl VARCHAR(30) UNIQUE,
    lock BOOLEAN NOT NULL DEFAULT FALSE,
    slug VARCHAR(20) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id)
);

DROP TRIGGER IF EXISTS update_roles_updated_at ON roles;
CREATE TRIGGER update_roles_updated_at
    BEFORE UPDATE ON roles
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

--bun:split

SELECT 2
