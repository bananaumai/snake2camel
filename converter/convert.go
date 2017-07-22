package converter

import (
	"strings"
	"regexp"
)

var re1 = regexp.MustCompile("_[a-z]")
var re2 = regexp.MustCompile("^[a-z]")

func Convert(str string, lower bool) string {
	str = re1.ReplaceAllStringFunc(str, func(m string) string {
		return strings.Trim(strings.ToUpper(m), "_")
	})
	if !lower {
		str = re2.ReplaceAllStringFunc(str, func(m string) string {
			return strings.ToUpper(m)
		})
	}
	return str
}
