package svc_test

import (
	"context"
	"testing"

	"github.com/KonstantinGasser/datalab/monolith/internal/domain/user"
	"github.com/KonstantinGasser/datalab/monolith/internal/domain/user/repo"
	"github.com/google/uuid"
)

func newInmem() user.Repository {
	return repo.NewInMem()
}

func makeUUID(t *testing.T) uuid.UUID {
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}
	return uuid
}

func TestStoreUser(t *testing.T) {

	db := newInmem()

	tt := []struct {
		name    string
		newUser user.User
		err     error
	}{
		{
			name: "insert non-existing user",
			newUser: user.User{
				Uuid:         makeUUID(t),
				Username:     "test-1",
				FirstName:    "t1",
				LastName:     "t1",
				Organization: "dev-test",
				Position:     "-",
			},
			err: nil,
		},
		{
			name: "insert existing user",
			newUser: user.User{
				Uuid:         makeUUID(t),
				Username:     "test-1",
				FirstName:    "t2",
				LastName:     "t2",
				Organization: "dev-test",
				Position:     "-",
			},
			err: user.ErrDuplicatedEntry,
		},
	}

	ctx := context.Background()
	for _, tc := range tt {
		if err := db.Store(ctx, tc.newUser); err != nil {
			if err != tc.err {
				t.Fatalf("[%s] want-err: %v, got-err: %v", tc.name, tc.err, err)
			}
		}
		// check if there
		var u user.User
		if err := db.ById(ctx, tc.newUser.Uuid, &u); err != nil {
			t.Fatalf("[%s] lookup user: %v", tc.name, err)
		}
		if u.Uuid != tc.newUser.Uuid {
			t.Fatalf("[%s] user not the same: want-uuid: %v, got-uuid: %v", tc.name, tc.newUser.Uuid, u.Uuid)
		}
	}
}
