package brute

import (
	"strings"
)

func generateNewHEX(hex string) string {
	// Splitted hex []string
	sh := strings.Split(hex, "")
	possible := "0123456789abcdef"

	for i := len(hex) - 1; i >= 0; i-- {
		point := strings.Index(possible, sh[i])
		if point == 15 {
			sh[i] = "0"
		} else {
			sh[i] = string(possible[point+1])
			break
		}
	}
	return strings.Join(sh, "")
}
