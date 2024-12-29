package utils

// Convert digit to keycap emoji (1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 0️⃣)
func DigitToKeycap(x int) string {
	if x < 0 || x > 9 {
		return ""
	}

	return string(rune(0x30+x)) + string(rune(0xFE0F)) + string(rune(0x20E3))
}

// Convert keycap emoji (1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 0️⃣) to digit
func KeycapToDigit(x string) int {
	if len([]rune(x)) != 3 || []rune(x)[1] != 0xFE0F || []rune(x)[2] != 0x20E3 {
		return -1
	}

	return int(x[0]) - 0x30
}
