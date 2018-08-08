package video

import (
	"LongTM/btq/btquan/o/auth/user"
	"LongTM/btq/btquan/x/db/mongodb"
)

var VideoTable = mongodb.NewTable("tk_booking", "tkbk", 18)

type Video struct {
	mongodb.BaseModel `bson:",inline"`
	Ip                string `json:"ip" bson:"ip"`
}

type VideoHst struct {
	ServiceID string       `json:"service_id" bson:"service_id"`
	BranchID  string       `json:"branch_id"  bson:"branch_id"`
	MTime     int64        `json:"ctime"  bson:"ctime"`
	Status    BookingState `json:"status"  bson:"status"`
}

type RateVideo struct {
	Video `bson:",inline"`
}

type UpdateCetm struct {
	BVideoID     string       `bson:"bVideo_id" json:"bVideo_id"`
	Teller       string       `json:"teller"  bson:"teller"`
	AvatarTeller string       `json:"avatar_teller"  bson:"avatar_teller"`
	TellerID     string       `json:"teller_id"  bson:"teller_id"`
	IdVideoCetm  string       `json:"id_Video_cetm"  bson:"id_Video_cetm"`
	Tracks       []VideoHst   `json:"tracks"  bson:"tracks"`
	CnumCetm     string       `json:"cnum_cetm"  bson:"cnum_cetm"`
	Status       BookingState `json:"status"  bson:"status"`
	ServingTime  int64        `json:"serving_time"  bson:"serving_time"`
	WaitingTime  int64        `json:"waiting_time"  bson:"waiting_time"`
}

type WhenCreateVideoInCetm struct {
	CheckInAt   int64  `json:"check_in_at"  bson:"check_in_at"`
	IdVideoCetm string `json:"id_Video_cetm"  bson:"id_Video_cetm"`
	CnumCetm    string `json:"cnum_cetm"  bson:"cnum_cetm"`
}

type VideoUpdate struct {
	UpdatedAt     int64        `json:"updated_at" bson:"updated_at"`
	BVideoID      string       `bson:"bVideo_id" json:"bVideo_id"`
	TimeGoBank    int64        `bson:"time_go_bank" json:"time_go_bank"`
	ServiceID     string       `json:"service_id" bson:"service_id"`
	ServiceName   string       `json:"service_name" bson:"service_name"`
	BranchName    string       `json:"branch_name"  bson:"branch_name"`
	BranchID      string       `json:"branch_id"  bson:"branch_id"`
	BranchAddress string       `json:"branch_address"  bson:"branch_address"`
	TypeVideo     TypeVideo    `json:"type_Video" bson:"type_Video"`
	Tracks        []VideoHst   `json:"tracks"  bson:"tracks"`
	Status        BookingState `json:"status"  bson:"status"`
}

type VideoSchedule struct {
	IdBranch    string `json:"id" bson:"_id"`
	CountPeople int    `json:"count" bson:"count"`
}

type VideoCreate struct {
	Customer      string    `bson:"customer" json:"customer"`
	TimeGoBank    int64     `bson:"time_go_bank" json:"time_go_bank"`
	BranchAddress string    `json:"branch_address"  bson:"branch_address"`
	BranchName    string    `json:"branch_name"  bson:"branch_name"`
	ServiceID     string    `json:"service_id" bson:"service_id"`
	ServiceName   string    `json:"service_name" bson:"service_name"`
	BranchID      string    `json:"branch_id"  bson:"branch_id"`
	Lang          string    `json:"lang" bson:"lang"`
	CustomerID    string    `json:"customer_id" bson:"customer_id"`
	TypeVideo     TypeVideo `json:"type_Video" bson:"type_Video"`
}

type BookingState string

type TypeVideo string

type TypeRate int

const (
	TypeDefaultRate = TypeRate(0)
	TypeRated       = TypeRate(1)
	TypeNoRate      = TypeRate(2)
)

const (
	TypeNow      = TypeVideo("book_now")
	TypeSchedule = TypeVideo("book_schedule")
)

const (
	BookingStateCreated    = BookingState("created")     // vừa tạo
	BookingStateConfirmed  = BookingState("confirmed")   // đã xác nhận đến ngân hàng
	BookingStateSancelled  = BookingState("cancelled")   // cetm gọi ko có mặt
	BookingStateDelete     = BookingState("deleted")     // kh xóa vé
	BookingStateFinished   = BookingState("finished")    // đã giao dịch xong
	BookingStateNotArrived = BookingState("not_arrived") // không đến
	BookingStateCheckCode  = BookingState("check_code")  // CheckCode
	BookingStateCreateCetm = BookingState("create_cetm") // CheckCode
	BookingCustomerUpdate  = BookingState("cus_update")  //kh update
)

type VideoBranches struct {
	BranchID string
	Videos   []*Video
}

type VideoUser struct {
	Video    `bson:",inline"`
	Customer *user.User `json:"customer" bson:"customer"`
}

type NotifyVideo struct {
	Video *Video `json:"Video_booking"`
}
