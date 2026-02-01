SET statement_timeout = 0;

--bun:split

CREATE TABLE IF NOT EXISTS role_permissions (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),

    CONSTRAINT uk_role_permission UNIQUE(role_id, permission_id),
    
    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
        REFERENCES roles(id) ON DELETE CASCADE,

    CONSTRAINT fk_permission
        FOREIGN KEY (permission_id)
        REFERENCES permissions(id) ON DELETE CASCADE
);

DROP TRIGGER IF EXISTS update_role_permissions_updated_at ON role_permissions;
CREATE TRIGGER update_role_permissions_updated_at
    BEFORE UPDATE ON role_permissions
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

--bun:split

SELECT 2
