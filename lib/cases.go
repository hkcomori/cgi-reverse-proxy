package lib

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// UpperSnakeCaseToTrainCase converts string format: UPPER_SNAKE_CASE to Train-Case
func UpperSnakeCaseToTrainCase(upperSnakeCase string) string {
	words := strings.Split(upperSnakeCase, "_")
	for i, word := range words {
		words[i] = cases.Title(language.Und).String(strings.ToLower(word))
	}
	return strings.Join(words, "-")
}
