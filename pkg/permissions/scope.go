package permissions

import (
	"errors"
)

type Scope struct {
	ID         int64
	PlayerRole PlayerRole

	SquadID   *string
	SquadRole SquadRole

	GuildID   *string
	GuildRole SquadRole
}

type PlayerRole int

const (
	PlayerRoleBanned PlayerRole = iota - 1
	PlayerRoleUnregistered
	PlayerRoleUser
	PlayerRoleInternal
)

type SquadRole int

const (
	SquadRoleNone SquadRole = iota
	SquadRoleMember
	SquadRoleSquire
	SquadRoleLeader
)

var ErrForbidden = errors.New("forbidden")

var InternalScope = Scope{PlayerRole: PlayerRoleInternal}
