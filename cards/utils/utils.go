package utils

import (
	"slices"
)

func ReverseString(card string) string {
	my_string := []rune(card)
	slices.Reverse(my_string)
	return string(my_string)
}
