package helpers

import (
	"strings"
)

func makeSlug(str string) string {
	str = strings.ToLower(str)
	return str
}
