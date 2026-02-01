SET statement_timeout = 0;

--bun:split

CREATE TABLE IF NOT EXISTS permissions (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    alias VARCHAR(30) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY(id)
);

DROP TRIGGER IF EXISTS update_permissions_updated_at ON permissions;
CREATE TRIGGER update_permissions_updated_at
    BEFORE UPDATE ON permissions
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

--bun:split

SELECT 2
