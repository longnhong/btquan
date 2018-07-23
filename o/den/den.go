package den

import (
	"btquan/x/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

var DenTable = mongodb.NewTable("dens", "d", 20)
var DenHstTable = mongodb.NewTable("den_hst", "dhst", 20)

type Den struct {
	mongodb.BaseModel `bson:",inline"`
	IP                string `json:"ip" bson:"ip"`
	Address           string `json:"address" bson:"address"`
	Cnum              string `json:"cnum" bson:"cnum"`
	Online            string `json:"online" bson:"online"`
	Pin               string `json:"pin" bson:"pin"`
	HenGio            string `json:"hen_gio" bson:"hen_gio"`
}

func (d *Den) InsertDen() error {
	return DenTable.Create(d)
}

func GetAllDen() ([]*Den, error) {
	var a []*Den
	return a, DenTable.FindWhere(bson.M{}, &a)
}
func GetDenByNum(num string) (*Den, error) {
	var a *Den
	return a, DenTable.FindWhere(bson.M{"cnum": num}, &a)
}

type DenHST struct {
	mongodb.BaseModel `bson:",inline"`
	IP                string `json:"ip" bson:"ip"`
	Cnum              string `json:"cnum" bson:"cnum"`
	Online            string `json:"online" bson:"online"`
	Pin               string `json:"pin" bson:"pin"`
	HenGio            string `json:"hen_gio" bson:"hen_gio"`
}

func (d *DenHST) InsertDenHst() error {
	return DenHstTable.Create(d)
}

func GetAllDenHst() ([]*DenHST, error) {
	var a []*DenHST
	return a, DenHstTable.FindWhere(bson.M{}, &a)
}
func GetDenHstByNum(num string) (*DenHST, error) {
	var a *DenHST
	return a, DenHstTable.FindWhere(bson.M{"cnum": num}, &a)
}
