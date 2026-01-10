package chars

import (
	"regexp"
	"strings"
)

// OnlyNumbers Удаляем все символы, кроме цифр
func OnlyNumbers(varchar string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(varchar, "")
}

func CheckStringInTexts(searchString string, texts []string) bool {
	for _, text := range texts {
		if strings.Contains(searchString, text) {
			return true
		}
	}
	return false
}
