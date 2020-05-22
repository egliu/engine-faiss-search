package utils

import "regexp"

var (
	patternLegalNodeName = regexp.MustCompile(`^[0-9a-zA-Z_.\[\]:-]+$`)
)

func ValidateZkNodeName(n string) bool {
	return patternLegalNodeName.MatchString(n)
}
