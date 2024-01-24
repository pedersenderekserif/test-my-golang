CREATE TABLE IF NOT EXISTS teams
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    name text,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    CONSTRAINT pk_teams PRIMARY KEY(id)
);