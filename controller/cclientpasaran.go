package controller

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/api_go/entities"
	"bitbucket.org/isbtotogroup/api_go/helpers"
	"bitbucket.org/isbtotogroup/api_go/model"
	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
)

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
	Pasaran_note           string `json:"pasaran_note"`
	Pasaran_url            string `json:"pasaran_url"`
	Pasaran_status         string `json:"pasaran_status"`
}

type responseredisinit_432 struct {
	Min_bet                  int     `json:"min_bet"`
	Max4d_bet                int     `json:"max4d_bet"`
	Max3d_bet                int     `json:"max3d_bet"`
	Max3dd_bet               int     `json:"max3dd_bet"`
	Max2d_bet                int     `json:"max2d_bet"`
	Max2dd_bet               int     `json:"max2dd_bet"`
	Max2dt_bet               int     `json:"max2dt_bet"`
	Max4d_fullbb_bet         int     `json:"max4d_fullbb_bet"`
	Max3d_fullbb_bet         int     `json:"max3d_fullbb_bet"`
	Max3dd_fullbb_bet        int     `json:"max3dd_fullbb_bet"`
	Max2d_fullbb_bet         int     `json:"max2d_fullbb_bet"`
	Max2dd_fullbb_bet        int     `json:"max2dd_fullbb_bet"`
	Max2dt_fullbb_bet        int     `json:"max2dt_fullbb_bet"`
	Max4d_buy                int     `json:"max4d_buy"`
	Max3d_buy                int     `json:"max3d_buy"`
	Max3dd_buy               int     `json:"max3dd_buy"`
	Max2d_buy                int     `json:"max2d_buy"`
	Max2dd_buy               int     `json:"max2dd_buy"`
	Max2dt_buy               int     `json:"max2dt_buy"`
	Disc4d_bet               float32 `json:"disc4d_bet"`
	Disc3d_bet               float32 `json:"disc3d_bet"`
	Disc3dd_bet              float32 `json:"disc3dd_bet"`
	Disc2d_bet               float32 `json:"disc2d_bet"`
	Disc2dd_bet              float32 `json:"disc2dd_bet"`
	Disc2dt_bet              float32 `json:"disc2dt_bet"`
	Win4d_bet                int     `json:"win4d_bet"`
	Win3d_bet                int     `json:"win3d_bet"`
	Win3dd_bet               int     `json:"win3dd_bet"`
	Win2d_bet                int     `json:"win2d_bet"`
	Win2dd_bet               int     `json:"win2dd_bet"`
	Win2dt_bet               int     `json:"win2dt_bet"`
	Win4dnodiskon_bet        int     `json:"win4dnodiskon_bet"`
	Win3dnodiskon_bet        int     `json:"win3dnodiskon_bet"`
	Win3ddnodiskon_bet       int     `json:"win3ddnodiskon_bet"`
	Win2dnodiskon_bet        int     `json:"win2dnodiskon_bet"`
	Win2ddnodiskon_bet       int     `json:"win2ddnodiskon_bet"`
	Win2dtnodiskon_bet       int     `json:"win2dtnodiskon_bet"`
	Win4dbb_kena_bet         int     `json:"win4dbb_kena_bet"`
	Win3dbb_kena_bet         int     `json:"win3dbb_kena_bet"`
	Win3ddbb_kena_bet        int     `json:"win3ddbb_kena_bet"`
	Win2dbb_kena_bet         int     `json:"win2dbb_kena_bet"`
	Win2ddbb_kena_bet        int     `json:"win2ddbb_kena_bet"`
	Win2dtbb_kena_bet        int     `json:"win2dtbb_kena_bet"`
	Win4dbb_bet              int     `json:"win4dbb_bet"`
	Win3dbb_bet              int     `json:"win3dbb_bet"`
	Win3ddbb_bet             int     `json:"win3ddbb_bet"`
	Win2dbb_bet              int     `json:"win2dbb_bet"`
	Win2ddbb_bet             int     `json:"win2ddbb_bet"`
	Win2dtbb_bet             int     `json:"win2dtbb_bet"`
	Bbfs                     int     `json:"bbfs"`
	Limitline_4d             int     `json:"limitline_4d"`
	Limitline_3d             int     `json:"limitline_3d"`
	Limitline_3dd            int     `json:"limitline_3dd"`
	Limitline_2d             int     `json:"limitline_2d"`
	Limitline_2dd            int     `json:"limitline_2dd"`
	Limitline_2dt            int     `json:"limitline_2dt"`
	Limittotal4d_bet         int     `json:"limittotal4d_bet"`
	Limittotal3d_bet         int     `json:"limittotal3d_bet"`
	Limittotal3dd_bet        int     `json:"limittotal3dd_bet"`
	Limittotal2d_bet         int     `json:"limittotal2d_bet"`
	Limittotal2dd_bet        int     `json:"limittotal2dd_bet"`
	Limittotal2dt_bet        int     `json:"limittotal2dt_bet"`
	Limittotal4d_fullbb_bet  int     `json:"limittotal4d_fullbb_bet"`
	Limittotal3d_fullbb_bet  int     `json:"limittotal3d_fullbb_bet"`
	Limittotal3dd_fullbb_bet int     `json:"limittotal3dd_fullbb_bet"`
	Limittotal2d_fullbb_bet  int     `json:"limittotal2d_fullbb_bet"`
	Limittotal2dd_fullbb_bet int     `json:"limittotal2dd_fullbb_bet"`
	Limittotal2dt_fullbb_bet int     `json:"limittotal2dt_fullbb_bet"`
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
	Max_buy_colokbebas        int     `json:"max_buy_colokbebas"`
	Max_buy_colokmacau        int     `json:"max_buy_colokmacau"`
	Max_buy_coloknaga         int     `json:"max_buy_coloknaga"`
	Max_buy_colokjitu         int     `json:"max_buy_colokjitu"`
	Disc_bet_colokbebas       float32 `json:"disc_bet_colokbebas"`
	Disc_bet_colokmacau       float32 `json:"disc_bet_colokmacau"`
	Disc_bet_coloknaga        float32 `json:"disc_bet_coloknaga"`
	Disc_bet_colokjitu        float32 `json:"disc_bet_colokjitu"`
	Win_bet_colokbebas        float32 `json:"win_bet_colokbebas"`
	Win_bet_colokmacau        float32 `json:"win_bet_colokmacau"`
	Win_bet_coloknaga         float32 `json:"win_bet_coloknaga"`
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
	Max_buy_5050umum                       int     `json:"max_buy_5050umum"`
	Max_buy_5050special                    int     `json:"max_buy_5050special"`
	Max_buy_5050kombinasi                  int     `json:"max_buy_5050kombinasi"`
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
	Max_buy     int     `json:"max_buy"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Win_bet     float32 `json:"win_bet"`
	Limit_total int     `json:"limit_total"`
}
type responseredisinit_dasar struct {
	Min_bet         int     `json:"min_bet"`
	Max_bet         int     `json:"max_bet"`
	Max_buy         int     `json:"max_buy"`
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
	Max_buy     int     `json:"max_buy"`
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
	Tipe      string  `json:"tipe"`
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
	Tipe      string  `json:"tipe"`
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
type responseaxios struct {
	Status          int    `json:"status"`
	Token           string `json:"token"`
	Member_company  string `json:"member_company"`
	Member_username string `json:"member_username"`
	Member_credit   int    `json:"member_credit"`
}
type responseaxioserror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var ctx = context.Background()

const fielddomain_redis = "LISTDOMAIN"
const fieldsetting_redis = "LISTSETTING_MASTER"
const fieldallpasaran_redis = "listpasaran_"
const fieldresult_redis = "listresult_"
const fieldconfig_redis = "config_"
const fieldinvoice_redis = "listinvoice_"
const fieldlimit_redis = "limitpasaran_"

func Fetch_token(c *fiber.Ctx) error {
	client := new(entities.Controller_clientToken)

	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	tglnow, _ := goment.New()
	website_status := "ONLINE"
	website_message := ""
	tglskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
	jamstart := ""
	jamend := ""
	resultredis, flag := helpers.GetRedis(fieldsetting_redis)
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
		website_message = "MAINTENANCE<br>START : " + jamstart + "<br>FINISH : " + jamend
	}
	member_company, member_username, member_saldo := Fetch_apimoney(client.Token)

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
		var obj entities.Controller_settingcustom
		var obj2 entities.Model_setting
		var arraobj2 []entities.Model_setting

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
		helpers.SetRedis(fieldsetting_redis, obj, 24*time.Hour)

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
func Fetch_apimoney(token string) (string, string, int) {
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responseaxios{}).
		SetError(responseaxioserror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"token": token,
		}).
		Post("http://128.199.241.112:6061/api/servicetoken")
	if err != nil {
		log.Println(err.Error())
	}
	result := resp.Result().(*responseaxios)
	return result.Member_company, result.Member_username, result.Member_credit
}
func FetchAll_pasaran(c *fiber.Ctx) error {
	client := new(entities.Controller_clientInit)

	if err := c.BodyParser(client); err != nil {
		return err
	}

	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	render_page := time.Now()
	tglnow, _ := goment.New()
	var obj responseredisfetch
	var arraobj []responseredisfetch
	resultredis, flag := helpers.GetRedis(fieldallpasaran_redis + strings.ToLower(client.Client_Company))
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
		pasaran_note, _ := jsonparser.GetString(value, "pasaran_note")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")

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
		obj.Pasaran_note = pasaran_note
		obj.Pasaran_url = pasaran_url
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
		helpers.SetRedis(fieldallpasaran_redis+strings.ToLower(client.Client_Company), result, 24*time.Hour)
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
	client := new(entities.Controller_clientResult)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
	var obj entities.Model_mclientpasaranResult
	var arraobj []entities.Model_mclientpasaranResult
	resultredis, flag := helpers.GetRedis(fieldresult_redis + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Pasaran_Code))
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
			helpers.SetRedis(fieldresult_redis+strings.ToLower(client.Client_Company)+"_"+strings.ToLower(client.Pasaran_Code), result, 24*time.Hour)
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
	client := new(entities.Controller_clientResultAll)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	var obj entities.Model_mclientpasaranResultAll
	var arraobj []entities.Model_mclientpasaranResultAll
	resultredis, flag := helpers.GetRedis(fieldresult_redis + strings.ToLower(client.Client_Company))
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
		obj.Pasaran_code = pasaran_code_RD
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
			helpers.SetRedis(fieldresult_redis+strings.ToLower(client.Client_Company), result, 24*time.Hour)
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
	client := new(entities.Controller_clientResult)

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
	client := new(entities.Controller_clientConfPasaran)
	render_page := time.Now()
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
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
	resultredis, flag := helpers.GetRedis(fieldconfig_redis + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Pasaran_Code) + "_" + strings.ToLower(client.Permainan))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch client.Permainan {
		case "4-3-2":
			minbet_RD, _ := jsonparser.GetInt(value, "min_bet")
			maxbet4d_RD, _ := jsonparser.GetInt(value, "max4d_bet")
			maxbet3d_RD, _ := jsonparser.GetInt(value, "max3d_bet")
			maxbet3dd_RD, _ := jsonparser.GetInt(value, "max3dd_bet")
			maxbet2d_RD, _ := jsonparser.GetInt(value, "max2d_bet")
			maxbet2dd_RD, _ := jsonparser.GetInt(value, "max2dd_bet")
			maxbet2dt_RD, _ := jsonparser.GetInt(value, "max2dt_bet")
			max4d_fullbb_bet, _ := jsonparser.GetInt(value, "max4d_fullbb_bet")
			max3d_fullbb_bet, _ := jsonparser.GetInt(value, "max3d_fullbb_bet")
			max3dd_fullbb_bet, _ := jsonparser.GetInt(value, "max3dd_fullbb_bet")
			max2d_fullbb_bet, _ := jsonparser.GetInt(value, "max2d_fullbb_bet")
			max2dd_fullbb_bet, _ := jsonparser.GetInt(value, "max2dd_fullbb_bet")
			max2dt_fullbb_bet, _ := jsonparser.GetInt(value, "max2dt_fullbb_bet")
			max4d_buy, _ := jsonparser.GetInt(value, "max4d_buy")
			max3d_buy, _ := jsonparser.GetInt(value, "max3d_buy")
			max3dd_buy, _ := jsonparser.GetInt(value, "max3dd_buy")
			max2d_buy, _ := jsonparser.GetInt(value, "max2d_buy")
			max2dd_buy, _ := jsonparser.GetInt(value, "max2dd_buy")
			max2dt_buy, _ := jsonparser.GetInt(value, "max2dt_buy")
			disc4d_RD, _ := jsonparser.GetFloat(value, "disc4d_bet")
			disc3d_RD, _ := jsonparser.GetFloat(value, "disc3d_bet")
			disc3dd_RD, _ := jsonparser.GetFloat(value, "disc3dd_bet")
			disc2d_RD, _ := jsonparser.GetFloat(value, "disc2d_bet")
			disc2dd_RD, _ := jsonparser.GetFloat(value, "disc2dd_bet")
			disc2dt_RD, _ := jsonparser.GetFloat(value, "disc2dt_bet")
			win4d_RD, _ := jsonparser.GetInt(value, "win4d_bet")
			win3d_RD, _ := jsonparser.GetInt(value, "win3d_bet")
			win3dd_RD, _ := jsonparser.GetInt(value, "win3dd_bet")
			win2d_RD, _ := jsonparser.GetInt(value, "win2d_bet")
			win2dd_RD, _ := jsonparser.GetInt(value, "win2dd_bet")
			win2dt_RD, _ := jsonparser.GetInt(value, "win2dt_bet")
			win4dnodiskon_RD, _ := jsonparser.GetInt(value, "win4dnodiskon_bet")
			win3dnodiskon_RD, _ := jsonparser.GetInt(value, "win3dnodiskon_bet")
			win3ddnodiskon_RD, _ := jsonparser.GetInt(value, "win3ddnodiskon_bet")
			win2dnodiskon_RD, _ := jsonparser.GetInt(value, "win2dnodiskon_bet")
			win2ddnodiskon_RD, _ := jsonparser.GetInt(value, "win2ddnodiskon_bet")
			win2dtnodiskon_RD, _ := jsonparser.GetInt(value, "win2dtnodiskon_bet")
			win4dbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win4dbb_kena_bet")
			win3dbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win3dbb_kena_bet")
			win3ddbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win3ddbb_kena_bet")
			win2dbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win2dbb_kena_bet")
			win2ddbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win2ddbb_kena_bet")
			win2dtbb_kena_bet_RD, _ := jsonparser.GetInt(value, "win2dtbb_kena_bet")
			win4dbb_bet_RD, _ := jsonparser.GetInt(value, "win4dbb_bet")
			win3dbb_bet_RD, _ := jsonparser.GetInt(value, "win3dbb_bet")
			win3ddbb_bet_RD, _ := jsonparser.GetInt(value, "win3ddbb_bet")
			win2dbb_bet_RD, _ := jsonparser.GetInt(value, "win2dbb_bet")
			win2ddbb_bet_RD, _ := jsonparser.GetInt(value, "win2ddbb_bet")
			win2dtbb_bet_RD, _ := jsonparser.GetInt(value, "win2dtbb_bet")
			limitline_4d_RD, _ := jsonparser.GetInt(value, "limitline_4d")
			limitline_3d_RD, _ := jsonparser.GetInt(value, "limitline_3d")
			limitline_3dd_RD, _ := jsonparser.GetInt(value, "limitline_3dd")
			limitline_2d_RD, _ := jsonparser.GetInt(value, "limitline_2d")
			limitline_2dd_RD, _ := jsonparser.GetInt(value, "limitline_2dd")
			limitline_2dt_RD, _ := jsonparser.GetInt(value, "limitline_2dt")
			limittotal4d_RD, _ := jsonparser.GetInt(value, "limittotal4d_bet")
			limittotal3d_RD, _ := jsonparser.GetInt(value, "limittotal3d_bet")
			limittotal3dd_RD, _ := jsonparser.GetInt(value, "limittotal3dd_bet")
			limittotal2d_RD, _ := jsonparser.GetInt(value, "limittotal2d_bet")
			limittotal2dd_RD, _ := jsonparser.GetInt(value, "limittotal2dd_bet")
			limittotal2dt_RD, _ := jsonparser.GetInt(value, "limittotal2dt_bet")
			limittotal4d_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal4d_fullbb_bet")
			limittotal3d_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal3d_fullbb_bet")
			limittotal3dd_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal3dd_fullbb_bet")
			limittotal2d_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal2d_fullbb_bet")
			limittotal2dd_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal2dd_fullbb_bet")
			limittotal2dt_fullbb_bet, _ := jsonparser.GetInt(value, "limittotal2dt_fullbb_bet")
			bbfs_RD, _ := jsonparser.GetInt(value, "bbfs")

			obj_432.Min_bet = int(minbet_RD)
			obj_432.Max4d_bet = int(maxbet4d_RD)
			obj_432.Max3d_bet = int(maxbet3d_RD)
			obj_432.Max3dd_bet = int(maxbet3dd_RD)
			obj_432.Max2d_bet = int(maxbet2d_RD)
			obj_432.Max2dd_bet = int(maxbet2dd_RD)
			obj_432.Max2dt_bet = int(maxbet2dt_RD)
			obj_432.Max4d_fullbb_bet = int(max4d_fullbb_bet)
			obj_432.Max3d_fullbb_bet = int(max3d_fullbb_bet)
			obj_432.Max3dd_fullbb_bet = int(max3dd_fullbb_bet)
			obj_432.Max2d_fullbb_bet = int(max2d_fullbb_bet)
			obj_432.Max2dd_fullbb_bet = int(max2dd_fullbb_bet)
			obj_432.Max2dt_fullbb_bet = int(max2dt_fullbb_bet)
			obj_432.Max4d_buy = int(max4d_buy)
			obj_432.Max3d_buy = int(max3d_buy)
			obj_432.Max3dd_buy = int(max3dd_buy)
			obj_432.Max2d_buy = int(max2d_buy)
			obj_432.Max2dd_buy = int(max2dd_buy)
			obj_432.Max2dt_buy = int(max2dt_buy)
			obj_432.Disc4d_bet = float32(disc4d_RD)
			obj_432.Disc3d_bet = float32(disc3d_RD)
			obj_432.Disc3dd_bet = float32(disc3dd_RD)
			obj_432.Disc2d_bet = float32(disc2d_RD)
			obj_432.Disc2dd_bet = float32(disc2dd_RD)
			obj_432.Disc2dt_bet = float32(disc2dt_RD)
			obj_432.Win4d_bet = int(win4d_RD)
			obj_432.Win3d_bet = int(win3d_RD)
			obj_432.Win3dd_bet = int(win3dd_RD)
			obj_432.Win2d_bet = int(win2d_RD)
			obj_432.Win2dd_bet = int(win2dd_RD)
			obj_432.Win2dt_bet = int(win2dt_RD)
			obj_432.Win4dnodiskon_bet = int(win4dnodiskon_RD)
			obj_432.Win3dnodiskon_bet = int(win3dnodiskon_RD)
			obj_432.Win3ddnodiskon_bet = int(win3ddnodiskon_RD)
			obj_432.Win2dnodiskon_bet = int(win2dnodiskon_RD)
			obj_432.Win2ddnodiskon_bet = int(win2ddnodiskon_RD)
			obj_432.Win2dtnodiskon_bet = int(win2dtnodiskon_RD)
			obj_432.Win4dbb_kena_bet = int(win4dbb_kena_bet_RD)
			obj_432.Win3dbb_kena_bet = int(win3dbb_kena_bet_RD)
			obj_432.Win3ddbb_kena_bet = int(win3ddbb_kena_bet_RD)
			obj_432.Win2dbb_kena_bet = int(win2dbb_kena_bet_RD)
			obj_432.Win2ddbb_kena_bet = int(win2ddbb_kena_bet_RD)
			obj_432.Win2dtbb_kena_bet = int(win2dtbb_kena_bet_RD)
			obj_432.Win4dbb_bet = int(win4dbb_bet_RD)
			obj_432.Win3dbb_bet = int(win3dbb_bet_RD)
			obj_432.Win3ddbb_bet = int(win3ddbb_bet_RD)
			obj_432.Win2dbb_bet = int(win2dbb_bet_RD)
			obj_432.Win2ddbb_bet = int(win2ddbb_bet_RD)
			obj_432.Win2dtbb_bet = int(win2dtbb_bet_RD)
			obj_432.Limitline_4d = int(limitline_4d_RD)
			obj_432.Limitline_3d = int(limitline_3d_RD)
			obj_432.Limitline_3dd = int(limitline_3dd_RD)
			obj_432.Limitline_2d = int(limitline_2d_RD)
			obj_432.Limitline_2dd = int(limitline_2dd_RD)
			obj_432.Limitline_2dt = int(limitline_2dt_RD)
			obj_432.Limittotal4d_bet = int(limittotal4d_RD)
			obj_432.Limittotal3d_bet = int(limittotal3d_RD)
			obj_432.Limittotal3dd_bet = int(limittotal3dd_RD)
			obj_432.Limittotal2d_bet = int(limittotal2d_RD)
			obj_432.Limittotal2dd_bet = int(limittotal2dd_RD)
			obj_432.Limittotal2dt_bet = int(limittotal2dt_RD)
			obj_432.Limittotal4d_fullbb_bet = int(limittotal4d_fullbb_bet)
			obj_432.Limittotal3d_fullbb_bet = int(limittotal3d_fullbb_bet)
			obj_432.Limittotal3dd_fullbb_bet = int(limittotal3dd_fullbb_bet)
			obj_432.Limittotal2d_fullbb_bet = int(limittotal2d_fullbb_bet)
			obj_432.Limittotal2dd_fullbb_bet = int(limittotal2dd_fullbb_bet)
			obj_432.Limittotal2dt_fullbb_bet = int(limittotal2dt_fullbb_bet)
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
			Max_buy_colokbebas_RD, _ := jsonparser.GetInt(value, "max_buy_colokbebas")
			Max_buy_colokmacau_RD, _ := jsonparser.GetInt(value, "max_buy_colokmacau")
			Max_buy_coloknaga_RD, _ := jsonparser.GetInt(value, "max_buy_coloknaga")
			Max_buy_colokjitu_RD, _ := jsonparser.GetInt(value, "max_buy_colokjitu")
			Disc_bet_colokbebas_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokbebas")
			Disc_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokmacau")
			Disc_bet_coloknaga_RD, _ := jsonparser.GetFloat(value, "disc_bet_coloknaga")
			Disc_bet_colokjitu_RD, _ := jsonparser.GetFloat(value, "disc_bet_colokjitu")
			Win_bet_colokbebas_RD, _ := jsonparser.GetFloat(value, "win_bet_colokbebas")
			Win_bet_colokmacau_RD, _ := jsonparser.GetFloat(value, "win_bet_colokmacau")
			Win_bet_coloknaga_RD, _ := jsonparser.GetFloat(value, "win_bet_coloknaga")
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
			obj_colok.Max_buy_colokbebas = int(Max_buy_colokbebas_RD)
			obj_colok.Max_buy_colokmacau = int(Max_buy_colokmacau_RD)
			obj_colok.Max_buy_coloknaga = int(Max_buy_coloknaga_RD)
			obj_colok.Max_buy_colokjitu = int(Max_buy_colokjitu_RD)
			obj_colok.Disc_bet_colokbebas = float32(Disc_bet_colokbebas_RD)
			obj_colok.Disc_bet_colokmacau = float32(Disc_bet_colokmacau_RD)
			obj_colok.Disc_bet_coloknaga = float32(Disc_bet_coloknaga_RD)
			obj_colok.Disc_bet_colokjitu = float32(Disc_bet_colokjitu_RD)
			obj_colok.Win_bet_colokbebas = float32(Win_bet_colokbebas_RD)
			obj_colok.Win_bet_colokmacau = float32(Win_bet_colokmacau_RD)
			obj_colok.Win_bet_coloknaga = float32(Win_bet_coloknaga_RD)
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
			Max_buy_5050umum_RD, _ := jsonparser.GetInt(value, "max_buy_5050umum")
			Max_buy_5050special_RD, _ := jsonparser.GetInt(value, "max_buy_5050special")
			Max_buy_5050kombinasi_RD, _ := jsonparser.GetInt(value, "max_buy_5050kombinasi")
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
			obj_5050.Max_buy_5050umum = int(Max_buy_5050umum_RD)
			obj_5050.Max_buy_5050special = int(Max_buy_5050special_RD)
			obj_5050.Max_buy_5050kombinasi = int(Max_buy_5050kombinasi_RD)
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
			Max_buy_RD, _ := jsonparser.GetInt(value, "max_buy")
			Diskon_bet_RD, _ := jsonparser.GetFloat(value, "diskon_bet")
			Win_bet_RD, _ := jsonparser.GetFloat(value, "win_bet")
			Limit_total_RD, _ := jsonparser.GetInt(value, "limit_total")

			obj_macaukombinasi.Min_bet = int(Min_bet_RD)
			obj_macaukombinasi.Max_bet = int(Max_bet_RD)
			obj_macaukombinasi.Max_buy = int(Max_buy_RD)
			obj_macaukombinasi.Diskon_bet = float32(Diskon_bet_RD)
			obj_macaukombinasi.Win_bet = float32(Win_bet_RD)
			obj_macaukombinasi.Limit_total = int(Limit_total_RD)

			arraobj_macaukombinasi = append(arraobj_macaukombinasi, obj_macaukombinasi)
		case "dasar":
			Min_bet_RD, _ := jsonparser.GetInt(value, "min_bet")
			Max_bet_RD, _ := jsonparser.GetInt(value, "max_bet")
			Max_buy_RD, _ := jsonparser.GetInt(value, "max_buy")
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
			obj_dasar.Max_buy = int(Max_buy_RD)
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
			Max_buy_RD, _ := jsonparser.GetInt(value, "max_buy")
			Diskon_bet_RD, _ := jsonparser.GetFloat(value, "diskon_bet")
			Win_bet_RD, _ := jsonparser.GetFloat(value, "win_bet")
			Limit_total_RD, _ := jsonparser.GetInt(value, "limit_total")

			obj_shio.Min_bet = int(Min_bet_RD)
			obj_shio.Max_bet = int(Max_bet_RD)
			obj_shio.Max_buy = int(Max_buy_RD)
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
			helpers.SetRedis(fieldconfig_redis+strings.ToLower(client.Client_Company)+"_"+strings.ToLower(client.Pasaran_Code)+"_"+strings.ToLower(client.Permainan), result, 24*time.Hour)
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
	client := new(entities.Controller_clientLimitPasaran)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
	render_page := time.Now()
	var obj entities.Model_mpasaranLimit
	resultredis, flag := helpers.GetRedis(
		fieldlimit_redis + strings.ToLower(client.Client_Company) + "_" +
			strings.ToLower(client.Client_Username) + "_" +
			strings.ToLower(client.Pasaran_Code) + "_" +
			strings.ToLower(client.Pasaran_Periode))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")

	total_4d, _ := jsonparser.GetInt(record_RD, "total_4d")
	total_3d, _ := jsonparser.GetInt(record_RD, "total_3d")
	total_3dd, _ := jsonparser.GetInt(record_RD, "total_3dd")
	total_2d, _ := jsonparser.GetInt(record_RD, "total_2d")
	total_2dd, _ := jsonparser.GetInt(record_RD, "total_2dd")
	total_2dt, _ := jsonparser.GetInt(record_RD, "total_2dt")
	total_colokbebas, _ := jsonparser.GetInt(record_RD, "total_colokbebas")
	total_colokmacau, _ := jsonparser.GetInt(record_RD, "total_colokmacau")
	total_coloknaga, _ := jsonparser.GetInt(record_RD, "total_coloknaga")
	total_colokjitu, _ := jsonparser.GetInt(record_RD, "total_colokjitu")
	total_5050umum, _ := jsonparser.GetInt(record_RD, "total_5050umum")
	total_5050special, _ := jsonparser.GetInt(record_RD, "total_5050special")
	total_5050kombinasi, _ := jsonparser.GetInt(record_RD, "total_5050kombinasi")
	total_macaukombinasi, _ := jsonparser.GetInt(record_RD, "total_macaukombinasi")
	total_dasar, _ := jsonparser.GetInt(record_RD, "total_dasar")
	total_shio, _ := jsonparser.GetInt(record_RD, "total_shio")
	total_4d_sum, _ := jsonparser.GetInt(record_RD, "total_4d_sum")
	total_3d_sum, _ := jsonparser.GetInt(record_RD, "total_3d_sum")
	total_3dd_sum, _ := jsonparser.GetInt(record_RD, "total_3dd_sum")
	total_2d_sum, _ := jsonparser.GetInt(record_RD, "total_2d_sum")
	total_2dd_sum, _ := jsonparser.GetInt(record_RD, "total_2dd_sum")
	total_2dt_sum, _ := jsonparser.GetInt(record_RD, "total_2dt_sum")
	total_colokbebas_sum, _ := jsonparser.GetInt(record_RD, "total_colokbebas_sum")
	total_colokmacau_sum, _ := jsonparser.GetInt(record_RD, "total_colokmacau_sum")
	total_coloknaga_sum, _ := jsonparser.GetInt(record_RD, "total_coloknaga_sum")
	total_colokjitu_sum, _ := jsonparser.GetInt(record_RD, "total_colokjitu_sum")
	total_5050umum_sum, _ := jsonparser.GetInt(record_RD, "total_5050umum_sum")
	total_5050special_sum, _ := jsonparser.GetInt(record_RD, "total_5050special_sum")
	total_5050kombinasi_sum, _ := jsonparser.GetInt(record_RD, "total_5050kombinasi_sum")
	total_macaukombinasi_sum, _ := jsonparser.GetInt(record_RD, "total_macaukombinasi_sum")
	total_dasar_sum, _ := jsonparser.GetInt(record_RD, "total_dasar_sum")
	total_shio_sum, _ := jsonparser.GetInt(record_RD, "total_shio_sum")

	obj.Total_4d = int(total_4d)
	obj.Total_3d = int(total_3d)
	obj.Total_3dd = int(total_3dd)
	obj.Total_2d = int(total_2d)
	obj.Total_2dd = int(total_2dd)
	obj.Total_2dt = int(total_2dt)
	obj.Total_colokbebas = int(total_colokbebas)
	obj.Total_colokmacau = int(total_colokmacau)
	obj.Total_coloknaga = int(total_coloknaga)
	obj.Total_colokjitu = int(total_colokjitu)
	obj.Total_5050umum = int(total_5050umum)
	obj.Total_5050special = int(total_5050special)
	obj.Total_5050kombinasi = int(total_5050kombinasi)
	obj.Total_macaukombinasi = int(total_macaukombinasi)
	obj.Total_dasar = int(total_dasar)
	obj.Total_shio = int(total_shio)
	obj.Total_4d_sum = int(total_4d_sum)
	obj.Total_3d_sum = int(total_3d_sum)
	obj.Total_3dd_sum = int(total_3dd_sum)
	obj.Total_2d_sum = int(total_2d_sum)
	obj.Total_2dd_sum = int(total_2dd_sum)
	obj.Total_2dt_sum = int(total_2dt_sum)
	obj.Total_colokbebas_sum = int(total_colokbebas_sum)
	obj.Total_colokmacau_sum = int(total_colokmacau_sum)
	obj.Total_coloknaga_sum = int(total_coloknaga_sum)
	obj.Total_colokjitu_sum = int(total_colokjitu_sum)
	obj.Total_5050umum_sum = int(total_5050umum_sum)
	obj.Total_5050special_sum = int(total_5050special_sum)
	obj.Total_5050kombinasi_sum = int(total_5050kombinasi_sum)
	obj.Total_macaukombinasi_sum = int(total_macaukombinasi_sum)
	obj.Total_dasar_sum = int(total_dasar_sum)
	obj.Total_shio_sum = int(total_shio_sum)

	if !flag {
		result, err := model.Fetch_LimitTransaksiPasaran432(client.Client_Username, client.Client_Company, client.Permainan, client.Client_Idinvoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(
			fieldlimit_redis+strings.ToLower(client.Client_Company)+"_"+
				strings.ToLower(client.Client_Username)+"_"+
				strings.ToLower(client.Pasaran_Code)+"_"+
				strings.ToLower(client.Pasaran_Periode), result, 30*time.Minute)
		log.Println("LIMIT MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LIMIT cache")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  obj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Fetch_listinvoicebet(c *fiber.Ctx) error {
	client := new(entities.Controller_clientInvoicePasaran)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
	render_page := time.Now()
	var obj responseredislistinvoicebet
	var arraobj []responseredislistinvoicebet
	resultredis, flag := helpers.GetRedis(fieldinvoice_redis + strings.ToLower(client.Client_Company) + "_" + strconv.Itoa(client.Client_Idinvoice) + "_" + strings.ToLower(client.Client_Username))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	totalbet_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	totalbayar_RD, _ := jsonparser.GetInt(jsonredis, "totalbayar")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tanggal, _ := jsonparser.GetString(value, "tanggal")
		tipe, _ := jsonparser.GetString(value, "tipe")
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
		obj.Tipe = tipe
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
		result, err := model.Fetch_invoicebet(client.Client_Username, client.Client_Company, client.Client_Idinvoice)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(fieldinvoice_redis+strings.ToLower(client.Client_Company)+"_"+strconv.Itoa(client.Client_Idinvoice)+"_"+strings.ToLower(client.Client_Username), result, 30*time.Minute)
		log.Println("LIST INVOICE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("LIST INVOICE cache")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"totalrecord": int(totalbet_RD),
			"totalbayar":  int(totalbayar_RD),
			"time":        time.Since(render_page).String(),
		})
	}
}
func Fetch_listinvoicebetid(c *fiber.Ctx) error {
	client := new(entities.Controller_clientInvoicePasaranId)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
	render_page := time.Now()
	var obj responseinvoiceidpermainan
	var arraobj []responseinvoiceidpermainan
	resultredis, flag := helpers.GetRedis(fieldinvoice_redis + strconv.Itoa(client.Client_Idinvoice) + "_" + strings.ToLower(client.Client_Company) + "_" + strings.ToLower(client.Client_Username) + "_" + strings.ToLower(client.Permainan))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		no, _ := jsonparser.GetInt(value, "no")
		status, _ := jsonparser.GetString(value, "status")
		tipe, _ := jsonparser.GetString(value, "tipe")
		permainan, _ := jsonparser.GetString(value, "permainan")
		nomor, _ := jsonparser.GetString(value, "nomor")
		bet, _ := jsonparser.GetInt(value, "bet")
		diskon, _ := jsonparser.GetFloat(value, "diskon")
		kei, _ := jsonparser.GetFloat(value, "kei")
		bayar, _ := jsonparser.GetInt(value, "bayar")
		win, _ := jsonparser.GetInt(value, "win")

		obj.No = int(no)
		obj.Status = status
		obj.Tipe = tipe
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
		helpers.SetRedis(fieldinvoice_redis+strconv.Itoa(client.Client_Idinvoice)+"_"+strings.ToLower(client.Client_Company)+"_"+strings.ToLower(client.Client_Username)+"_"+strings.ToLower(client.Permainan), result, 5*time.Minute)
		log.Println("MYSQL")
		return c.JSON(result)
	} else {
		log.Println("CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Fetch_slipperiode(c *fiber.Ctx) error {
	client := new(entities.Controller_clientSlipPeriode)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
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
	client := new(entities.Controller_clientSlipPeriodeAll)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
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
	client := new(entities.Controller_clientSlipPeriodeDetail)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	flag_domain := _domainsecurity(client.Hostname)
	if !flag_domain {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
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

func _domainsecurity(nmdomain string) bool {
	log.Printf("Domain Client : %s", nmdomain)
	resultredis, flag_domain := helpers.GetRedis(fielddomain_redis)
	flag := false
	if len(nmdomain) > 0 {
		if !flag_domain {
			result, err := model.Get_Domain()
			if err != nil {
				flag = false
			}
			log.Println("DOMAIN MYSQL")
			helpers.SetRedis(fielddomain_redis, result, 24*time.Hour)
			flag = true
		} else {
			jsonredis_domain := []byte(resultredis)
			record_RD, _, _, _ := jsonparser.Get(jsonredis_domain, "record")
			jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				domain, _ := jsonparser.GetString(value, "domain")
				if nmdomain == domain {
					flag = true
					log.Println("DOMAIN CACHE")
				}
			})
		}
	}
	return flag
}
