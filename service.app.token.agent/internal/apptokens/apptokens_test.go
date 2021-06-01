package apptokens

import (
	"testing"

	"github.com/matryer/is"
)

func TestAppTokenIssuing(t *testing.T) {
	is := is.New(t)

	apptoken, _ := NewDefault("someUUID", "someHash", "someOwner")
	newtoken, _ := apptoken.Issue()

	// issue must fail since the app token must still be valid
	_, err := newtoken.Issue()
	is.Equal(err, ErrAppTokenStillValid)
}
