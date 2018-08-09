package hen_gio

import (
	"LongTM/btq/btquan/o/tucthoi"
	"LongTM/btq/btquan/x/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

var HenGioTable = mongodb.NewTable("hen_gio", "hen", 20)
var HenGioHstTable = mongodb.NewTable("hen_gio_hst", "hg_hst", 20)

type HenGio struct {
	mongodb.BaseModel `bson:",inline"`
	HGio              `bson:",inline"`
}

type HGio struct {
	tucthoi.ThongSoStruct `bson:",inline"`
	Hens                  []*HenGioDen `json:"hen_gios" bson:"hen_gios"`
}

type HenGioDen struct {
	DenNum string `json:"den_so" bson:"den_so"`
	GioBat string `json:"gio_bat" bson:"gio_bat"`
	GioTat string `json:"gio_tat" bson:"gio_tat"`
}

func GetAll() ([]*HenGio, error) {
	var hs []*HenGio
	var err = HenGioTable.FindWhere(bson.M{}, &hs)
	return hs, err
}

func InsertHenGio(d HGio) error {
	var hs, _ = GetAll()
	if hs != nil && len(hs) > 0 {
		UpdateHenGio(&d)
	}
	var hg = &HenGio{
		HGio: d,
	}
	return HenGioTable.Create(hg)
}

func UpdateHenGio(u *HGio) error {
	var _, err = HenGioTable.UpdateAll(bson.M{}, bson.M{"$set": u})
	return err
}

func GetHenGio() (*HenGio, error) {
	var h *HenGio
	var err = HenGioTable.FindOne(bson.M{}, &h)
	return h, err
}

func (d *HenGio) InsertHst() error {
	return HenGioHstTable.Create(d)
}

func GetAllHst() ([]*HenGio, error) {
	var hgs []*HenGio
	return hgs, HenGioHstTable.FindWhere(bson.M{}, &hgs)
}
