CREATE TYPE castle AS ENUM (
    'red',
    'yellow',
    'green',
    'blue',
    'unknown'
    );

CREATE TYPE squad_role AS ENUM (
    'stranger',
    'member',

    'bartender',
    'squire',
    'commander',

    'leader'
    );

CREATE TABLE squads (
    id         UUID PRIMARY KEY,
    chat_id    UUID         NOT NULL,
    name       VARCHAR(255) NOT NULL,
    castle     castle       NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE guilds (
    id         UUID PRIMARY KEY,
    squad_id   UUID         REFERENCES squads (id) ON DELETE SET NULL,
    name       VARCHAR(255) NOT NULL,
    tag        VARCHAR(16)  NOT NULL,
    leader_id  BIGINT       NOT NULL,
    level      INTEGER      NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE players (
    id                 BIGINT PRIMARY KEY,
    username           VARCHAR(32)  NOT NULL,
    player_role        INTEGER      NOT NULL,

    language           VARCHAR(32)  NOT NULL,

    squad_id           UUID         REFERENCES squads (id) ON DELETE SET NULL,
    guild_id           UUID         REFERENCES guilds (id) ON DELETE SET NULL,
    squad_role         squad_role   NOT NULL DEFAULT 'stranger',

    first_seen         TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    last_seen          TIMESTAMPTZ  NOT NULL DEFAULT NOW(),

    castle             castle       NOT NULL DEFAULT 'unknown',
    player_name        VARCHAR(255) NOT NULL DEFAULT '',

    level              INTEGER      NOT NULL DEFAULT 0,
    current_exp        INTEGER      NOT NULL DEFAULT 0,
    next_level_exp     INTEGER      NOT NULL DEFAULT 0,

    rank               INTEGER      NOT NULL DEFAULT 0,
    str                INTEGER      NOT NULL DEFAULT 0,
    dex                INTEGER      NOT NULL DEFAULT 0,
    vit                INTEGER      NOT NULL DEFAULT 0,
    detailed_stats     JSONB,
    profile_updated_at TIMESTAMPTZ  NOT NULL DEFAULT 'epoch',

    schools            JSONB,
    schools_updated_at TIMESTAMPTZ  NOT NULL DEFAULT 'epoch',

    player_balance     INTEGER      NOT NULL DEFAULT 0,
    bank_balance       INTEGER      NOT NULL DEFAULT 0,
    balance_updated_at TIMESTAMPTZ  NOT NULL DEFAULT 'epoch'
);
