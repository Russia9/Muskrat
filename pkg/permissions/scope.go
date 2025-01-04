package permissions

import (
	"errors"
)

type Scope struct {
	ID         int64
	PlayerRole PlayerRole

	SquadID   *string
	GuildID   *string
	SquadRole SquadRole
}

type PlayerRole int

const (
	PlayerRoleBanned PlayerRole = iota - 1
	PlayerRoleUnregistered
	PlayerRoleUser
)

type SquadRole string

const (
	SquadRoleNone   SquadRole = ""
	SquadRoleMember SquadRole = "member"

	SquadRoleBartender SquadRole = "bartender"
	SquadRoleSquire    SquadRole = "squire"
	SquadRoleCommander SquadRole = "commander"

	SquadRoleLeader SquadRole = "leader"
)

var ErrForbidden = errors.New("forbidden")
