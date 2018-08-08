package api

import (
	"LongTM/btq/btquan/o/den"
	"LongTM/btq/btquan/o/tucthoi"
	"LongTM/btq/btquan/x/rest"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func InitApi(root *gin.RouterGroup) {
	NewAuthenServer(root)
}

type AuthenServer struct {
	*gin.RouterGroup
	rest.JsonRender
	den.Ds
}

func NewAuthenServer(parent *gin.RouterGroup) {
	var ds, _ = den.GetAllDen()
	var s = AuthenServer{
		RouterGroup: parent,
		Ds:          ds,
	}
	s.GET("/get_all", s.handleGetAll)
	s.GET("/update", s.handleUpdate)
	s.GET("/get_num", s.handleGetNum)
	s.GET("/get_hst_num", s.handleGetHst)
	s.POST("/insert_light", s.handleInsert)
}

func (s *AuthenServer) handleGetAll(ctx *gin.Context) {
	var ds = s.Ds
	s.SendData(ctx, ds)
}

func (s *AuthenServer) handleInsert(ctx *gin.Context) {
	var body *den.Den
	rest.AssertNil(ctx.BindJSON(&body))
	s.Ds = append(s.Ds, body)
	s.SendData(ctx, body.InsertDen())
}

func (s *AuthenServer) handleUpdate(ctx *gin.Context) {
	var request = ctx.Request.URL.Query()
	var data = request.Get("value")
	data = strings.Trim(data, "[]")
	var arr = strings.Split(data, ",")

	if len(arr) == 20 {
		var nhietDo, _ = strconv.ParseFloat(arr[6], 64)
		var doAm, _ = strconv.Atoi(arr[7])
		var daAcquy, _ = strconv.ParseFloat(arr[9], 64)
		var dlAcquy, _ = strconv.Atoi(arr[10])
		var dnPin, _ = strconv.ParseFloat(arr[11], 64)
		var dnTua, _ = strconv.ParseFloat(arr[12], 64)
		var dongXa, _ = strconv.Atoi(arr[13])
		var ts = tucthoi.ThongSoStruct{
			NhietDo:        strconv.FormatFloat(nhietDo/float64(10), 'f', -1, 64),
			DoAm:           strconv.Itoa(doAm),
			DienAp:         strconv.FormatFloat(daAcquy/float64(10), 'f', -1, 64),
			DungLuongAcquy: strconv.Itoa(dlAcquy),
			DongNapPin:     strconv.FormatFloat(dnPin/float64(10), 'f', -1, 64),
			DongNapTuaBin:  strconv.FormatFloat(dnTua/float64(10), 'f', -1, 64),
			DongXa:         strconv.Itoa(dongXa),
		}
		var err = tucthoi.UpdateThongSo(ts)
		rest.AssertNil(err)
		tucthoi.InsertThongSoHst(&ts)
	} else if len(arr) > 20 {
		var ac1 = den.ActivityDenHST{
			On:  string(arr[12][0]) + string(arr[12][1]) + ":" + string(arr[12][2]) + string(arr[12][3]),
			Off: string(arr[13][0]) + string(arr[13][1]) + ":" + string(arr[13][2]) + string(arr[13][3]),
		}
		ac1.InsertActivityDenHst()

		var ac2 = den.ActivityDenHST{
			On:  string(arr[20][0]) + string(arr[20][1]) + ":" + string(arr[20][2]) + string(arr[20][3]),
			Off: string(arr[21][0]) + string(arr[21][1]) + ":" + string(arr[21][2]) + string(arr[21][3]),
		}
		ac2.InsertActivityDenHst()
	}
	s.SendData(ctx, nil)
}

func (s *AuthenServer) handleGetHst(ctx *gin.Context) {
	var cnum = ctx.Request.URL.Query().Get("cnum")
	var denhst, err = den.GetDenHstByNum(cnum)
	rest.AssertNil(err)
	s.SendData(ctx, denhst)
}

func (s *AuthenServer) handleGetNum(ctx *gin.Context) {
	var cnum = ctx.Request.URL.Query().Get("cnum")
	var denhst, err = den.GetDenByNum(cnum)
	rest.AssertNil(err)
	s.SendData(ctx, denhst)
}
