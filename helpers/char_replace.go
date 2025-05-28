package helpers

import (
	"strings"
)

func CharReplace(text string) string {
	chars := []string{"?", "/", `/\/`, ":", "<", ">", `"`, "|", "*"}
	for _, c := range chars {
		text = strings.ReplaceAll(text, c, " ")
	}
	return text
}
