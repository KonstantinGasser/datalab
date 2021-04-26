package user

import "regexp"

// orgnAllowed check if the provided organization domain follows the
// allowed syntax
func orgnAllowed(orgnDomain string) bool {
	re := regexp.MustCompile("/")
	matches := re.Find([]byte(orgnDomain))
	return matches == nil
}
