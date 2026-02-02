SET statement_timeout = 0;

--bun:split

CREATE TABLE IF NOT EXISTS google_identities (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    google_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id, google_id),

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

--bun:split

SELECT 2
