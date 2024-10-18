package domain

type Castle string

const (
	CastleRed     Castle = "red"
	CastleYellow  Castle = "yellow"
	CastleGreen   Castle = "green"
	CastleBlue    Castle = "blue"
	CastleUnknown Castle = "unknown"
)

func FlagToCastle(flag string) Castle {
	switch flag {
	case "🇮🇲":
		return CastleRed
	case "🇻🇦":
		return CastleYellow
	case "🇲🇴":
		return CastleGreen
	case "🇪🇺":
		return CastleBlue
	}
	return CastleUnknown
}
