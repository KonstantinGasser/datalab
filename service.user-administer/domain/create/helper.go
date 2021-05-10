package create

import "regexp"

func orgnNameAllowed(orgnDomain string) bool {
	re := regexp.MustCompile("/")
	matches := re.Find([]byte(orgnDomain))
	return matches == nil
}
