CREATE TABLE IF NOT EXISTS players
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    team_id uuid NULL,
    name text,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp without time zone,
    CONSTRAINT pk_players PRIMARY KEY(id),
    CONSTRAINT fk_players_team FOREIGN KEY(team_id) REFERENCES teams(id)
);