package setup

import (
	tt "LongTM/btq/btquan/o/tucthoi"
	"LongTM/btq/btquan/x/db/mongodb"

	"gopkg.in/mgo.v2/bson"
)

var SetupTable = mongodb.NewTable("setup", "set", 20)

type Setup struct {
	mongodb.BaseModel `bson:",inline"`
	UpdateSetup       `bson:",inline"`
}

type UpdateSetup struct {
	tt.ThongSoStruct `bson:",inline"`
	TimeOn1          string `json:"time_on1" bson:"time_on1"`
	TimeOff1         string `json:"time_off1" bson:"time_off1"`
	TimeOn2          string `json:"time_on2" bson:"time_on2"`
	TimeOff2         string `json:"time_off2" bson:"time_off2"`
	ValueClient      string `json:"value_client" bson:"value_client"`
	TimeUpClient     int64  `json:"time_client" bson:"time_client"`
	Value            string `json:"value" bson:"value"`
	TimeUpServer     int64  `json:"time_up_server" bson:"time_up_server"`
	Manual           string `json:"manual" bson:"manual"`
	ManualOn1        string `json:"manual_on1" bson:"manual_on1"`
	ManualOn2        string `json:"manual_on2" bson:"manual_on2"`
	IntUpServer      int    `json:"-" bson:"int_up_server"`
	IntUpClient      int    `json:"-" bson:"int_up_client"`
	TimeOldOn1       string `json:"-" bson:"time_old_on1"`
	TimeOldOff1      string `json:"-" bson:"time_old_off1"`
	TimeOldOn2       string `json:"-" bson:"time_old_on2"`
	TimeOldOff2      string `json:"-" bson:"time_old_off2"`
	IsUp             bool   `json:"-" bson:"is_up"`
}

func InsertSetup(d UpdateSetup) error {
	var s, _ = GetSetup()
	if s != nil {
		return Update(s.ID, d)
	}
	var ds = Setup{
		UpdateSetup: d,
	}
	return SetupTable.Create(&ds)
}

func GetSetup() (*Setup, error) {
	var set *Setup
	return set, SetupTable.FindOne(bson.M{}, &set)
}

func Update(id string, set UpdateSetup) error {
	return SetupTable.UnsafeUpdateByID(id, set)
}
