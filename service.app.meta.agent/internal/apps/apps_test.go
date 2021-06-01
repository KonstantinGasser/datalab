package apps

import (
	"testing"
)

func TestAddInvite(t *testing.T) {
	// is := is.New(t)

	app, _ := NewDefault("test", "url", "0", "domain", "some")

	_, _ = app.AddInvite("1")

}
