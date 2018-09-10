package tucthoi

import (
	"LongTM/btq/btquan/x/db/mongodb"

	"gopkg.in/mgo.v2/bson"
)

type ThongSo struct {
	mongodb.BaseModel `bson:",inline"`
	ThongSoTemp       `bson:",inline"`
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

type ThongSoTemp struct {
	ThongSoStruct `bson:",inline"`
	Status1       string `json:"status1" bson:"status1"`
	Status2       string `json:"status2" bson:"status2"`
}

var ThongSoTable = mongodb.NewTable("thong_so", "tstt", 20)

func (d *ThongSo) InsertThongSo() error {
	return ThongSoTable.Create(d)
}

func GetThongSo() (*ThongSo, error) {
	var ts *ThongSo
	err := ThongSoTable.FindOne(bson.M{}, &ts)
	return ts, err
}

func UpdateThongSo(ts ThongSoTemp) error {
	_, err := ThongSoTable.UpdateAll(bson.M{}, bson.M{"$set": ts})
	return err
}

var ThongSoHSTTable = mongodb.NewTable("thong_so_hst", "tshst", 20)

func InsertThongSoHst(ts *ThongSoTemp) error {
	var tts = ThongSo{}
	tts.DienAp = ts.DienAp
	tts.DoAm = ts.DoAm
	tts.DongNapPin = ts.DongNapPin
	tts.DongNapTuaBin = ts.DongNapTuaBin
	tts.DongXa = ts.DongXa
	tts.DungLuongAcquy = ts.DungLuongAcquy
	tts.NhietDo = ts.NhietDo
	tts.Status1 = ts.Status1
	tts.Status2 = ts.Status2
	return ThongSoHSTTable.Create(&tts)
}

func GetThongSoHst() ([]*ThongSo, error) {
	var ts []*ThongSo
	err := ThongSoTable.FindWhere(bson.M{}, &ts)
	return ts, err
}
