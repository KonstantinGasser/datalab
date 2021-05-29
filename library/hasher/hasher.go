package hasher

import (
	"strings"

	"github.com/KonstantinGasser/datalab/library/utils/hash"
)

// Build generates the app hash
func Build(appName string, organization string) string {
	concated := strings.Join([]string{appName, organization}, "/")
	return hash.Sha256([]byte(concated)).String()
}
