package den

import (
	"LongTM/btq/btquan/x/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

var DenTable = mongodb.NewTable("dens", "d", 20)
var DenHstTable = mongodb.NewTable("den_hst", "dhst", 20)
var ActivityDenHstTable = mongodb.NewTable("activity_den_hst", "achst", 20)

type Ds []*Den

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

func UpdateDen(ip string, online string, pin string, hengio string, cnum string) error {
	var hst = &DenHST{
		Cnum:   cnum,
		IP:     ip,
		Online: online,
		HenGio: hengio,
		Pin:    pin,
	}
	hst.InsertDenHst()
	return DenTable.Update(bson.M{"cnum": cnum},
		bson.M{"online": online,
			"pin":     pin,
			"hen_gio": hengio,
			"ip":      cnum,
		})
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
func GetDenHstByNum(num string) ([]*DenHST, error) {
	var a []*DenHST
	return a, DenHstTable.FindWhere(bson.M{"cnum": num}, &a)
}

type ActivityDenHST struct {
	mongodb.BaseModel `bson:",inline"`
	On                string `json:"on" bson:"on"`
	Off               string `json:"off" bson:"off"`
}

func (d *ActivityDenHST) InsertActivityDenHst() error {
	return ActivityDenHstTable.Create(d)
}
