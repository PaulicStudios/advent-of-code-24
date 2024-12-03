package utils

import "strconv"

func ConvertToInt(s string) int {
	nbr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nbr
}
