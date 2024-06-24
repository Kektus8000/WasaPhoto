package api

import (
	"strconv"
	"strings"
)

func Authenticate(auth string) int {
	// Formato autenticazione: Bearer ID
	elements := strings.Split(auth, " ")

	// Se l'ID è del formato sbagliato, verrà ritornato un errore
	// Altrimenti, viene recuperato l'ID dell'utente
	if len(elements) == 2 {
		identifier, errConv := strconv.Atoi(elements[1])
		if errConv != nil {
			return -1
		}
		return identifier
	}
	return -1
}
