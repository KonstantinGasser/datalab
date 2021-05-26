package domain

type ConfigInfo struct {
	AppUuid  string   `bson:"_id"`
	Funnel   []Stage  `bson:"funnel"`
	Campaign []Record `bson:"campaign"`
	BtnTime  []BtnDef `bson:"btn_time"`
}

type Stage struct {
	mustImpleToBeValidConfig
	ID         int32  `bson:"id"`
	Name       string `bson:"name"`
	Transition string `bson:"transition"`
}

type Record struct {
	mustImpleToBeValidConfig
	ID     int32  `bson:"id"`
	Name   string `bson:"name"`
	Suffix string `bson:"suffix"`
}

type BtnDef struct {
	mustImpleToBeValidConfig
	ID      int32  `bson:"id"`
	Name    string `bson:"name"`
	BtnName string `bson:"btn_name"`
}

type Config interface {
	MustImpleToBeValidConfig()
}

type mustImpleToBeValidConfig struct{}

func (mustImpleToBeValidConfig) MustImpleToBeValidConfig() {}
