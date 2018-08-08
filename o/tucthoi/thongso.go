package tucthoi

import (
	"LongTM/btq/btquan/x/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

type ThongSo struct {
	mongodb.BaseModel `bson:",inline"`
	NhietDo           string `json:"nhiet_do" bson:"nhiet_do"`
	DienAp            string `json:"dien_ap" bson:"dien_ap"`
	DoAm              string `json:"do_am" bson:"do_am"`
	DungLuongAcquy    string `json:"dung_luong_acquy" bson:"dung_luong_acquy"`
	DongNapTuaBin     string `json:"dong_nap_tuabin" bson:"dong_nap_tuabin"`
	DongNapPin        string `json:"dong_nap_pin" bson:"dong_nap_pin"`
	DongXa            string `json:"dong_xa" bson:"dong_xa"`
}

type ThongSoStruct struct {
	NhietDo        string `json:"nhiet_do" bson:"nhiet_do"`
	DienAp         string `json:"dien_ap" bson:"dien_ap"`
	DoAm           string `json:"do_am" bson:"do_am"`
	DungLuongAcquy string `json:"dung_luong_acquy" bson:"dung_luong_acquy"`
	DongNapTuaBin  string `json:"dong_nap_tuabin" bson:"dong_nap_tuabin"`
	DongNapPin     string `json:"dong_nap_pin" bson:"dong_nap_pin"`
	DongXa         string `json:"dong_xa" bson:"dong_xa"`
}

var ThongSoTable = mongodb.NewTable("thong_so", "tstt", 20)

func (d *ThongSo) InsertThongSo() error {
	return ThongSoTable.Create(d)
}

func UpdateThongSo(ts ThongSoStruct) error {
	_, err := ThongSoTable.UpdateAll(bson.M{}, bson.M{"$set": ts})
	return err
}

var ThongSoHSTTable = mongodb.NewTable("thong_so_hst", "tshst", 20)

func InsertThongSoHst(ts *ThongSoStruct) error {
	var tts = ThongSo{
		DienAp:         ts.DienAp,
		DoAm:           ts.DoAm,
		DongNapPin:     ts.DongNapPin,
		DongNapTuaBin:  ts.DongNapTuaBin,
		DongXa:         ts.DongXa,
		DungLuongAcquy: ts.DungLuongAcquy,
		NhietDo:        ts.NhietDo,
	}
	return ThongSoHSTTable.Create(&tts)
}
