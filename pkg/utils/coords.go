package utils

import "strings"

// CoordsToGoto converts a G 3#3 format into g_3_3 format
func CoordsToGoto(input string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(input, " ", "_"), "#", "_"))
}

// GotoToCoords converts a g_3_3 format into G 3#3 format
func GotoToCoords(input string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(input, " ", "_"), "#", "_"))
}
