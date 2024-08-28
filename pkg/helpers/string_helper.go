package helpers

import (
	"strings"
)

func makeSlug(str string) string {

	return strings.ToLower(str)
}
