package helps

import "strings"

func VerifyContractAddress(s string) string {
	if s == "" || strings.TrimSpace(s) == "" {
		return "0x0000000000000000000000000000000000000000"
	}
	return s
}
