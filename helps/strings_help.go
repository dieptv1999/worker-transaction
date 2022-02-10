package helps

import (
	"net/mail"
	"strings"
)

func IsEmptyString(str string) bool {
	return strings.Compare(str, "") == 0
}

func IsEmailString(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}
