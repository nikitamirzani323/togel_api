package controller

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/api_go/helpers"
	"bitbucket.org/isbtotogroup/api_go/model"
	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

type ClientToken struct {
	Token string `json:"token"`
}
type ClientInit struct {
	Client_Company string `json:"client_company"`
}
type ClientResult struct {
	Client_Company string `json:"client_company"`
	Pasaran_Code   string `json:"pasaran_code"`
}
type ClientResultAll struct {
	Client_Company string `json:"client_company"`
}
type ClientConfPasaran struct {
	Client_Company string `json:"client_company"`
	Pasaran_Code   string `json:"pasaran_code"`
	Permainan      string `json:"permainan"`
}
type ClientLimitPasaran struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Pasaran_Code    string `json:"pasaran_code"`
	Pasaran_Periode string `json:"pasaran_periode"`
	Permainan       string `json:"permainan"`
}
type ClientInvoicePasaran struct {
	Client_Idinvoice int    `json:"client_idinvoice"`
	Client_Username  string `json:"client_username"`
	Client_Company   string `json:"client_company"`
	Pasaran_Code     string `json:"pasaran_code"`
	Pasaran_Periode  string `json:"pasaran_periode"`
}
type ClientInvoicePasaranId struct {
	Client_Idinvoice int    `json:"client_idinvoice"`
	Client_Username  string `json:"client_username"`
	Client_Company   string `json:"client_company"`
	Permainan        string `json:"permainan"`
}
type ClientSlipPeriode struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Pasaran_Code    string `json:"pasaran_code"`
}
type ClientSlipPeriodeAll struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
}
type ClientSlipPeriodeDetail struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Idtrxkeluaran   string `json:"idtrxkeluaran"`
}
type ClientSaveTogel struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Idtrxkeluaran   string `json:"idtrxkeluaran"`
	Idcomppasaran   string `json:"idcomppasaran"`
	Devicemember    string `json:"devicemember"`
	Formipaddress   string `json:"formipaddress"`
	Timezone        string `json:"timezone"`
	Totalbayarbet   int    `json:"totalbayarbet"`
	List4d          string `json:"list4d"`
}
type parsingjson struct {
	Record []ytRecord `json:"record"`
}
type ytRecord struct {
	PasaranId      string `json:"pasaran_id"`
	PasaranTogel   string `json:"pasaran_togel"`
	PasaranPeriode string `json:"pasaran_periode"`
}
type responseredisfetch struct {
	Pasaran_id             string `json:"pasaran_id"`
	Pasaran_togel          string `json:"pasaran_togel"`
	Pasaran_periode        string `json:"pasaran_periode"`
	Pasaran_tglkeluaran    string `json:"pasaran_tglkeluaran"`
	Pasaran_marketclose    string `json:"pasaran_marketclose"`
	Pasaran_marketschedule string `json:"pasaran_marketschedule"`
	Pasaran_marketopen     string `json:"pasaran_marketopen"`
	Pasaran_status         string `json:"pasaran_status"`
}
type responseredis struct {
	No      int    `json:"no"`
	Date    string `json:"date"`
	Periode string `json:"periode"`
	Result  string `json:"result"`
}
type responseredisall struct {
	No          int    `json:"no"`
	Date        string `json:"date"`
	Pasaran     string `json:"pasaran"`
	Pasarancode string `json:"pasaran_code"`
	Periode     string `json:"periode"`
	Result      string `json:"result"`
}
type responseredisinit_432 struct {
	Min_bet           int     `json:"min_bet"`
	Max4d_bet         int     `json:"max4d_bet"`
	Max3d_bet         int     `json:"max3d_bet"`
	Max2d_bet         int     `json:"max2d_bet"`
	Max2dd_bet        int     `json:"max2dd_bet"`
	Max2dt_bet        int     `json:"max2dt_bet"`
	Disc4d_bet        float32 `json:"disc4d_bet"`
	Disc3d_bet        float32 `json:"disc3d_bet"`
	Disc2d_bet        float32 `json:"disc2d_bet"`
	Disc2dd_bet       float32 `json:"disc2dd_bet"`
	Disc2dt_bet       float32 `json:"disc2dt_bet"`
	Win4d_bet         int     `json:"win4d_bet"`
	Win3d_bet         int     `json:"win3d_bet"`
	Win2d_bet         int     `json:"win2d_bet"`
	Win2dd_bet        int     `json:"win2dd_bet"`
	Win2dt_bet        int     `json:"win2dt_bet"`
	Bbfs              int     `json:"bbfs"`
	Limitline_4d      int     `json:"limitline_4d"`
	Limitline_3d      int     `json:"limitline_3d"`
	Limitline_2d      int     `json:"limitline_2d"`
	Limitline_2dd     int     `json:"limitline_2dd"`
	Limitline_2dt     int     `json:"limitline_2dt"`
	Limittotal4d_bet  int     `json:"limittotal4d_bet"`
	Limittotal3d_bet  int     `json:"limittotal3d_bet"`
	Limittotal2d_bet  int     `json:"limittotal2d_bet"`
	Limittotal2dd_bet int     `json:"limittotal2dd_bet"`
	Limittotal2dt_bet int     `json:"limittotal2dt_bet"`
}
type responseredisinit_colok struct {
	Min_bet_colokbebas        int     `json:"min_bet_colokbebas"`
	Min_bet_colokmacau        int     `json:"min_bet_colokmacau"`
	Min_bet_coloknaga         int     `json:"min_bet_coloknaga"`
	Min_bet_colokjitu         int     `json:"min_bet_colokjitu"`
	Max_bet_colokbebas        int     `json:"max_bet_colokbebas"`
	Max_bet_colokmacau        int     `json:"max_bet_colokmacau"`
	Max_bet_coloknaga         int     `json:"max_bet_coloknaga"`
	Max_bet_colokjitu         int     `json:"max_bet_colokjitu"`
	Disc_bet_colokbebas       float32 `json:"disc_bet_colokbebas"`
	Disc_bet_colokmacau       float32 `json:"disc_bet_colokmacau"`
	Disc_bet_coloknaga        float32 `json:"disc_bet_coloknaga"`
	Disc_bet_colokjitu        float32 `json:"disc_bet_colokjitu"`
	Win_bet_colokbebas        float32 `json:"win_bet_colokbebas"`
	Win_bet_colokmacau        float32 `json:"win_bet_colokmacau"`
	Win3_bet_colokmacau       float32 `json:"win3_bet_colokmacau"`
	Win4_bet_colokmacau       float32 `json:"win4_bet_colokmacau"`
	Win4_bet_coloknaga        float32 `json:"win4_bet_coloknaga"`
	Winas_bet_colokjitu       float32 `json:"winas_bet_colokjitu"`
	Winekor_bet_colokjitu     float32 `json:"winekor_bet_colokjitu"`
	Winkepala_bet_colokjitu   float32 `json:"winkepala_bet_colokjitu"`
	Winkop_bet_colokjitu      float32 `json:"winkop_bet_colokjitu"`
	Limittotal_bet_colokbebas int     `json:"limittotal_bet_colokbebas"`
	Limittotal_bet_colokmacau int     `json:"limittotal_bet_colokmacau"`
	Limittotal_bet_coloknaga  int     `json:"limittotal_bet_coloknaga"`
	Limittotal_bet_colokjitu  int     `json:"limittotal_bet_colokjitu"`
}
type responseredisinit_5050 struct {
	Min_bet_5050umum                       int     `json:"min_bet_5050umum"`
	Min_bet_5050special                    int     `json:"min_bet_5050special"`
	Min_bet_5050kombinasi                  int     `json:"min_bet_5050kombinasi"`
	Max_bet_5050umum                       int     `json:"max_bet_5050umum"`
	Max_bet_5050special                    int     `json:"max_bet_5050special"`
	Max_bet_5050kombinasi                  int     `json:"max_bet_5050kombinasi"`
	Discbesar_bet_5050umum                 float32 `json:"discbesar_bet_5050umum"`
	Disckecil_bet_5050umum                 float32 `json:"disckecil_bet_5050umum"`
	Discganjil_bet_5050umum                float32 `json:"discganjil_bet_5050umum"`
	Discgenap_bet_5050umum                 float32 `json:"discgenap_bet_5050umum"`
	Disctengah_bet_5050umum                float32 `json:"disctengah_bet_5050umum"`
	Disctepi_bet_5050umum                  float32 `json:"disctepi_bet_5050umum"`
	Discasbesar_bet_5050special            float32 `json:"discasbesar_bet_5050special"`
	Discaskecil_bet_5050special            float32 `json:"discaskecil_bet_5050special"`
	Discasganjil_bet_5050special           float32 `json:"discasganjil_bet_5050special"`
	Discasgenap_bet_5050special            float32 `json:"discasgenap_bet_5050special"`
	Discekorbesar_bet_5050special          float32 `json:"discekorbesar_bet_5050special"`
	Discekorkecil_bet_5050special          float32 `json:"discekorkecil_bet_5050special"`
	Discekorganjil_bet_5050special         float32 `json:"discekorganjil_bet_5050special"`
	Discekorgenap_bet_5050special          float32 `json:"discekorgenap_bet_5050special"`
	Disckepalabesar_bet_5050special        float32 `json:"disckepalabesar_bet_5050special"`
	Disckepalakecil_bet_5050special        float32 `json:"disckepalakecil_bet_5050special"`
	Disckepalaganjil_bet_5050special       float32 `json:"disckepalaganjil_bet_5050special"`
	Disckepalagenap_bet_5050special        float32 `json:"disckepalagenap_bet_5050special"`
	Disckopbesar_bet_5050special           float32 `json:"disckopbesar_bet_5050special"`
	Disckopkecil_bet_5050special           float32 `json:"disckopkecil_bet_5050special"`
	Disckopganjil_bet_5050special          float32 `json:"disckopganjil_bet_5050special"`
	Disckopgenap_bet_5050special           float32 `json:"disckopgenap_bet_5050special"`
	Disc_belakangkembang_bet_5050kombinasi float32 `json:"disc_belakangkembang_bet_5050kombinasi"`
	Disc_belakangkembar_bet_5050kombinasi  float32 `json:"disc_belakangkembar_bet_5050kombinasi"`
	Disc_belakangkempis_bet_5050kombinasi  float32 `json:"disc_belakangkempis_bet_5050kombinasi"`
	Disc_belakangmono_bet_5050kombinasi    float32 `json:"disc_belakangmono_bet_5050kombinasi"`
	Disc_belakangstereo_bet_5050kombinasi  float32 `json:"disc_belakangstereo_bet_5050kombinasi"`
	Disc_depankembang_bet_5050kombinasi    float32 `json:"disc_depankembang_bet_5050kombinasi"`
	Disc_depankembar_bet_5050kombinasi     float32 `json:"disc_depankembar_bet_5050kombinasi"`
	Disc_depankempis_bet_5050kombinasi     float32 `json:"disc_depankempis_bet_5050kombinasi"`
	Disc_depanmono_bet_5050kombinasi       float32 `json:"disc_depanmono_bet_5050kombinasi"`
	Disc_depanstereo_bet_5050kombinasi     float32 `json:"disc_depanstereo_bet_5050kombinasi"`
	Disc_tengahkembang_bet_5050kombinasi   float32 `json:"disc_tengahkembang_bet_5050kombinasi"`
	Disc_tengahkembar_bet_5050kombinasi    float32 `json:"disc_tengahkembar_bet_5050kombinasi"`
	Disc_tengahkempis_bet_5050kombinasi    float32 `json:"disc_tengahkempis_bet_5050kombinasi"`
	Disc_tengahmono_bet_5050kombinasi      float32 `json:"disc_tengahmono_bet_5050kombinasi"`
	Disc_tengahstereo_bet_5050kombinasi    float32 `json:"disc_tengahstereo_bet_5050kombinasi"`
	Keibesar_bet_5050umum                  float32 `json:"keibesar_bet_5050umum"`
	Keikecil_bet_5050umum                  float32 `json:"keikecil_bet_5050umum"`
	Keiganjil_bet_5050umum                 float32 `json:"keiganjil_bet_5050umum"`
	Keigenap_bet_5050umum                  float32 `json:"keigenap_bet_5050umum"`
	Keitengah_bet_5050umum                 float32 `json:"keitengah_bet_5050umum"`
	Keitepi_bet_5050umum                   float32 `json:"keitepi_bet_5050umum"`
	Keiasbesar_bet_5050special             float32 `json:"keiasbesar_bet_5050special"`
	Keiaskecil_bet_5050special             float32 `json:"keiaskecil_bet_5050special"`
	Keiasganjil_bet_5050special            float32 `json:"keiasganjil_bet_5050special"`
	Keiasgenap_bet_5050special             float32 `json:"keiasgenap_bet_5050special"`
	Keiekorbesar_bet_5050special           float32 `json:"keiekorbesar_bet_5050special"`
	Keiekorkecil_bet_5050special           float32 `json:"keiekorkecil_bet_5050special"`
	Keiekorganjil_bet_5050special          float32 `json:"keiekorganjil_bet_5050special"`
	Keiekorgenap_bet_5050special           float32 `json:"keiekorgenap_bet_5050special"`
	Keikepalabesar_bet_5050special         float32 `json:"keikepalabesar_bet_5050special"`
	Keikepalakecil_bet_5050special         float32 `json:"keikepalakecil_bet_5050special"`
	Keikepalaganjil_bet_5050special        float32 `json:"keikepalaganjil_bet_5050special"`
	Keikepalagenap_bet_5050special         float32 `json:"keikepalagenap_bet_5050special"`
	Keikopbesar_bet_5050special            float32 `json:"keikopbesar_bet_5050special"`
	Keikopkecil_bet_5050special            float32 `json:"keikopkecil_bet_5050special"`
	Keikopganjil_bet_5050special           float32 `json:"keikopganjil_bet_5050special"`
	Keikopgenap_bet_5050special            float32 `json:"keikopgenap_bet_5050special"`
	Kei_belakangkembang_bet_5050kombinasi  float32 `json:"kei_belakangkembang_bet_5050kombinasi"`
	Kei_belakangkembar_bet_5050kombinasi   float32 `json:"kei_belakangkembar_bet_5050kombinasi"`
	Kei_belakangkempis_bet_5050kombinasi   float32 `json:"kei_belakangkempis_bet_5050kombinasi"`
	Kei_belakangmono_bet_5050kombinasi     float32 `json:"kei_belakangmono_bet_5050kombinasi"`
	Kei_belakangstereo_bet_5050kombinasi   float32 `json:"kei_belakangstereo_bet_5050kombinasi"`
	Kei_depankembang_bet_5050kombinasi     float32 `json:"kei_depankembang_bet_5050kombinasi"`
	Kei_depankembar_bet_5050kombinasi      float32 `json:"kei_depankembar_bet_5050kombinasi"`
	Kei_depankempis_bet_5050kombinasi      float32 `json:"kei_depankempis_bet_5050kombinasi"`
	Kei_depanmono_bet_5050kombinasi        float32 `json:"kei_depanmono_bet_5050kombinasi"`
	Kei_depanstereo_bet_5050kombinasi      float32 `json:"kei_depanstereo_bet_5050kombinasi"`
	Kei_tengahkembang_bet_5050kombinasi    float32 `json:"kei_tengahkembang_bet_5050kombinasi"`
	Kei_tengahkembar_bet_5050kombinasi     float32 `json:"kei_tengahkembar_bet_5050kombinasi"`
	Kei_tengahkempis_bet_5050kombinasi     float32 `json:"kei_tengahkempis_bet_5050kombinasi"`
	Kei_tengahmono_bet_5050kombinasi       float32 `json:"kei_tengahmono_bet_5050kombinasi"`
	Kei_tengahstereo_bet_5050kombinasi     float32 `json:"kei_tengahstereo_bet_5050kombinasi"`
	Limittotal_bet_5050umum                int     `json:"limittotal_bet_5050umum"`
	Limittotal_bet_5050special             int     `json:"limittotal_bet_5050special"`
	Limittotal_bet_5050kombinasi           int     `json:"limittotal_bet_5050kombinasi"`
}
type responseredisinit_macaukombinasi struct {
	Min_bet     int     `json:"min_bet"`
	Max_bet     int     `json:"max_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Win_bet     float32 `json:"win_bet"`
	Limit_total int     `json:"limit_total"`
}
type responseredisinit_dasar struct {
	Min_bet         int     `json:"min_bet"`
	Max_bet         int     `json:"max_bet"`
	Disc_besar_bet  float32 `json:"disc_besar_bet"`
	Disc_kecil_bet  float32 `json:"disc_kecil_bet"`
	Disc_ganjil_bet float32 `json:"disc_ganjil_bet"`
	Disc_genap_bet  float32 `json:"disc_genap_bet"`
	Kei_besar_bet   float32 `json:"kei_besar_bet"`
	Kei_kecil_bet   float32 `json:"kei_kecil_bet"`
	Kei_ganjil_bet  float32 `json:"kei_ganjil_bet"`
	Kei_genap_bet   float32 `json:"kei_genap_bet"`
	Limit_total     int     `json:"limit_total"`
}
type responseredisinit_shio struct {
	Min_bet     int     `json:"min_bet"`
	Max_bet     int     `json:"max_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Win_bet     float32 `json:"win_bet"`
	Limit_total int     `json:"limit_total"`
}
type responseinvoiceall struct {
	Tglkeluaran     string `json:"tglkeluaran"`
	Idinvoice       string `json:"idinvoice"`
	Pasaran         string `json:"pasaran"`
	Periode         string `json:"periode"`
	Totalbet        int    `json:"totalbet"`
	Totalbayar      int    `json:"totalbayar"`
	Totalwin        int    `json:"totalwin"`
	Totallose       int    `json:"totallose"`
	Status          string `json:"status"`
	Color_lost      string `json:"color_lost"`
	Background      string `json:"background"`
	Color_totallose string `json:"color_totallose"`
}
type responseinvoiceidpermainan struct {
	No        int     `json:"no"`
	Status    string  `json:"status"`
	Permainan string  `json:"permainan"`
	Nomor     string  `json:"nomor"`
	Bet       int     `json:"bet"`
	Diskon    float32 `json:"diskon"`
	Kei       float32 `json:"kei"`
	Bayar     int     `json:"bayar"`
	Win       int     `json:"win"`
}
type responseredislistinvoicebet struct {
	Tanggal   string  `json:"tanggal"`
	Permainan string  `json:"permainan"`
	Periode   string  `json:"periode"`
	Nomor     string  `json:"nomor"`
	Bet       int     `json:"bet"`
	Diskon    float32 `json:"diskon"`
	Kei       float32 `json:"kei"`
	Bayar     int     `json:"bayar"`
	Win       int     `json:"win"`
	Menang    int     `json:"menang"`
}
type settingcustom struct {
	Status int         `json:"status"`
	Record interface{} `json:"record"`
}
type settingcustom2 struct {
	StartMaintenance string `json:"maintenance_start"`
	EndMaintenance   string `json:"maintenance_end"`
}

var ctx = context.Background()

func Fetch_token(c *fiber.Ctx) error {
	client := new(ClientToken)

	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	field_redis := "LISTSETTING_MASTER"
	tglnow, _ := goment.New()
	website_status := "ONLINE"
	website_message := ""
	tglskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
	jamstart := ""
	jamend := ""
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, err := jsonparser.Get(jsonredis, "record")
	log.Println(err)
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		maintenance_start, _ := jsonparser.GetString(value, "maintenance_start")
		maintenance_end, _ := jsonparser.GetString(value, "maintenance_end")
		jamstart = tglnow.Format("YYYY-MM-DD") + " " + maintenance_start
		jamend = tglnow.Format("YYYY-MM-DD") + " " + maintenance_end

	})
	if tglskrg >= jamstart && tglskrg <= jamend {
		website_status = "OFFLINE"
		website_message = "MAINTENANCE START : " + jamstart + ", FINISH : " + jamend
	}
	member_username := ""
	member_company := ""
	member_saldo := 0
	switch client.Token {
	case "qC5YmBvXzabGp34jJlKvnC6wCrr3pLCwBzsLoSzl4k=":
		member_username = "developer"
		member_company = "MMD"
		member_saldo = 5000000
	case "qwertyuiop":
		member_username = "Jhon Wick"
		member_company = "MMD"
		member_saldo = 1000000
	case "asdfghjkl":
		member_username = "Edo Febrian"
		member_company = "MMD"
		member_saldo = 200000
	case "1234567890":
		member_username = "developerisb"
		member_company = "ISB"
		member_saldo = 1000000
	case "0987654321`":
		member_username = "developernuke"
		member_company = "NUKE"
		member_saldo = 2000000
	}
	if !flag {
		result, err := model.Fetch_Setting()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		var obj settingcustom
		var obj2 settingcustom2
		var arraobj2 []settingcustom2

		dataresult, _ := json.Marshal(result)
		status_rd, _ := jsonparser.GetInt(dataresult, "status")
		record_RD, _, _, _ := jsonparser.Get(dataresult, "record")

		log.Println("INIT MYSQL")
		jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			maintenance_start, _ := jsonparser.GetString(value, "maintenance_start")
			maintenance_end, _ := jsonparser.GetString(value, "maintenance_end")
			jamstart = tglnow.Format("YYYY-MM-DD") + " " + maintenance_start
			jamend = tglnow.Format("YYYY-MM-DD") + " " + maintenance_end
			obj2.StartMaintenance = maintenance_start
			obj2.EndMaintenance = maintenance_end
			arraobj2 = append(arraobj2, obj2)
		})
		obj.Status = int(status_rd)
		obj.Record = arraobj2
		helpers.SetRedis(field_redis, obj, 0)

		if tglskrg >= jamstart && tglskrg <= jamend {
			website_status = "OFFLINE"
			website_message = "START : " + jamstart + ", FINISH : " + jamend
		}
		return c.JSON(fiber.Map{
			"status":          fiber.StatusOK,
			"token":           client.Token,
			"website_status":  website_status,
			"website_message": website_message,
			"member_username": member_username,
			"member_company":  member_company,
			"member_credit":   member_saldo,
		})
	} else {
		log.Println("INIT cache")
		return c.JSON(fiber.Map{
			"status":          fiber.StatusOK,
			"token":           client.Token,
			"website_status":  website_status,
			"website_message": website_message,
			"member_username": member_username,
			"member_company":  member_company,
			"member_credit":   member_saldo,
		})
	}
}
func FetchAll_pasaran(c *fiber.Ctx) error {
	client := new(ClientInit)

	if err := c.BodyParser(client); err != nil {
		return err
	}
	field_redis := "listpasaran_" + strings.ToLower(client.Client_Company)
	render_page := time.Now()
	tglnow, _ := goment.New()
	var obj responseredisfetch
	var arraobj []responseredisfetch
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	statuspasaran := "ONLINE"
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_id, _ := jsonparser.GetString(value, "pasaran_id")
		pasaran_togel, _ := jsonparser.GetString(value, "pasaran_togel")
		pasaran_periode, _ := jsonparser.GetString(value, "pasaran_periode")
		pasaran_tglkeluaran, _ := jsonparser.GetString(value, "pasaran_tglkeluaran")
		pasaran_marketclose, _ := jsonparser.GetString(value, "pasaran_marketclose")
		pasaran_marketschedule, _ := jsonparser.GetString(value, "pasaran_marketschedule")
		pasaran_marketopen, _ := jsonparser.GetString(value, "pasaran_marketopen")
		pasaran_jamtutup, _ := jsonparser.GetString(value, "pasaran_jamtutup")
		pasaran_jamopen, _ := jsonparser.GetString(value, "pasaran_jamopen")
		pasaran_hari, _ := jsonparser.GetString(value, "pasaran_hari")

		tgltutup, _ := goment.New(pasaran_marketclose)
		jamtutup2 := tgltutup.Format("YYYY-MM-DD") + " " + pasaran_jamtutup
		// tglopen, _ := goment.New(pasaran_marketopen)

		taiskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
		jamtutup := tglnow.Format("YYYY-MM-DD") + " " + pasaran_jamtutup
		jamopen := tglnow.Format("YYYY-MM-DD") + " " + pasaran_jamopen

		if taiskrg >= jamtutup2 {
			statuspasaran = "OFFLINE"
		} else {
			if pasaran_hari == "ONLINE" {
				if taiskrg >= jamtutup && taiskrg <= jamopen {
					statuspasaran = "OFFLINE"
				} else {
					statuspasaran = "ONLINE"
				}
			} else {
				statuspasaran = "ONLINE"
			}
		}

		obj.Pasaran_id = pasaran_id
		obj.Pasaran_togel = pasaran_togel
		obj.Pasaran_periode = pasaran_periode
		obj.Pasaran_tglkeluaran = pasaran_tglkeluaran
		obj.Pasaran_marketclose = pasaran_marketclose
		obj.Pasaran_marketschedule = pasaran_marketschedule
		obj.Pasaran_marketopen = pasaran_marketopen
		obj.Pasaran_status = statuspasaran
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := model.FetchAll_MclientPasaran(client.Client_Company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 0)
		log.Println("PASARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func FetchAll_resultbypasaran(c *fiber.Ctx) error {
	client := new(ClientResult)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}

	field_redis := "listresult_" + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Pasaran_Code)
	var obj responseredis
	var arraobj []responseredis
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")

	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		no_RD, _ := jsonparser.GetInt(value, "no")
		date_RD, _ := jsonparser.GetString(value, "date")
		periode_RD, _ := jsonparser.GetString(value, "periode")
		result_RD, _ := jsonparser.GetString(value, "result")

		obj.No = int(no_RD)
		obj.Date = date_RD
		obj.Periode = periode_RD
		obj.Result = result_RD
		arraobj = append(arraobj, obj)

	})

	if !flag {
		result, err := model.FetchAll_MclientPasaranResult(client.Client_Company, client.Pasaran_Code)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}

		log.Println("RESulT mysql")
		if result.Status == 200 {
			helpers.SetRedis(field_redis, result, 0)
		}
		return c.JSON(result)
	} else {
		log.Println("RESulT cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})

	}
}
func FetchAll_result(c *fiber.Ctx) error {
	client := new(ClientResultAll)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}

	field_redis := "listresult_" + strings.ToLower(client.Client_Company)
	var obj responseredisall
	var arraobj []responseredisall
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")

	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		no_RD, _ := jsonparser.GetInt(value, "no")
		date_RD, _ := jsonparser.GetString(value, "date")
		pasaran_RD, _ := jsonparser.GetString(value, "pasaran")
		pasaran_code_RD, _ := jsonparser.GetString(value, "pasaran_code")
		periode_RD, _ := jsonparser.GetString(value, "periode")
		result_RD, _ := jsonparser.GetString(value, "result")

		obj.No = int(no_RD)
		obj.Date = date_RD
		obj.Pasaran = pasaran_RD
		obj.Pasarancode = pasaran_code_RD
		obj.Periode = periode_RD
		obj.Result = result_RD
		arraobj = append(arraobj, obj)

	})

	if !flag {
		result, err := model.FetchAll_MclientPasaranResultAll(client.Client_Company)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}

		log.Println("PASARAN mysql")
		if result.Status == 200 {
			helpers.SetRedis(field_redis, result, 0)
		}
		return c.JSON(result)
	} else {
		log.Println("cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})

	}
}
func Fetch_CheckPasaran(c *fiber.Ctx) error {
	client := new(ClientResult)

	if err := c.BodyParser(client); err != nil {
		return err
	}
	result, err := model.CheckPasaran(client.Client_Company, client.Pasaran_Code)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func Fetch_InitPasaran(c *fiber.Ctx) error {
	client := new(ClientConfPasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}
	field_redis := "config_" + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Pasaran_Code) + "_" + strings.ToLower(client.Permainan)
	var obj_432 responseredisinit_432
	var arraobj_432 []responseredisinit_432
	var obj_colok responseredisinit_colok
	var arraobj_colok []responseredisinit_colok
	var obj_5050 responseredisinit_5050
	var arraobj_5050 []responseredisinit_5050
	var obj_macaukombinasi responseredisinit_macaukombinasi
	var arraobj_macaukombinasi []responseredisinit_macaukombinasi
	var obj_dasar responseredisinit_dasar
	var arraobj_dasar []responseredisinit_dasar
	var obj_shio responseredisinit_shio
	var arraobj_shio []responseredisinit_shio
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch client.Permainan {
		case "4-3-2":
			minbet_RD, _ := jsonparser.GetInt(value, "min_bet")
			maxbet4d_RD, _ := jsonparser.GetInt(value, "max4d_bet")
			maxbet3d_RD, _ := jsonparser.GetInt(value, "max3d_bet")
			maxbet2d_RD, _ := jsonparser.GetInt(value, "max2d_bet")
			maxbet2dd_RD, _ := jsonparser.GetInt(value, "max2dd_bet")
			maxbet2dt_RD, _ := jsonparser.GetInt(value, "max2dt_bet")
			disc4d_RD, _ := jsonparser.GetFloat(value, "disc4d_bet")
			disc3d_RD, _ := jsonparser.GetFloat(value, "disc3d_bet")
			disc2d_RD, _ := jsonparser.GetFloat(value, "disc2d_bet")
			disc2dd_RD, _ := jsonparser.GetFloat(value, "disc2dd_bet")
			disc2dt_RD, _ := jsonparser.GetFloat(value, "disc2dt_bet")
			win4d_RD, _ := jsonparser.GetInt(value, "win4d_bet")
			win3d_RD, _ := jsonparser.GetInt(value, "win3d_bet")
			win2d_RD, _ := jsonparser.GetInt(value, "win2d_bet")
			win2dd_RD, _ := jsonparser.GetInt(value, "win2dd_bet")
			win2dt_RD, _ := jsonparser.GetInt(value, "win2dt_bet")
			limitline_4d_RD, _ := jsonparser.GetInt(value, "limitline_4d")
			limitline_3d_RD, _ := jsonparser.GetInt(value, "limitline_3d")
			limitline_2d_RD, _ := jsonparser.GetInt(value, "limitline_2d")
			limitline_2dd_RD, _ := jsonparser.GetInt(value, "limitline_2dd")
			limitline_2dt_RD, _ := jsonparser.GetInt(value, "limitline_2dt")
			limittotal4d_RD, _ := jsonparser.GetInt(value, "limittotal4d_bet")
			limittotal3d_RD, _ := jsonparser.GetInt(value, "limittotal3d_bet")
			limittotal2d_RD, _ := jsonparser.GetInt(value, "limittotal2d_bet")
			limittotal2dd_RD, _ := jsonparser.GetInt(value, "limittotal2dd_bet")
			limittotal2dt_RD, _ := jsonparser.GetInt(value, "limittotal2dt_bet")
			bbfs_RD, _ := jsonparser.GetInt(value, "bbfs")

			obj_432.Min_bet = int(minbet_RD)
			obj_432.Max4d_bet = int(maxbet4d_RD)
			obj_432.Max3d_bet = int(maxbet3d_RD)
			obj_432.Max2d_bet = int(maxbet2d_RD)
			obj_432.Max2dd_bet = int(maxbet2dd_RD)
			obj_432.Max2dt_bet = int(maxbet2dt_RD)
			obj_432.Disc4d_bet = float32(disc4d_RD)
			obj_432.Disc3d_bet = float32(disc3d_RD)
			obj_432.Disc2d_bet = float32(disc2d_RD)
			obj_432.Disc2dd_bet = float32(disc2dd_RD)
			obj_432.Disc2dt_bet = float32(disc2dt_RD)
			obj_432.Win4d_bet = int(win4d_RD)
			obj_432.Win3d_bet = int(win3d_RD)
			obj_432.Win2d_bet = int(win2d_RD)
			obj_432.Win2dd_bet = int(win2dd_RD)
			obj_432.Win2dt_bet = int(win2dt_RD)
			obj_432.Limitline_4d = int(limitline_4d_RD)
			obj_432.Limitline_3d = int(limitline_3d_RD)
			obj_432.Limitline_2d = int(limitline_2d_RD)
			obj_432.Limitline_2dd = int(limitline_2dd_RD)
			obj_432.Limitline_2dt = int(limitline_2dt_RD)
			obj_432.Limittotal4d_bet = int(limittotal4d_RD)
			obj_432.Limittotal3d_bet = int(limittotal3d_RD)
			obj_432.Limittotal2d_bet = int(limittotal2d_RD)
			obj_432.Limittotal2dd_bet = int(limittotal2dd_RD)
			obj_432.Limittotal2dt_bet = int(limittotal2dt_RD)
			obj_432.Bbfs = int(bbfs_RD)

			arraobj_432 = append(arraobj_432, obj_432)
		case "colok":
			Min_bet_colokbebas_RD, _ := jsonparser.GetInt(value, "min_bet_colokbebas")
			Min_bet_colokmacau_RD, _ := jsonparser.GetInt(value, "min_bet_colokmacau")
			Min_bet_coloknaga_RD, _ := jsonparser.GetInt(value, "min_bet_coloknaga")
			Min_bet_colokjitu_RD, _ := jsonparser.GetInt(value, "min_bet_colokjitu")
			Max_bet_colokbebas_RD, _ := jsonparser.GetInt(value, "max_bet_colokbebas")
			Max_bet_colokmacau_RD, _ := jsonparser.GetInt(value, "max_bet_colokmacau")
			Max_bet_coloknaga_RD, _ := jsonparser.GetInt(value, "max_bet_coloknaga")
			Max_bet_colokjitu_RD, _ := jsonparser.GetInt(value, "max_bet_colokjitu")
			Disc_bet_colokbebas_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokbebas")
			Disc_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokmacau")
			Disc_bet_coloknaga_RD, _ := jsonparser.GetFloat(value, "disc_bet_coloknaga")
			Disc_bet_colokjitu_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokjitu")
			Win_bet_colokbebas_RD, _ := jsonparser.GetFloat(value, "win_bet_colokbebas")
			Win_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "win_bet_colokmacau")
			Win3_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "win3_bet_colokmacau")
			Win4_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "win4_bet_colokmacau")
			Win4_bet_coloknaga_RD, _ := jsonparser.GetFloat(value, "win4_bet_coloknaga")
			Winas_bet_colokjitu, _ := jsonparser.GetFloat(value, "winas_bet_colokjitu")
			Winekor_bet_colokjitu, _ := jsonparser.GetFloat(value, "winekor_bet_colokjitu")
			Winkepala_bet_colokjitu, _ := jsonparser.GetFloat(value, "winkepala_bet_colokjitu")
			Winkop_bet_colokjitu, _ := jsonparser.GetFloat(value, "winkop_bet_colokjitu")
			Limittotal_bet_colokbebas, _ := jsonparser.GetInt(value, "limittotal_bet_colokbebas")
			Limittotal_bet_colokmacau, _ := jsonparser.GetInt(value, "limittotal_bet_colokmacau")
			Limittotal_bet_coloknaga, _ := jsonparser.GetInt(value, "limittotal_bet_coloknaga")
			Limittotal_bet_colokjitu, _ := jsonparser.GetInt(value, "limittotal_bet_colokjitu")

			obj_colok.Min_bet_colokbebas = int(Min_bet_colokbebas_RD)
			obj_colok.Min_bet_colokmacau = int(Min_bet_colokmacau_RD)
			obj_colok.Min_bet_coloknaga = int(Min_bet_coloknaga_RD)
			obj_colok.Min_bet_colokjitu = int(Min_bet_colokjitu_RD)
			obj_colok.Max_bet_colokbebas = int(Max_bet_colokbebas_RD)
			obj_colok.Max_bet_colokmacau = int(Max_bet_colokmacau_RD)
			obj_colok.Max_bet_coloknaga = int(Max_bet_coloknaga_RD)
			obj_colok.Max_bet_colokjitu = int(Max_bet_colokjitu_RD)
			obj_colok.Disc_bet_colokbebas = float32(Disc_bet_colokbebas_RD)
			obj_colok.Disc_bet_colokmacau = float32(Disc_bet_colokmacau_RD)
			obj_colok.Disc_bet_coloknaga = float32(Disc_bet_coloknaga_RD)
			obj_colok.Disc_bet_colokjitu = float32(Disc_bet_colokjitu_RD)
			obj_colok.Win_bet_colokbebas = float32(Win_bet_colokbebas_RD)
			obj_colok.Win_bet_colokmacau = float32(Win_bet_colokmacau_RD)
			obj_colok.Win3_bet_colokmacau = float32(Win3_bet_colokmacau_RD)
			obj_colok.Win4_bet_colokmacau = float32(Win4_bet_colokmacau_RD)
			obj_colok.Win4_bet_coloknaga = float32(Win4_bet_coloknaga_RD)
			obj_colok.Winas_bet_colokjitu = float32(Winas_bet_colokjitu)
			obj_colok.Winekor_bet_colokjitu = float32(Winekor_bet_colokjitu)
			obj_colok.Winkepala_bet_colokjitu = float32(Winkepala_bet_colokjitu)
			obj_colok.Winkop_bet_colokjitu = float32(Winkop_bet_colokjitu)
			obj_colok.Limittotal_bet_colokbebas = int(Limittotal_bet_colokbebas)
			obj_colok.Limittotal_bet_colokmacau = int(Limittotal_bet_colokmacau)
			obj_colok.Limittotal_bet_coloknaga = int(Limittotal_bet_coloknaga)
			obj_colok.Limittotal_bet_colokjitu = int(Limittotal_bet_colokjitu)

			arraobj_colok = append(arraobj_colok, obj_colok)
		case "5050":
			Min_bet_5050umum_RD, _ := jsonparser.GetInt(value, "min_bet_5050umum")
			Min_bet_5050special_RD, _ := jsonparser.GetInt(value, "min_bet_5050special")
			Min_bet_5050kombinasi_RD, _ := jsonparser.GetInt(value, "min_bet_5050kombinasi")
			Max_bet_5050umum_RD, _ := jsonparser.GetInt(value, "max_bet_5050umum")
			Max_bet_5050special_RD, _ := jsonparser.GetInt(value, "max_bet_5050special")
			Max_bet_5050kombinasi_RD, _ := jsonparser.GetInt(value, "max_bet_5050kombinasi")
			Discbesar_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "discbesar_bet_5050umum")
			Disckecil_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "disckecil_bet_5050umum")
			Discganjil_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "discganjil_bet_5050umum")
			Discgenap_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "discgenap_bet_5050umum")
			Disctengah_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "disctengah_bet_5050umum")
			Disctepi_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "disctepi_bet_5050umum")
			Discasbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discasbesar_bet_5050special")
			Discaskecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discaskecil_bet_5050special")
			Discasganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discasganjil_bet_5050special")
			Discasgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discasgenap_bet_5050special")
			Discekorbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discekorbesar_bet_5050special")
			Discekorkecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discekorkecil_bet_5050special")
			Discekorganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discekorganjil_bet_5050special")
			Discekorgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "discekorgenap_bet_5050special")
			Disckepalabesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckepalabesar_bet_5050special")
			Disckepalakecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckepalakecil_bet_5050special")
			Disckepalaganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckepalaganjil_bet_5050special")
			Disckepalagenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckepalagenap_bet_5050special")
			Disckopbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckopbesar_bet_5050special")
			Disckopkecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckopkecil_bet_5050special")
			Disckopganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckopganjil_bet_5050special")
			Disckopgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "disckopgenap_bet_5050special")
			Disc_belakangkembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_belakangkembang_bet_5050kombinasi")
			Disc_belakangkembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_belakangkembar_bet_5050kombinasi")
			Disc_belakangkempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_belakangkempis_bet_5050kombinasi")
			Disc_belakangmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_belakangmono_bet_5050kombinasi")
			Disc_belakangstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_belakangstereo_bet_5050kombinasi")
			Disc_depankembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_depankembang_bet_5050kombinasi")
			Disc_depankembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_depankembar_bet_5050kombinasi")
			Disc_depankempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_depankempis_bet_5050kombinasi")
			Disc_depanmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_depanmono_bet_5050kombinasi")
			Disc_depanstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_depanstereo_bet_5050kombinasi")
			Disc_tengahkembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_tengahkembang_bet_5050kombinasi")
			Disc_tengahkembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_tengahkembar_bet_5050kombinasi")
			Disc_tengahkempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_tengahkempis_bet_5050kombinasi")
			Disc_tengahmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_tengahmono_bet_5050kombinasi")
			Disc_tengahstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "disc_tengahstereo_bet_5050kombinasi")
			Keibesar_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keibesar_bet_5050umum")
			Keikecil_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keikecil_bet_5050umum")
			Keiganjil_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keiganjil_bet_5050umum")
			Keigenap_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keigenap_bet_5050umum")
			Keitengah_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keitengah_bet_5050umum")
			Keitepi_bet_5050umum_RD, _ := jsonparser.GetFloat(value, "keitepi_bet_5050umum")
			Keiasbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiasbesar_bet_5050special")
			Keiaskecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiaskecil_bet_5050special")
			Keiasganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiasganjil_bet_5050special")
			Keiasgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiasgenap_bet_5050special")
			Keiekorbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiekorbesar_bet_5050special")
			Keiekorkecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiekorkecil_bet_5050special")
			Keiekorganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiekorganjil_bet_5050special")
			Keiekorgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keiekorgenap_bet_5050special")
			Keikepalabesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikepalabesar_bet_5050special")
			Keikepalakecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikepalakecil_bet_5050special")
			Keikepalaganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikepalaganjil_bet_5050special")
			Keikepalagenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikepalagenap_bet_5050special")
			Keikopbesar_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikopbesar_bet_5050special")
			Keikopkecil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikopkecil_bet_5050special")
			Keikopganjil_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikopganjil_bet_5050special")
			Keikopgenap_bet_5050special_RD, _ := jsonparser.GetFloat(value, "keikopgenap_bet_5050special")
			Kei_belakangkembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_belakangkembang_bet_5050kombinasi")
			Kei_belakangkembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_belakangkembar_bet_5050kombinasi")
			Kei_belakangkempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_belakangkempis_bet_5050kombinasi")
			Kei_belakangmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_belakangmono_bet_5050kombinasi")
			Kei_belakangstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_belakangstereo_bet_5050kombinasi")
			Kei_depankembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_depankembang_bet_5050kombinasi")
			Kei_depankembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_depankembar_bet_5050kombinasi")
			Kei_depankempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_depankempis_bet_5050kombinasi")
			Kei_depanmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_depanmono_bet_5050kombinasi")
			Kei_depanstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_depanstereo_bet_5050kombinasi")
			Kei_tengahkembang_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_tengahkembang_bet_5050kombinasi")
			Kei_tengahkembar_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_tengahkembar_bet_5050kombinasi")
			Kei_tengahkempis_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_tengahkempis_bet_5050kombinasi")
			Kei_tengahmono_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_tengahmono_bet_5050kombinasi")
			Kei_tengahstereo_bet_5050kombinasi_RD, _ := jsonparser.GetFloat(value, "kei_tengahstereo_bet_5050kombinasi")
			Limittotal_bet_5050umum_RD, _ := jsonparser.GetInt(value, "limittotal_bet_5050umum")
			Limittotal_bet_5050special_RD, _ := jsonparser.GetInt(value, "limittotal_bet_5050special")
			Limittotal_bet_5050kombinasi_RD, _ := jsonparser.GetInt(value, "limittotal_bet_5050kombinasi")
			obj_5050.Min_bet_5050umum = int(Min_bet_5050umum_RD)
			obj_5050.Min_bet_5050special = int(Min_bet_5050special_RD)
			obj_5050.Min_bet_5050kombinasi = int(Min_bet_5050kombinasi_RD)
			obj_5050.Max_bet_5050umum = int(Max_bet_5050umum_RD)
			obj_5050.Max_bet_5050special = int(Max_bet_5050special_RD)
			obj_5050.Max_bet_5050kombinasi = int(Max_bet_5050kombinasi_RD)
			obj_5050.Discbesar_bet_5050umum = float32(Discbesar_bet_5050umum_RD)
			obj_5050.Disckecil_bet_5050umum = float32(Disckecil_bet_5050umum_RD)
			obj_5050.Discganjil_bet_5050umum = float32(Discganjil_bet_5050umum_RD)
			obj_5050.Discgenap_bet_5050umum = float32(Discgenap_bet_5050umum_RD)
			obj_5050.Disctengah_bet_5050umum = float32(Disctengah_bet_5050umum_RD)
			obj_5050.Disctepi_bet_5050umum = float32(Disctepi_bet_5050umum_RD)
			obj_5050.Discasbesar_bet_5050special = float32(Discasbesar_bet_5050special_RD)
			obj_5050.Discaskecil_bet_5050special = float32(Discaskecil_bet_5050special_RD)
			obj_5050.Discasganjil_bet_5050special = float32(Discasganjil_bet_5050special_RD)
			obj_5050.Discasgenap_bet_5050special = float32(Discasgenap_bet_5050special_RD)
			obj_5050.Discekorbesar_bet_5050special = float32(Discekorbesar_bet_5050special_RD)
			obj_5050.Discekorkecil_bet_5050special = float32(Discekorkecil_bet_5050special_RD)
			obj_5050.Discekorganjil_bet_5050special = float32(Discekorganjil_bet_5050special_RD)
			obj_5050.Discekorgenap_bet_5050special = float32(Discekorgenap_bet_5050special_RD)
			obj_5050.Disckepalabesar_bet_5050special = float32(Disckepalabesar_bet_5050special_RD)
			obj_5050.Disckepalakecil_bet_5050special = float32(Disckepalakecil_bet_5050special_RD)
			obj_5050.Disckepalaganjil_bet_5050special = float32(Disckepalaganjil_bet_5050special_RD)
			obj_5050.Disckepalagenap_bet_5050special = float32(Disckepalagenap_bet_5050special_RD)
			obj_5050.Disckopbesar_bet_5050special = float32(Disckopbesar_bet_5050special_RD)
			obj_5050.Disckopkecil_bet_5050special = float32(Disckopkecil_bet_5050special_RD)
			obj_5050.Disckopganjil_bet_5050special = float32(Disckopganjil_bet_5050special_RD)
			obj_5050.Disckopgenap_bet_5050special = float32(Disckopgenap_bet_5050special_RD)
			obj_5050.Disc_belakangkembang_bet_5050kombinasi = float32(Disc_belakangkembang_bet_5050kombinasi_RD)
			obj_5050.Disc_belakangkembar_bet_5050kombinasi = float32(Disc_belakangkembar_bet_5050kombinasi_RD)
			obj_5050.Disc_belakangkempis_bet_5050kombinasi = float32(Disc_belakangkempis_bet_5050kombinasi_RD)
			obj_5050.Disc_belakangmono_bet_5050kombinasi = float32(Disc_belakangmono_bet_5050kombinasi_RD)
			obj_5050.Disc_belakangstereo_bet_5050kombinasi = float32(Disc_belakangstereo_bet_5050kombinasi_RD)
			obj_5050.Disc_depankembang_bet_5050kombinasi = float32(Disc_depankembang_bet_5050kombinasi_RD)
			obj_5050.Disc_depankembar_bet_5050kombinasi = float32(Disc_depankembar_bet_5050kombinasi_RD)
			obj_5050.Disc_depankempis_bet_5050kombinasi = float32(Disc_depankempis_bet_5050kombinasi_RD)
			obj_5050.Disc_depanmono_bet_5050kombinasi = float32(Disc_depanmono_bet_5050kombinasi_RD)
			obj_5050.Disc_depanstereo_bet_5050kombinasi = float32(Disc_depanstereo_bet_5050kombinasi_RD)
			obj_5050.Disc_tengahkembang_bet_5050kombinasi = float32(Disc_tengahkembang_bet_5050kombinasi_RD)
			obj_5050.Disc_tengahkembar_bet_5050kombinasi = float32(Disc_tengahkembar_bet_5050kombinasi_RD)
			obj_5050.Disc_tengahkempis_bet_5050kombinasi = float32(Disc_tengahkempis_bet_5050kombinasi_RD)
			obj_5050.Disc_tengahmono_bet_5050kombinasi = float32(Disc_tengahmono_bet_5050kombinasi_RD)
			obj_5050.Disc_tengahstereo_bet_5050kombinasi = float32(Disc_tengahstereo_bet_5050kombinasi_RD)
			obj_5050.Keibesar_bet_5050umum = float32(Keibesar_bet_5050umum_RD)
			obj_5050.Keikecil_bet_5050umum = float32(Keikecil_bet_5050umum_RD)
			obj_5050.Keiganjil_bet_5050umum = float32(Keiganjil_bet_5050umum_RD)
			obj_5050.Keigenap_bet_5050umum = float32(Keigenap_bet_5050umum_RD)
			obj_5050.Keitengah_bet_5050umum = float32(Keitengah_bet_5050umum_RD)
			obj_5050.Keitepi_bet_5050umum = float32(Keitepi_bet_5050umum_RD)
			obj_5050.Keiasbesar_bet_5050special = float32(Keiasbesar_bet_5050special_RD)
			obj_5050.Keiaskecil_bet_5050special = float32(Keiaskecil_bet_5050special_RD)
			obj_5050.Keiasganjil_bet_5050special = float32(Keiasganjil_bet_5050special_RD)
			obj_5050.Keiasgenap_bet_5050special = float32(Keiasgenap_bet_5050special_RD)
			obj_5050.Keiekorbesar_bet_5050special = float32(Keiekorbesar_bet_5050special_RD)
			obj_5050.Keiekorkecil_bet_5050special = float32(Keiekorkecil_bet_5050special_RD)
			obj_5050.Keiekorganjil_bet_5050special = float32(Keiekorganjil_bet_5050special_RD)
			obj_5050.Keiekorgenap_bet_5050special = float32(Keiekorgenap_bet_5050special_RD)
			obj_5050.Keikepalabesar_bet_5050special = float32(Keikepalabesar_bet_5050special_RD)
			obj_5050.Keikepalakecil_bet_5050special = float32(Keikepalakecil_bet_5050special_RD)
			obj_5050.Keikepalaganjil_bet_5050special = float32(Keikepalaganjil_bet_5050special_RD)
			obj_5050.Keikepalagenap_bet_5050special = float32(Keikepalagenap_bet_5050special_RD)
			obj_5050.Keikopbesar_bet_5050special = float32(Keikopbesar_bet_5050special_RD)
			obj_5050.Keikopkecil_bet_5050special = float32(Keikopkecil_bet_5050special_RD)
			obj_5050.Keikopganjil_bet_5050special = float32(Keikopganjil_bet_5050special_RD)
			obj_5050.Keikopgenap_bet_5050special = float32(Keikopgenap_bet_5050special_RD)
			obj_5050.Kei_belakangkembang_bet_5050kombinasi = float32(Kei_belakangkembang_bet_5050kombinasi_RD)
			obj_5050.Kei_belakangkembar_bet_5050kombinasi = float32(Kei_belakangkembar_bet_5050kombinasi_RD)
			obj_5050.Kei_belakangkempis_bet_5050kombinasi = float32(Kei_belakangkempis_bet_5050kombinasi_RD)
			obj_5050.Kei_belakangmono_bet_5050kombinasi = float32(Kei_belakangmono_bet_5050kombinasi_RD)
			obj_5050.Kei_belakangstereo_bet_5050kombinasi = float32(Kei_belakangstereo_bet_5050kombinasi_RD)
			obj_5050.Kei_depankembang_bet_5050kombinasi = float32(Kei_depankembang_bet_5050kombinasi_RD)
			obj_5050.Kei_depankembar_bet_5050kombinasi = float32(Kei_depankembar_bet_5050kombinasi_RD)
			obj_5050.Kei_depankempis_bet_5050kombinasi = float32(Kei_depankempis_bet_5050kombinasi_RD)
			obj_5050.Kei_depanmono_bet_5050kombinasi = float32(Kei_depanmono_bet_5050kombinasi_RD)
			obj_5050.Kei_depanstereo_bet_5050kombinasi = float32(Kei_depanstereo_bet_5050kombinasi_RD)
			obj_5050.Kei_tengahkembang_bet_5050kombinasi = float32(Kei_tengahkembang_bet_5050kombinasi_RD)
			obj_5050.Kei_tengahkembar_bet_5050kombinasi = float32(Kei_tengahkembar_bet_5050kombinasi_RD)
			obj_5050.Kei_tengahkempis_bet_5050kombinasi = float32(Kei_tengahkempis_bet_5050kombinasi_RD)
			obj_5050.Kei_tengahmono_bet_5050kombinasi = float32(Kei_tengahmono_bet_5050kombinasi_RD)
			obj_5050.Kei_tengahstereo_bet_5050kombinasi = float32(Kei_tengahstereo_bet_5050kombinasi_RD)
			obj_5050.Limittotal_bet_5050umum = int(Limittotal_bet_5050umum_RD)
			obj_5050.Limittotal_bet_5050special = int(Limittotal_bet_5050special_RD)
			obj_5050.Limittotal_bet_5050kombinasi = int(Limittotal_bet_5050kombinasi_RD)
			arraobj_5050 = append(arraobj_5050, obj_5050)
		case "macaukombinasi":
			Min_bet_RD, _ := jsonparser.GetInt(value, "min_bet")
			Max_bet_RD, _ := jsonparser.GetInt(value, "max_bet")
			Diskon_bet_RD, _ := jsonparser.GetFloat(value, "diskon_bet")
			Win_bet_RD, _ := jsonparser.GetFloat(value, "Win_bet")
			Limit_total_RD, _ := jsonparser.GetInt(value, "limit_total")

			obj_macaukombinasi.Min_bet = int(Min_bet_RD)
			obj_macaukombinasi.Max_bet = int(Max_bet_RD)
			obj_macaukombinasi.Diskon_bet = float32(Diskon_bet_RD)
			obj_macaukombinasi.Win_bet = float32(Win_bet_RD)
			obj_macaukombinasi.Limit_total = int(Limit_total_RD)

			arraobj_macaukombinasi = append(arraobj_macaukombinasi, obj_macaukombinasi)
		case "dasar":
			Min_bet_RD, _ := jsonparser.GetInt(value, "min_bet")
			Max_bet_RD, _ := jsonparser.GetInt(value, "max_bet")
			Disc_besar_bet_RD, _ := jsonparser.GetFloat(value, "disc_besar_bet")
			Disc_kecil_bet_RD, _ := jsonparser.GetFloat(value, "disc_kecil_bet")
			Disc_ganjil_bet_RD, _ := jsonparser.GetFloat(value, "disc_ganjil_bet")
			Disc_genap_bet_RD, _ := jsonparser.GetFloat(value, "disc_genap_bet")
			Kei_besar_bet_RD, _ := jsonparser.GetFloat(value, "kei_besar_bet")
			Kei_kecil_bet_RD, _ := jsonparser.GetFloat(value, "kei_kecil_bet")
			Kei_ganjil_bet_RD, _ := jsonparser.GetFloat(value, "kei_ganjil_bet")
			Kei_genap_bet_RD, _ := jsonparser.GetFloat(value, "kei_genap_bet")
			Limit_total_RD, _ := jsonparser.GetInt(value, "limit_total")

			obj_dasar.Min_bet = int(Min_bet_RD)
			obj_dasar.Max_bet = int(Max_bet_RD)
			obj_dasar.Disc_besar_bet = float32(Disc_besar_bet_RD)
			obj_dasar.Disc_kecil_bet = float32(Disc_kecil_bet_RD)
			obj_dasar.Disc_ganjil_bet = float32(Disc_ganjil_bet_RD)
			obj_dasar.Disc_genap_bet = float32(Disc_genap_bet_RD)
			obj_dasar.Kei_besar_bet = float32(Kei_besar_bet_RD)
			obj_dasar.Kei_kecil_bet = float32(Kei_kecil_bet_RD)
			obj_dasar.Kei_ganjil_bet = float32(Kei_ganjil_bet_RD)
			obj_dasar.Kei_genap_bet = float32(Kei_genap_bet_RD)
			obj_dasar.Limit_total = int(Limit_total_RD)

			arraobj_dasar = append(arraobj_dasar, obj_dasar)
		case "shio":
			Min_bet_RD, _ := jsonparser.GetInt(value, "min_bet")
			Max_bet_RD, _ := jsonparser.GetInt(value, "max_bet")
			Diskon_bet_RD, _ := jsonparser.GetFloat(value, "diskon_bet")
			Win_bet_RD, _ := jsonparser.GetFloat(value, "win_bet")
			Limit_total_RD, _ := jsonparser.GetInt(value, "limit_total")

			obj_shio.Min_bet = int(Min_bet_RD)
			obj_shio.Max_bet = int(Max_bet_RD)
			obj_shio.Diskon_bet = float32(Diskon_bet_RD)
			obj_shio.Win_bet = float32(Win_bet_RD)
			obj_shio.Limit_total = int(Limit_total_RD)

			arraobj_shio = append(arraobj_shio, obj_shio)
		}

	})
	if !flag {
		result, err := model.FetchAll_MinitPasaran(client.Client_Company, client.Pasaran_Code, client.Permainan)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}

		log.Println("mysql")
		if result.Status == 200 {
			helpers.SetRedis(field_redis, result, 0)
		}
		return c.JSON(result)
	} else {
		log.Println("cache")
		switch client.Permainan {
		case "colok":
			arraobj := &arraobj_colok
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		case "5050":
			arraobj := &arraobj_5050
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		case "macaukombinasi":
			arraobj := &arraobj_macaukombinasi
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		case "dasar":
			arraobj := &arraobj_dasar
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		case "shio":
			arraobj := &arraobj_shio
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		default:
			arraobj := &arraobj_432
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "Success",
				"record":  arraobj,
				"time":    time.Since(render_page).String(),
			})
		}
	}
}
func Fetch_LimitPasaran432(c *fiber.Ctx) error {
	client := new(ClientLimitPasaran)
	if err := c.BodyParser(client); err != nil {
		return err
	}

	result, err := model.Fetch_LimitTransaksiPasaran432(client.Client_Username, client.Client_Company, client.Pasaran_Code, client.Pasaran_Periode, client.Permainan)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func Fetch_listinvoicebet(c *fiber.Ctx) error {
	client := new(ClientInvoicePasaran)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	field_redis := "listinvoice_" + strings.ToLower(client.Client_Company) + "_" + strconv.Itoa(client.Client_Idinvoice) + "_" + strings.ToLower(client.Client_Username)
	render_page := time.Now()
	var obj responseredislistinvoicebet
	var arraobj []responseredislistinvoicebet
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tanggal, _ := jsonparser.GetString(value, "tanggal")
		permainan, _ := jsonparser.GetString(value, "permainan")
		periode, _ := jsonparser.GetString(value, "periode")
		nomor, _ := jsonparser.GetString(value, "nomor")
		bet, _ := jsonparser.GetInt(value, "bet")
		diskon, _ := jsonparser.GetFloat(value, "diskon")
		kei, _ := jsonparser.GetFloat(value, "kei")
		bayar, _ := jsonparser.GetInt(value, "bayar")
		win, _ := jsonparser.GetInt(value, "win")
		menang, _ := jsonparser.GetInt(value, "menang")

		obj.Tanggal = tanggal
		obj.Permainan = permainan
		obj.Periode = periode
		obj.Nomor = nomor
		obj.Bet = int(bet)
		obj.Diskon = float32(diskon)
		obj.Kei = float32(kei)
		obj.Bayar = int(bayar)
		obj.Win = int(win)
		obj.Menang = int(menang)
		arraobj = append(arraobj, obj)

	})

	if !flag {
		result, err := model.Fetch_invoicebet(client.Client_Username, client.Client_Company, client.Pasaran_Code, client.Pasaran_Periode)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 30*time.Minute)
		log.Println("LIST INVOICE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LIST INVOICE cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Fetch_listinvoicebetid(c *fiber.Ctx) error {
	client := new(ClientInvoicePasaranId)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	field_redis := "listinvoice_" + strconv.Itoa(client.Client_Idinvoice) + "_" + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Client_Username) + "_" + strings.ToLower(client.Permainan)
	render_page := time.Now()
	var obj responseinvoiceidpermainan
	var arraobj []responseinvoiceidpermainan
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		no, _ := jsonparser.GetInt(value, "no")
		status, _ := jsonparser.GetString(value, "status")
		permainan, _ := jsonparser.GetString(value, "permainan")
		nomor, _ := jsonparser.GetString(value, "nomor")
		bet, _ := jsonparser.GetInt(value, "bet")
		diskon, _ := jsonparser.GetFloat(value, "diskon")
		kei, _ := jsonparser.GetFloat(value, "kei")
		bayar, _ := jsonparser.GetInt(value, "bayar")
		win, _ := jsonparser.GetInt(value, "win")

		obj.No = int(no)
		obj.Status = status
		obj.Permainan = permainan
		obj.Nomor = nomor
		obj.Bet = int(bet)
		obj.Diskon = float32(diskon)
		obj.Kei = float32(kei)
		obj.Bayar = int(bayar)
		obj.Win = int(win)
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := model.Fetch_invoicebetbyid(
			client.Client_Idinvoice, client.Client_Username,
			client.Client_Company, client.Permainan)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 5*time.Minute)
		log.Println("MYSQL")
		return c.JSON(result)
	} else {
		log.Println("cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Fetch_slipperiode(c *fiber.Ctx) error {
	client := new(ClientSlipPeriode)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	result, err := model.Fetch_invoiceperiode(client.Client_Username, client.Client_Company, client.Pasaran_Code)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func Fetch_slipperiodeall(c *fiber.Ctx) error {
	client := new(ClientSlipPeriodeAll)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	field_redis := "listinvoiceall_" + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Client_Username)
	render_page := time.Now()
	var obj responseinvoiceall
	var arraobj []responseinvoiceall
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tglkeluaran, _ := jsonparser.GetString(value, "tglkeluaran")
		idinvoice, _ := jsonparser.GetString(value, "idinvoice")
		pasaran, _ := jsonparser.GetString(value, "pasaran")
		periode, _ := jsonparser.GetString(value, "periode")
		totalbet, _ := jsonparser.GetInt(value, "totalbet")
		totalbayar, _ := jsonparser.GetInt(value, "totalbayar")
		totalwin, _ := jsonparser.GetInt(value, "totalwin")
		totallose, _ := jsonparser.GetInt(value, "totallose")
		status, _ := jsonparser.GetString(value, "status")
		color_lost, _ := jsonparser.GetString(value, "color_lost")
		background, _ := jsonparser.GetString(value, "background")
		color_totallose, _ := jsonparser.GetString(value, "color_totallose")

		obj.Tglkeluaran = tglkeluaran
		obj.Idinvoice = idinvoice
		obj.Pasaran = pasaran
		obj.Periode = periode
		obj.Totalbet = int(totalbet)
		obj.Totalbayar = int(totalbayar)
		obj.Totalwin = int(totalwin)
		obj.Totallose = int(totallose)
		obj.Status = status
		obj.Color_lost = color_lost
		obj.Background = background
		obj.Color_totallose = color_totallose
		arraobj = append(arraobj, obj)

	})
	if !flag {
		result, err := model.Fetch_invoiceperiodeall(client.Client_Username, client.Client_Company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(field_redis, result, 5*time.Minute)
		log.Println("MYSQL")
		return c.JSON(result)
	} else {
		log.Println("cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}

}
func Fetch_slipperiodedetail(c *fiber.Ctx) error {
	client := new(ClientSlipPeriodeDetail)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	result, err := model.Fetch_invoiceperiodedetail(client.Client_Username, client.Client_Company, client.Idtrxkeluaran)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func SaveTogel(c *fiber.Ctx) error {
	client := new(ClientSaveTogel)
	if err := c.BodyParser(client); err != nil {
		panic(err.Error())
	}

	result, err := model.Savetransaksi(
		client.Client_Username,
		client.Client_Company, client.Idtrxkeluaran, client.Idcomppasaran, client.Devicemember, client.Formipaddress, client.Timezone, client.Totalbayarbet, client.List4d)
	if err != nil {
		// panic(err.Error())
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusAccepted,
			"message": err.Error(),
			"record":  nil,
		})
	}
	field_redis := "listinvoice_" + strings.ToLower(client.Client_Company) + "_" + client.Idtrxkeluaran + "_" + strings.ToLower(client.Client_Username)
	val := helpers.DeleteRedis(field_redis)
	log.Printf("DELETE REDIS INVOICE %d\n", val)

	//AGEN
	val_agen := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran)
	val_agenlistmember := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTMEMBER")
	val_agenlistbettable := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTBETTABLE")
	val_agen4d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_4D")
	val_agen3d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_3D")
	val_agen2d := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_2D")
	val_agen2dd := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_2DD")
	val_agen2dt := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_2DT")
	val_agencolokbebas := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_COLOK_BEBAS")
	val_agencolokmacau := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_COLOK_MACAU")
	val_agencoloknaga := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_COLOK_NAGA")
	val_agencolokjitu := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_COLOK_JITU")
	val_agen5050umum := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_50_50_UMUM")
	val_agen5050special := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_50_50_SPECIAL")
	val_agen5050kombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_50_50_KOMBINASI")
	val_agenmacaukombinasi := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_MACAU_KOMBINASI")
	val_agendasar := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_DASAR")
	val_agenshio := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTPERMAINAN_SHIO")
	val_agenall := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTBET_all")
	val_agenwinner := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTBET_winner")
	val_agencancel := helpers.DeleteRedis("LISTPERIODE_AGENT_" + strings.ToLower(client.Client_Company) + "_INVOICE_" + client.Idtrxkeluaran + "_LISTBET_cancel")
	log.Printf("DELETE REDIS AGEN PERIODE %d\n", val_agen)
	log.Printf("DELETE REDIS AGEN PERIODE LISTMEMBER %d\n", val_agenlistmember)
	log.Printf("DELETE REDIS AGEN PERIODE LISTBETTABEL %d\n", val_agenlistbettable)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 4D %d\n", val_agen4d)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 3D %d\n", val_agen3d)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 2D %d\n", val_agen2d)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 2DD %d\n", val_agen2dd)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 2DT %d\n", val_agen2dt)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN COLOK BEBAS %d\n", val_agencolokbebas)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN COLOK MACAU %d\n", val_agencolokmacau)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN COLOK NAGA %d\n", val_agencoloknaga)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN COLOK JITU %d\n", val_agencolokjitu)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 5050UMUM %d\n", val_agen5050umum)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 5050SPECIAL %d\n", val_agen5050special)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN 5050KOMBINASI %d\n", val_agen5050kombinasi)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN MACAU KOMBINASI %d\n", val_agenmacaukombinasi)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN DASAR %d\n", val_agendasar)
	log.Printf("DELETE REDIS AGEN PERIODE PERMAINAN SHIO %d\n", val_agenshio)
	log.Printf("DELETE REDIS AGEN PERIODE LIST BET ALL %d\n", val_agenall)
	log.Printf("DELETE REDIS AGEN PERIODE LIST BET WINNER %d\n", val_agenwinner)
	log.Printf("DELETE REDIS AGEN PERIODE LIST BET CANCEL %d\n", val_agencancel)
	return c.JSON(result)
}
