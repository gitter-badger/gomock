package gomock

import (
	"bytes"
	"regexp"
	"strings"
)

// newRegepxRoute parse a route and returns a routeRegexp
func newRegexRoute(route string) (*regexp.Regexp, error) {
	cnt := strings.Count(route, "?")
	var routetpl = make([]byte, 0, len(route)+cnt)

	for i, uPoint := range route {
		switch ch := string(uPoint); {
		case ch == "?":
			routetpl = append(routetpl, '\\', route[i])
		case ch == "}":
			cnt := 0
			for _, v := range routetpl {
				if string(v) == "{" {
					break
				}
				cnt++
			}

			routetpl = routetpl[:cnt]
			routetpl = append(routetpl, '[', '^', '/', '&', ']', '+', '?')
		default:
			for _, v := range []byte(string(uPoint)) {
				routetpl = append(routetpl, v)
			}
		}
	}

	routetpl = append(routetpl, '$')
	return regexp.Compile(bytes.NewBuffer(routetpl).String())
}
