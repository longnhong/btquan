package api

import (
	"btquan/o/den"
	"btquan/x/rest"
	"github.com/gin-gonic/gin"
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
	cnum := "1"
	var ip = arr[1]
	var hengio = arr[2]
	var pin = arr[3]
	var on = arr[4]
	var d *den.Den
	for _, u := range s.Ds {
		if u.Cnum == cnum {
			u.IP = ip
			u.HenGio = hengio
			u.Online = hengio
			u.Pin = pin
			d = u
			var er = den.UpdateDen(ip, on, pin, hengio, cnum)
			rest.AssertNil(er)
		}
	}
	s.SendData(ctx, d)
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
