package caps

import (
	"regexp"
	"strings"
)

func capitaliseSentence(s string) string {
	re := regexp.MustCompile(`[a-zA-Z]`)
	sp := regexp.MustCompile(`[ ]`)

	cap := true
	newStr := ""

	for _, d := range s {
		str := string(d)

		if cap && re.MatchString(str) {
			newStr = newStr + strings.ToUpper(str)
			cap = false
		} else if sp.MatchString(str) {
			newStr = newStr + str
			cap = true
		} else {
			newStr = newStr + str
		}
	}

	return newStr
}
