package libconfig

import "context"

type LibClientRepo interface {
	Load(ctx context.Context, appUuid string) (*Config, error)
}

type StageType int

const (
	OnURL StageType = iota + 1 // 0 will be droped by gRPC soooo +1
	OnClick
)

type Config struct {
	Stages  []Stage  `bson:"funnel"`
	BtnDefs []BtnDef `bson:"btntime"`
}

type Stage struct {
	ID         int32     `bson:"id"`
	Type       StageType `bson:"trigger"`
	Transition string    `bson:"transition"`
	Regex      string    `bson:"regex"`
}

type BtnDef struct {
	ID   int32  `bson:"id"`
	Name string `bson:"btn_name"`
}
