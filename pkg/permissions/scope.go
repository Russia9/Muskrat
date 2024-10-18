package permissions

import (
	"errors"
)

type Scope struct {
	ID         int64
	PlayerRole PlayerRole

	SquadID   string
	GuildID   string
	SquadRole SquadRole
}

type PlayerRole int

const (
	PlayerRoleBanned PlayerRole = iota - 1
	PlayerRoleUnregistered
	PlayerRoleUser
)

type SquadRole int

const (
	SquadRoleStranger SquadRole = iota
	SquadRoleMember
	SquadRoleSquire
	SquadRoleLeader
)

var ErrForbidden = errors.New("forbidden")
