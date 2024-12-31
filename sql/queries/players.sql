-- name: CreatePlayer :exec
INSERT
INTO players (id, username, player_role, language, first_seen, last_seen)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: DeletePlayer :exec
DELETE
FROM players
WHERE id = $1;

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

-- name: UpdatePlayer :exec
UPDATE players
SET username           = $2,
    player_role        = $3,
    language           = $4,
    squad_id           = $5,
    guild_id           = $6,
    squad_role         = $7,
    first_seen         = $8,
    last_seen          = $9,
    castle             = $10,
    player_name        = $11,
    level              = $12,
    current_exp        = $13,
    next_level_exp     = $14,
    rank               = $15,
    str                = $16,
    dex                = $17,
    vit                = $18,
    detailed_stats     = $19,
    profile_updated_at = $20,
    schools            = $21,
    schools_updated_at = $22,
    player_balance     = $23,
    bank_balance       = $24,
    balance_updated_at = $25
WHERE id = $1;
