/** careful this is an overwrite operation, only uncomment the lines you actually need to run **/
/** tests will need to be updated if the seeds are updated so update carefully **/

\COPY teams TO '.db/seeds/CSV_teams.csv' CSV DELIMITER E'\t' QUOTE E'\b' ESCAPE '\' HEADER;

\COPY players TO '.db/seeds/CSV_players.csv' CSV DELIMITER E'\t' QUOTE E'\b' ESCAPE '\' HEADER;