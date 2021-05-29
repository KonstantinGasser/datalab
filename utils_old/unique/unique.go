package unique

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// UUID generate a new random user id as document id for the user
// generates a NewV4 as defined in the github.com/gofrs/uuid package
func UUID() (string, error) {
	UUID, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("could not generate UUID for user: %v", err)
	}
	return UUID.String(), nil
}
