package api

import (
	"LongTM/btq/btquan/o/den"
	"LongTM/btq/btquan/o/hen_gio"
	"LongTM/btq/btquan/o/setup"
	"LongTM/btq/btquan/o/tucthoi"
	"LongTM/btq/btquan/x/rest"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitApi(root *gin.RouterGroup) {
	NewAuthenServer(root)
}

type AuthenServer struct {
	*gin.RouterGroup
	rest.JsonRender
	den.Ds
	*setup.Setup
}

func NewAuthenServer(parent *gin.RouterGroup) {
	var ds, _ = den.GetAllDen()
	var set, _ = setup.GetSetup()
	var s = AuthenServer{
		RouterGroup: parent,
		Ds:          ds,
		Setup:       set,
	}
	s.GET("/get_all", s.handleGetAll)
	s.GET("/update", s.handleUpdate)
	s.POST("/server_update", s.handleServerUpdate)
	s.GET("/get_setup", s.handleSettup)
	s.GET("/get_thongso", s.handleGetThongSo)
	s.GET("/get_num", s.handleGetNum)
	s.GET("/get_hst_num", s.handleGetHst)
	s.GET("/get_hst", s.handleGetAllHst)
	s.GET("/get_ts_hst", s.handleGetTsHst)
	s.GET("/get_den_hst", s.handleDenHst)
	s.GET("/get_hen_gio", s.handleGetHenGio)
	s.POST("/hen_gio", s.handleHenGio)

	s.POST("/insert_light", s.handleInsert)
}
func (s *AuthenServer) handleSettup(ctx *gin.Context) {
	s.SendData(ctx, s.Setup)
}
func (s *AuthenServer) handleHenGio(ctx *gin.Context) {
	var body = hen_gio.HGio{}
	rest.AssertNil(ctx.BindJSON(&body))
	err := hen_gio.InsertHenGio(body)
	rest.AssertNil(err)
	var hg = &hen_gio.HenGio{
		HGio: body,
	}
	hg.InsertHst()
	s.SendData(ctx, nil)
}

func (s *AuthenServer) handleGetHenGio(ctx *gin.Context) {
	var ts, err = hen_gio.GetHenGio()
	rest.AssertNil(err)
	s.SendData(ctx, ts)
}

func (s *AuthenServer) handleGetTsHst(ctx *gin.Context) {
	var ts, err = tucthoi.GetThongSoHst()
	rest.AssertNil(err)
	s.SendData(ctx, ts)
}

func (s *AuthenServer) handleDenHst(ctx *gin.Context) {
	var ts, err = den.GetAllDenHst()
	rest.AssertNil(err)
	s.SendData(ctx, ts)
}

func (s *AuthenServer) handleGetAll(ctx *gin.Context) {
	var ds = s.Ds
	s.SendData(ctx, ds)
}

func (s *AuthenServer) handleGetAllHst(ctx *gin.Context) {
	var ds, err = den.GetAllDenHst()
	rest.AssertNil(err)
	s.SendData(ctx, ds)
}

func (s *AuthenServer) handleGetThongSo(ctx *gin.Context) {
	var ts, err = tucthoi.GetThongSo()
	rest.AssertNil(err)

	s.SendData(ctx, map[string]interface{}{
		"thong_so": ts,
		"dens":     s.Ds,
	})
}

func (s *AuthenServer) handleInsert(ctx *gin.Context) {
	var body *den.Den
	rest.AssertNil(ctx.BindJSON(&body))
	s.Ds = append(s.Ds, body)
	s.SendData(ctx, body.InsertDen())
}

func (s *AuthenServer) handleServerUpdate(ctx *gin.Context) {
	var body = setup.UpdateSetup{}
	rest.AssertNil(ctx.BindJSON(&body))
	var set = setup.Setup{
		UpdateSetup: body,
	}
	set.IntUpServer = s.Setup.IntUpServer + 1
	s.ValueClient = s.Setup.ValueClient
	set.TimeOldOff1 = s.Setup.TimeOldOff1
	set.TimeOldOff2 = s.Setup.TimeOldOff2
	set.TimeOldOn1 = s.Setup.TimeOldOn1
	set.TimeOldOn2 = s.Setup.TimeOldOn2
	s.Setup = &set
	s.SendData(ctx, setup.InsertSetup(body))
}

func (s *AuthenServer) handleUpdate(ctx *gin.Context) {
	var request = ctx.Request.URL.Query()
	var data = request.Get("value")
	data = strings.Trim(data, "[]")
	var arr = strings.Split(data, ",")
	fmt.Printf("ARR", len(arr))
	//var isCheck = true
	if len(arr) == 20 {
		var nhietDo, _ = strconv.ParseFloat(arr[6], 64)
		var doAm, _ = strconv.Atoi(arr[7])
		var daAcquy, _ = strconv.ParseFloat(arr[9], 64)
		var dlAcquy, _ = strconv.Atoi(arr[10])
		var dnPin, _ = strconv.ParseFloat(arr[11], 64)
		var dnTua, _ = strconv.ParseFloat(arr[12], 64)
		var dongXa, _ = strconv.Atoi(arr[13])
		var status1 = arr[17]
		var status2 = arr[18]
		if status1 == "0001" {
			status1 = "online"
		} else {
			status1 = "offline"
		}
		if status2 == "0001" {
			status2 = "online"
		} else {
			status2 = "offline"
		}
		var ts = tucthoi.ThongSoTemp{
			Status1: status1,
			Status2: status2,
		}
		ts.NhietDo = strconv.FormatFloat(nhietDo/float64(10), 'f', -1, 64)
		ts.DoAm = strconv.Itoa(doAm)
		ts.DienAp = strconv.FormatFloat(daAcquy/float64(10), 'f', -1, 64)
		ts.DungLuongAcquy = strconv.Itoa(dlAcquy)
		ts.DongNapPin = strconv.FormatFloat(dnPin/float64(10), 'f', -1, 64)
		ts.DongNapTuaBin = strconv.FormatFloat(dnTua/float64(10), 'f', -1, 64)
		ts.DongXa = strconv.Itoa(dongXa)
		var err = tucthoi.UpdateThongSo(ts)
		rest.AssertNil(err)
		tucthoi.InsertThongSoHst(&ts)
		s.SendData(ctx, nil)
	} else if len(arr) == 49 {
		if s.Setup != nil && s.Setup.ValueClient == data && s.Setup.IntUpServer == s.Setup.IntUpClient {
			fmt.Println("Vô Trùng")
			return
		} else if s.Setup != nil && s.Setup.ValueClient != data && s.Setup.IntUpServer == s.Setup.IntUpClient {
			fmt.Println("Vô cài đặt phía client")
			var nhietDo, _ = strconv.Atoi(arr[6])
			var doAm, _ = strconv.Atoi(arr[7])

			var settup = setup.UpdateSetup{}
			settup.NhietDo = strconv.Itoa(nhietDo)
			settup.DoAm = strconv.Itoa(doAm)
			settup.TimeOn1 = arr[12]
			fmt.Printf("Time On1", arr[12])
			settup.TimeOff1 = arr[13]
			fmt.Printf("Time Off 1", arr[13])
			settup.TimeOn2 = arr[20]
			fmt.Printf("Time On2", arr[20])
			settup.TimeOff2 = arr[21]
			fmt.Printf("Time Off 2", arr[21])
			settup.ValueClient = data

			settup.Manual = arr[43]
			settup.ManualOn1 = arr[44]
			settup.ManualOn2 = arr[45]
			var set = setup.Setup{
				UpdateSetup: settup,
			}

			var a = strings.Split(settup.TimeOn1, ":")
			var timeOn1 = settup.TimeOn1
			if len(a) > 1 {
				timeOn1 = a[0] + a[1]
			}
			var b = strings.Split(settup.TimeOn2, ":")
			var timeOn2 = settup.TimeOn2
			if len(b) > 1 {
				timeOn2 = b[0] + b[1]
			}
			var c = strings.Split(settup.TimeOff1, ":")
			var timeOff1 = settup.TimeOff1
			if len(c) > 1 {
				timeOff1 = b[0] + b[1]
			}
			var d = strings.Split(settup.TimeOff2, ":")
			var timeOff2 = settup.TimeOff2
			if len(d) > 1 {
				timeOff2 = b[0] + b[1]
			}

			if settup.TimeOn1 != s.Setup.TimeOldOn1 {
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0510,"+timeOn1+",2368,]")
			}
			set.TimeOldOn1 = settup.TimeOn1
			if settup.TimeOn2 != s.Setup.TimeOldOn2 {
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0518,"+timeOn2+",2368,]")
			}
			set.TimeOldOn2 = settup.TimeOn2
			if settup.TimeOff1 != s.Setup.TimeOldOff1 {

				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0511,"+timeOff1+",2368,]")
			}
			set.TimeOldOff1 = settup.TimeOff1
			if settup.TimeOff2 != s.Setup.TimeOldOff2 {
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0519,"+timeOff2+",2368,]")
			}
			set.TimeOldOff2 = settup.TimeOff2
			s.Setup = &set
		} else if s.Setup == nil {
			fmt.Println("Vô nil")
			var nhietDo, _ = strconv.Atoi(arr[6])
			var doAm, _ = strconv.Atoi(arr[7])

			var settup = setup.UpdateSetup{}
			settup.NhietDo = strconv.Itoa(nhietDo)
			settup.DoAm = strconv.Itoa(doAm)
			settup.TimeOn1 = arr[12]
			settup.TimeOff1 = arr[13]
			settup.TimeOn2 = arr[20]
			settup.TimeOff2 = arr[21]
			settup.ValueClient = data

			settup.Manual = arr[43]
			settup.ManualOn1 = arr[44]
			settup.ManualOn2 = arr[45]
			settup.IntUpClient = 0
			settup.IntUpServer = 0
			settup.TimeOldOn1 = arr[12]
			settup.TimeOldOff1 = arr[13]
			settup.TimeOldOff2 = arr[21]
			settup.TimeOldOn2 = arr[20]
			var set = setup.Setup{
				UpdateSetup: settup,
			}

			s.Setup = &set
			setup.InsertSetup(settup)

			s.SendDataString(ctx, "[,000000F1,0022,10,0001,0510,"+settup.TimeOn1+",2368,]")
			s.SendDataString(ctx, "[,000000F1,0022,10,0001,0511,"+settup.TimeOff1+",2368,]")
			s.SendDataString(ctx, "[,000000F1,0022,10,0001,0518,"+settup.TimeOn2+",2368,]")
			s.SendDataString(ctx, "[,000000F1,0022,10,0001,0519,"+settup.TimeOff2+",2368,]")
			return

		} else if s.Setup != nil && s.Setup.IntUpServer > s.Setup.IntUpClient {

			fmt.Println("Vô cài đặt tại server")
			s.Setup.ValueClient = data
			var a = strings.Split(s.Setup.TimeOn1, ":")
			var timeOn1 = s.Setup.TimeOn1
			if len(a) > 1 {
				timeOn1 = a[0] + a[1]
			}
			var b = strings.Split(s.Setup.TimeOn2, ":")
			var timeOn2 = s.Setup.TimeOn2
			if len(b) > 1 {
				timeOn2 = b[0] + b[1]
			}
			var c = strings.Split(s.Setup.TimeOff1, ":")
			var timeOff1 = s.Setup.TimeOff1
			if len(c) > 1 {
				timeOff1 = b[0] + b[1]
			}
			var d = strings.Split(s.Setup.TimeOff2, ":")
			var timeOff2 = s.Setup.TimeOff2
			if len(d) > 1 {
				timeOff2 = b[0] + b[1]
			}
			fmt.Printf("Time On1 Server", s.Setup.TimeOn1)
			fmt.Printf("Time On 1 client", arr[12])
			fmt.Printf("Time On2", arr[20])
			fmt.Printf("Time Off 2", arr[21])
			var isUp = true
			var timeOn1s = arr[12]
			var timeOff1s = arr[13]
			var timeOn2s = arr[20]
			var timeOff2s = arr[21]
			var nhietDo, _ = strconv.Atoi(arr[6])
			var doAm, _ = strconv.Atoi(arr[7])
			var nhietDoS = strconv.Itoa(nhietDo)
			var doAmS = strconv.Itoa(doAm)
			if nhietDoS != s.Setup.NhietDo {
				isUp = false
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,050A,00"+s.Setup.NhietDo+",2368,]")
			}
			if doAmS != s.Setup.DoAm {
				isUp = false
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,050B,00"+s.Setup.DoAm+",2368,]")
			}

			if s.Setup.TimeOn1 != timeOn1s {
				isUp = false
				s.Setup.TimeOldOn1 = s.Setup.TimeOn1
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0510,"+timeOn1+",2368,]")
			}
			if s.Setup.TimeOn2 != timeOn2s {
				isUp = false
				s.Setup.TimeOldOn2 = s.Setup.TimeOn2
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0518,"+timeOn2+",2368,]")
			}

			if s.Setup.TimeOff1 != timeOff1s {
				isUp = false
				s.Setup.TimeOldOff1 = s.Setup.TimeOff1
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0511,"+timeOff1+",2368,]")
			}
			if s.Setup.TimeOff2 != timeOff2s {
				isUp = false
				s.Setup.TimeOldOff2 = s.Setup.TimeOff2
				s.SendDataString(ctx, "[,000000F1,0022,10,0001,0519,"+timeOff2+",2368,]")
			}

			if isUp {
				s.Setup.IntUpClient = s.Setup.IntUpServer
			}
		}
	}

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
