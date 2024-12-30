-- name: CreatePlayer :exec
INSERT INTO players (id, username, player_role, language, first_seen, last_seen)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetPlayer :one
SELECT *
FROM players
WHERE id = $1
LIMIT 1;

-- name: GetPlayerByUsername :one
SELECT *
FROM players
WHERE username = $1
LIMIT 1;
