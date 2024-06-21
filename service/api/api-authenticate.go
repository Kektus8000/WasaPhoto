package api

import (
	"strconv"
	"strings"
)

func Authenticate(auth string) int {
	elements := strings.Split(auth, " ")
	if len(elements) == 2 {
		identifier, errConv := strconv.Atoi(elements[1])
		if errConv != nil {
			return -1
		}
		return identifier
	}
	return -1
}
