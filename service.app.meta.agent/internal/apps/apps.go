package apps

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/library/hasher"
	"github.com/KonstantinGasser/datalab/library/utils/unique"
	"github.com/KonstantinGasser/required"
)

type InviteStatus int

const (
	// InvitePending means the request has been send but not yet acknowledged
	InvitePending InviteStatus = iota + 1 // plus one else grpc will drop zero value
	// InviteAccepted means the requested user has acknowledged and accepted the invite
	InviteAccepted
	// InviteRejected means the requested user has acknowledged and rejected the invite
	InviteRejected
)

var (
	ErrNotAppOwner           = fmt.Errorf("user is not the owner of App")
	ErrNotAMember            = fmt.Errorf("user is not a member of App")
	ErrRequiredFieldsMissing = fmt.Errorf("missing fields to create App")
	ErrReadAccess            = fmt.Errorf("user has not read access on App")
	ErrAlreadyMember         = fmt.Errorf("user is already member of App")
)

type AppsRepository interface {
	Store(ctx context.Context, app App) error
	GetById(ctx context.Context, uuid string, stored interface{}) error
	GetAll(ctx context.Context, userUuid string, stored interface{}) error
	SetAppLock(ctx context.Context, uuid string) error
	AddMember(ctx context.Context, appUuid string, invitedMember Member) error
	MemberStatus(ctx context.Context, appUuid string, openInvite Member) error
}

type App struct {
	// mongoDB pk (document key)
	Uuid        string   `bson:"_id" required:"yes"`
	Name        string   `bson:"name" required:"yes"`
	URL         string   `bson:"url" required:"yes"`
	OwnerUuid   string   `bson:"owner_uuid" required:"yes"`
	OwnerOrgn   string   `bson:"owner_orgn" required:"yes"`
	Description string   `bson:"description" required:"yes"`
	Members     []Member `bson:"member"`
	Hash        string   `bson:"hash" required:"yes"`
	Locked      bool     `bson:"locked"`
}

// Member refers to a user added to an app
type Member struct {
	Uuid   string       `bson:"uuid"`
	Status InviteStatus `bson:"status"`
}

// AppSubset is a light weighted representation
// of the app
type AppSubset struct {
	Uuid string
	Name string
}

// NewDefault creates a new App with its default values and initialized the App with
// a UUID and App Hash (hash of orgn/name)
func NewDefault(name, URL, ownerUuid, ownerOrgn, desc string) (*App, error) {
	app := App{
		Name:        name,
		URL:         URL,
		OwnerUuid:   ownerUuid,
		OwnerOrgn:   ownerOrgn,
		Description: desc,
		Members: []Member{ // Owner is listed as member per default
			{
				Uuid:   ownerUuid,
				Status: InviteAccepted,
			},
		},
		Locked: false,
	}
	app.Init()
	if err := required.Atomic(&app); err != nil {
		return nil, ErrRequiredFieldsMissing
	}
	return &app, nil
}

// SubsetOf returns a slice of AppSubset created from the
// input apps
func SubsetOf(apps ...App) []AppSubset {
	var appSubsets = make([]AppSubset, len(apps))
	for i, app := range apps {
		appSubsets[i] = AppSubset{
			Uuid: app.Uuid,
			Name: app.Name,
		}
	}
	return appSubsets
}

// Init initializes the App with its computed properties such as
// its UUID and Hash
func (app *App) Init() {
	app.Uuid, _ = unique.UUID()
	app.Hash = hasher.Build(app.Name, app.OwnerOrgn)
}

// AddInvite appends the App.Member slice if the user if not
// already listed as member of App
func (app *App) AddInvite(userUuid string) (*Member, error) {
	if ok := app.IsNotMember(userUuid); !ok {
		return nil, ErrAlreadyMember
	}
	member := Member{
		Uuid:   userUuid,
		Status: InvitePending,
	}
	app.Members = append(app.Members, member)
	return &member, nil
}

// InviteReminderOk checks if a user qualifys to be send an reminder
// for an invite again
func (app App) InviteReminderOk(userUuid string) bool {
	for _, member := range app.Members {
		if member.Uuid == userUuid && member.Status == InvitePending {
			return true
		}
	}
	return false
}

func (app App) HasReadAccess(userUuid string) error {
	errO := app.IsOwner(userUuid)
	errM := app.IsMember(userUuid, InviteAccepted)
	if errM != nil && errO != nil {
		return ErrReadAccess
	}
	return nil
}

func (app App) IsOwner(userUuid string) error {
	if app.OwnerUuid != userUuid {
		return ErrNotAppOwner
	}
	return nil
}

func (app App) IsMember(userUuid string, status InviteStatus) error {
	for _, member := range app.Members {
		if member.Uuid == userUuid && member.Status == status {
			return nil
		}
	}
	return ErrNotAMember
}

func (app App) OpenInvite(userUuid string) *Member {
	for _, member := range app.Members {
		if member.Uuid == userUuid && member.Status == InvitePending {
			return &member
		}
	}
	return nil
}

func (app App) IsLocked() bool { return app.Locked }

// IsNotMember checks if user is listed as member of app regardles of InviteStatus
func (app App) IsNotMember(userUuid string) bool {
	for _, member := range app.Members {
		if member.Uuid == userUuid {
			return false
		}
	}
	return true
}
