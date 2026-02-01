SET statement_timeout = 0;

--bun:split

WITH

inserted_admin_permission AS (
    INSERT INTO permissions (alias) VALUES ('admin')
    RETURNING id
),

inserted_common_user_permission AS (
    INSERT INTO permissions (alias) VALUES ('user:common')
    RETURNING id
),

inserted_admin_role AS (
    INSERT INTO roles (name_en, name_es, name_fr, name_pt, name_nl, slug, lock)
    VALUES (
        'Administrator',
        'Administrador',
        'Administrateur',
        'Administrador',
        'Beheerder',
        'admin',
        TRUE
    )
    RETURNING id
),

inserted_common_user_role AS (
    INSERT INTO roles (name_en, name_es, name_fr, name_pt, name_nl, slug, lock)
    VALUES (
        'User',
        'Usuario',
        'Utilisateur',
        'Usuário',
        'Gebruiker',
        'common:user',
        TRUE
    )
    RETURNING id
),

inserted_admin_role_permissions AS (
    INSERT INTO role_permissions (permission_id, role_id)
    SELECT p.id, r.id
    FROM inserted_admin_permission p, inserted_admin_role r
    RETURNING id
),

inserted_common_user_role_permissions AS (
    INSERT INTO role_permissions (permission_id, role_id)
    SELECT p.id, r.id
    FROM inserted_common_user_permission p, inserted_common_user_role r
    RETURNING id
),

inserted_admin_user AS (
    INSERT INTO users (name, surname, role_id)
    SELECT
        'Augusto Lautaro',
        'Kazalukian',
        r.id
    FROM inserted_admin_role r
    RETURNING id
),

inserted_common_user AS (
    INSERT INTO users (name, surname, role_id)
    SELECT
        'Ana Beatriz',
        'Schultheis',
        r.id
    FROM inserted_common_user_role r
    RETURNING id
),

inserted_admin_credential AS (
    INSERT INTO credentials (
        email,
        username,
        password,
        user_id
    )
    SELECT
        'lautarok2004@gmail.com',
        'lautaro',
        '$2y$10$fRqMwqGciCVEUCq.z7hi1e9TEsIDRTckabaXxvpn2xctT8JwNGipW', --$Admin$2004
        u.id
    FROM inserted_admin_user u
    RETURNING id
)

INSERT INTO credentials (
    email,
    username,
    password,
    user_id
)
SELECT
    'anabeatrizschultheis@gmail.com',
    'beatriz',
    '$2y$10$KocvWS1DekSQqqvsbHbtaeNV.naDfgOtytmDcLRtlzZeauGtEaPbS', --$User$1957
    u.id
FROM inserted_common_user u;

--bun:split

SELECT 2
