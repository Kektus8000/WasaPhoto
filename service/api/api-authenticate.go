package api

import (
	"strconv"
	"strings"
)

func Authenticate(userID int, auth string) bool {
	elements := strings.Split(auth, " ")
	if len(elements) == 2 {
		identifier, errConv := strconv.Atoi(elements[1])
		if errConv != nil {
			return false
		}
		return userID == identifier
	}
	return false
}
