package appconfigs

import (
	"context"
	"fmt"
	"regexp"
)

const (
	FlagFunnel   = "funnel"
	FlagCampaign = "campaign"
	FlagBtnTime  = "btntime"
)

var (
	ErrInvalidFlag       = fmt.Errorf("provided config-flag is invalid")
	ErrNoReadWriteAccess = fmt.Errorf("user read/write access for AppToken")
	ErrNoReadAccess      = fmt.Errorf("user has no read access for AppToken")

	ErrInvalidRegex = fmt.Errorf("provided stage-regex is not a regex")
)

type AppconfigRepo interface {
	Initialize(ctx context.Context, appConfig AppConfig) error
	GetById(ctx context.Context, uuid string, result interface{}) error
	UpdateByFlag(ctx context.Context, uuid, flag string, data interface{}) error

	AddMember(ctx context.Context, uuid, userUuid string) error
	RollbackAddMember(ctx context.Context, uuid, userUuid string) error

	SetAppConfigLock(ctx context.Context, uuid string, lock bool) error
}

type AppConfig struct {
	AppUuid     string   `bson:"_id"`
	ConfigOwner string   `bson:"config_owner"`
	Member      []string `bson:"member"`
	Locked      bool     `bson:"locked"`
	Funnel      []Stage  `bson:"funnel"`
	Campaign    []Record `bson:"campaign"`
	BtnTime     []BtnDef `bson:"btntime"`
}

type Stage struct {
	Id         int32  `bson:"id"`
	Name       string `bson:"name"`
	Transition string `bson:"transition"`
	Regex      string `bson:"regex"`
	Trigger    int32  `bson:"trigger"`
}

type Record struct {
	Id     int32  `bson:"id"`
	Name   string `bson:"name"`
	Suffix string `bson:"suffix"`
}

type BtnDef struct {
	Id      int32  `bson:"id"`
	Name    string `bson:"name"`
	BtnName string `bson:"btn_name"`
}

// NewDefault creates a new AppConfig with only meta data
func NewDefault(appRefUuid, configOwner string) *AppConfig {
	return &AppConfig{
		AppUuid:     appRefUuid,
		ConfigOwner: configOwner,
		Member:      make([]string, 0),
		Funnel:      make([]Stage, 0),
		Campaign:    make([]Record, 0),
		BtnTime:     make([]BtnDef, 0),
	}
}

func (appConf *AppConfig) ApplyFunnel(stages ...Stage) error {
	// check if regex is set and if regex is valid
	for _, stage := range stages {
		if len(stage.Regex) > 0 {
			if err := isRegex(stage.Regex); err != nil {
				return err
			}
		}
	}
	appConf.Funnel = stages
	return nil
}

func (appConf *AppConfig) ApplyCampaign(records ...Record) {
	appConf.Campaign = records
}

func (appConf *AppConfig) ApplyBtnTime(btnDefs ...BtnDef) {
	appConf.BtnTime = btnDefs
}

// AddMember appends the member list of an AppConfig. Skips operation if already exists
func (appConf *AppConfig) AddMember(userUuid string) {
	for _, member := range appConf.Member {
		if member == userUuid {
			return
		}
	}
	appConf.Member = append(appConf.Member, userUuid)
}

// HasReadWrite checks if the provided user uuid is listed as owner of
// AppToken
func (appConfig AppConfig) HasReadWrite(userUuid string) error {
	if appConfig.ConfigOwner != userUuid {
		return ErrNoReadWriteAccess
	}
	return nil
}

// HasRead checks if the user has read access on the AppToken
func (appConfig AppConfig) HasRead(userUuid string) error {
	for _, member := range appConfig.Member {
		if member == userUuid {
			return nil
		}
	}
	return ErrNoReadAccess
}

// HasReadOrWrite checks if the user has either read or write acces on the AppToken
func (appConfig AppConfig) HasReadOrWrite(userUuid string) error {
	readErr := appConfig.HasRead(userUuid)
	rwErr := appConfig.HasReadWrite(userUuid)
	if readErr != nil && rwErr != nil {
		return ErrNoReadAccess
	}
	return nil
}

// isRegex checks if a given pattern can be parsed to a regex
func isRegex(pattern string) error {
	_, err := regexp.Compile(pattern)
	if err != nil {
		return ErrInvalidRegex
	}
	return nil
}
