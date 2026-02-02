SET statement_timeout = 0;

--bun:split

CREATE TABLE IF NOT EXISTS credentials (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    email VARCHAR(200) UNIQUE NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    password VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

DROP TRIGGER IF EXISTS update_credentials_updated_at ON credentials;
CREATE TRIGGER update_credentials_updated_at
    BEFORE UPDATE ON credentials
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();

--bun:split

SELECT 2
