TRUNCATE teams CASCADE;

TRUNCATE players CASCADE;

\COPY teams FROM '.db/seeds/CSV_teams.csv' CSV DELIMITER E'\t' QUOTE E'\b' ESCAPE '\' HEADER;

\COPY players FROM '.db/seeds/CSV_players.csv' CSV DELIMITER E'\t' QUOTE E'\b' ESCAPE '\' HEADER;