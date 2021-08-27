package model

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiberapi/db"
	"github.com/nleeper/goment"
)

type Mclientpasaran struct {
	PasaranId          string `json:"pasaran_id"`
	PasaranTogel       string `json:"pasaran_togel"`
	PasaranPeriode     string `json:"pasaran_periode"`
	PasaranTglKeluaran string `json:"pasaran_tglkeluaran"`
	PasaranJamTutup    string `json:"pasaran_marketclose"`
	PasaranJamJadwal   string `json:"pasaran_marketschedule"`
	PasaranJamOpen     string `json:"pasaran_marketopen"`
	PasaranStatus      string `json:"pasaran_status"`
}
type MclientpasaranResult struct {
	No      uint16 `json:"no"`
	Date    string `json:"date"`
	Periode string `json:"periode"`
	Result  string `json:"result"`
}
type MclientpasaranCheckPasaran struct {
	PasaranIdtansaction string `json:"pasaran_idtransaction"`
	PasaranName         string `json:"pasaran_name"`
	PasaranPeriode      string `json:"pasaran_periode"`
	PasaranIdcomp       string `json:"pasaran_idcomp"`
	PasaranStatus       string `json:"pasaran_status"`
}
type Mpasarantogel432 struct {
	Min_bet           float32 `json:"min_bet"`
	Max4d_bet         float32 `json:"max4d_bet"`
	Max3d_bet         float32 `json:"max3d_bet"`
	Max2d_bet         float32 `json:"max2d_bet"`
	Max2dd_bet        float32 `json:"max2dd_bet"`
	Max2dt_bet        float32 `json:"max2dt_bet"`
	Disc4d_bet        float32 `json:"disc4d_bet"`
	Disc3d_bet        float32 `json:"disc3d_bet"`
	Disc2d_bet        float32 `json:"disc2d_bet"`
	Disc2dd_bet       float32 `json:"disc2dd_bet"`
	Disc2dt_bet       float32 `json:"disc2dt_bet"`
	Win4d_bet         float32 `json:"win4d_bet"`
	Win3d_bet         float32 `json:"win3d_bet"`
	Win2d_bet         float32 `json:"win2d_bet"`
	Win2dd_bet        float32 `json:"win2dd_bet"`
	Win2dt_bet        float32 `json:"win2dt_bet"`
	Limittotal4d_bet  float32 `json:"limittotal4d_bet"`
	Limittotal3d_bet  float32 `json:"limittotal3d_bet"`
	Limittotal2d_bet  float32 `json:"limittotal2d_bet"`
	Limittotal2dd_bet float32 `json:"limittotal2dd_bet"`
	Limittotal2dt_bet float32 `json:"limittotal2dt_bet"`
	Limitline_4d      uint32  `json:"limitline_4d"`
	Limitline_3d      uint32  `json:"limitline_3d"`
	Limitline_2d      uint32  `json:"limitline_2d"`
	Limitline_2dd     uint32  `json:"limitline_2dd"`
	Limitline_2dt     uint32  `json:"limitline_2dt"`
	Bbfs              uint8   `json:"bbfs"`
}
type MpasarantogelColok struct {
	Min_bet_colokbebas        float32 `json:"min_bet_colokbebas"`
	Max_bet_colokbebas        float32 `json:"max_bet_colokbebas"`
	Disc_bet_colokbebas       float32 `json:"disc_bet_colokbebas"`
	Win_bet_colokbebas        float32 `json:"win_bet_colokbebas"`
	Limittotal_bet_colokbebas float32 `json:"limittotal_bet_colokbebas"`
	Min_bet_colokmacau        float32 `json:"min_bet_colokmacau"`
	Max_bet_colokmacau        float32 `json:"max_bet_colokmacau"`
	Disc_bet_colokmacau       float32 `json:"disc_bet_colokmacau"`
	Win_bet_colokmacau        float32 `json:"win_bet_colokmacau"`
	Win3_bet_colokmacau       float32 `json:"win3_bet_colokmacau"`
	Win4_bet_colokmacau       float32 `json:"win4_bet_colokmacau"`
	Limittotal_bet_colokmacau float32 `json:"limittotal_bet_colokmacau"`
	Min_bet_coloknaga         float32 `json:"min_bet_coloknaga"`
	Max_bet_coloknaga         float32 `json:"max_bet_coloknaga"`
	Disc_bet_coloknaga        float32 `json:"disc_bet_coloknaga"`
	Win_bet_coloknaga         float32 `json:"win_bet_coloknaga"`
	Win4_bet_coloknaga        float32 `json:"win4_bet_coloknaga"`
	Limittotal_bet_coloknaga  float32 `json:"limittotal_bet_coloknaga"`
	Min_bet_colokjitu         float32 `json:"min_bet_colokjitu"`
	Max_bet_colokjitu         float32 `json:"max_bet_colokjitu"`
	Disc_bet_colokjitu        float32 `json:"disc_bet_colokjitu"`
	Winas_bet_colokjitu       float32 `json:"winas_bet_colokjitu"`
	Winkop_bet_colokjitu      float32 `json:"winkop_bet_colokjitu"`
	Winkepala_bet_colokjitu   float32 `json:"winkepala_bet_colokjitu"`
	Winekor_bet_colokjitu     float32 `json:"winekor_bet_colokjitu"`
	Limittotal_bet_colokjitu  float32 `json:"limittotal_bet_colokjitu"`
}
type Mpasarantogel5050 struct {
	Min_bet_5050umum                       float32 `json:"min_bet_5050umum"`
	Max_bet_5050umum                       float32 `json:"max_bet_5050umum"`
	Keibesar_bet_5050umum                  float32 `json:"keibesar_bet_5050umum"`
	Keikecil_bet_5050umum                  float32 `json:"keikecil_bet_5050umum"`
	Keigenap_bet_5050umum                  float32 `json:"keigenap_bet_5050umum"`
	Keiganjil_bet_5050umum                 float32 `json:"keiganjil_bet_5050umum"`
	Keitengah_bet_5050umum                 float32 `json:"keitengah_bet_5050umum"`
	Keitepi_bet_5050umum                   float32 `json:"keitepi_bet_5050umum"`
	Discbesar_bet_5050umum                 float32 `json:"discbesar_bet_5050umum"`
	Disckecil_bet_5050umum                 float32 `json:"disckecil_bet_5050umum"`
	Discgenap_bet_5050umum                 float32 `json:"discgenap_bet_5050umum"`
	Discganjil_bet_5050umum                float32 `json:"discganjil_bet_5050umum"`
	Disctengah_bet_5050umum                float32 `json:"disctengah_bet_5050umum"`
	Disctepi_bet_5050umum                  float32 `json:"disctepi_bet_5050umum"`
	Limittotal_bet_5050umum                float32 `json:"limittotal_bet_5050umum"`
	Min_bet_5050special                    float32 `json:"min_bet_5050special"`
	Max_bet_5050special                    float32 `json:"max_bet_5050special"`
	Keiasganjil_bet_5050special            float32 `json:"keiasganjil_bet_5050special"`
	Keiasgenap_bet_5050special             float32 `json:"keiasgenap_bet_5050special"`
	Keiasbesar_bet_5050special             float32 `json:"keiasbesar_bet_5050special"`
	Keiaskecil_bet_5050special             float32 `json:"keiaskecil_bet_5050special"`
	Keikopganjil_bet_5050special           float32 `json:"keikopganjil_bet_5050special"`
	Keikopgenap_bet_5050special            float32 `json:"keikopgenap_bet_5050special"`
	Keikopbesar_bet_5050special            float32 `json:"keikopbesar_bet_5050special"`
	Keikopkecil_bet_5050special            float32 `json:"keikopkecil_bet_5050special"`
	Keikepalaganjil_bet_5050special        float32 `json:"keikepalaganjil_bet_5050special"`
	Keikepalagenap_bet_5050special         float32 `json:"keikepalagenap_bet_5050special"`
	Keikepalabesar_bet_5050special         float32 `json:"keikepalabesar_bet_5050special"`
	Keikepalakecil_bet_5050special         float32 `json:"keikepalakecil_bet_5050special"`
	Keiekorganjil_bet_5050special          float32 `json:"keiekorganjil_bet_5050special"`
	Keiekorgenap_bet_5050special           float32 `json:"keiekorgenap_bet_5050special"`
	Keiekorbesar_bet_5050special           float32 `json:"keiekorbesar_bet_5050special"`
	Keiekorkecil_bet_5050special           float32 `json:"keiekorkecil_bet_5050special"`
	Discasganjil_bet_5050special           float32 `json:"discasganjil_bet_5050special"`
	Discasgenap_bet_5050special            float32 `json:"discasgenap_bet_5050special"`
	Discasbesar_bet_5050special            float32 `json:"discasbesar_bet_5050special"`
	Discaskecil_bet_5050special            float32 `json:"discaskecil_bet_5050special"`
	Disckopganjil_bet_5050special          float32 `json:"disckopganjil_bet_5050special"`
	Disckopgenap_bet_5050special           float32 `json:"disckopgenap_bet_5050special"`
	Disckopbesar_bet_5050special           float32 `json:"disckopbesar_bet_5050special"`
	Disckopkecil_bet_5050special           float32 `json:"disckopkecil_bet_5050special"`
	Disckepalaganjil_bet_5050special       float32 `json:"disckepalaganjil_bet_5050special"`
	Disckepalagenap_bet_5050special        float32 `json:"disckepalagenap_bet_5050special"`
	Disckepalabesar_bet_5050special        float32 `json:"disckepalabesar_bet_5050special"`
	Disckepalakecil_bet_5050special        float32 `json:"disckepalakecil_bet_5050special"`
	Discekorganjil_bet_5050special         float32 `json:"discekorganjil_bet_5050special"`
	Discekorgenap_bet_5050special          float32 `json:"discekorgenap_bet_5050special"`
	Discekorbesar_bet_5050special          float32 `json:"discekorbesar_bet_5050special"`
	Discekorkecil_bet_5050special          float32 `json:"discekorkecil_bet_5050special"`
	Limittotal_bet_5050special             float32 `json:"limittotal_bet_5050special"`
	Min_bet_5050kombinasi                  float32 `json:"min_bet_5050kombinasi"`
	Max_bet_5050kombinasi                  float32 `json:"max_bet_5050kombinasi"`
	Kei_belakangmono_bet_5050kombinasi     float32 `json:"kei_belakangmono_bet_5050kombinasi"`
	Kei_belakangstereo_bet_5050kombinasi   float32 `json:"kei_belakangstereo_bet_5050kombinasi"`
	Kei_belakangkembang_bet_5050kombinasi  float32 `json:"kei_belakangkembang_bet_5050kombinasi"`
	Kei_belakangkempis_bet_5050kombinasi   float32 `json:"kei_belakangkempis_bet_5050kombinasi"`
	Kei_belakangkembar_bet_5050kombinasi   float32 `json:"kei_belakangkembar_bet_5050kombinasi"`
	Kei_tengahmono_bet_5050kombinasi       float32 `json:"kei_tengahmono_bet_5050kombinasi"`
	Kei_tengahstereo_bet_5050kombinasi     float32 `json:"kei_tengahstereo_bet_5050kombinasi"`
	Kei_tengahkembang_bet_5050kombinasi    float32 `json:"kei_tengahkembang_bet_5050kombinasi"`
	Kei_tengahkempis_bet_5050kombinasi     float32 `json:"kei_tengahkempis_bet_5050kombinasi"`
	Kei_tengahkembar_bet_5050kombinasi     float32 `json:"kei_tengahkembar_bet_5050kombinasi"`
	Kei_depanmono_bet_5050kombinasi        float32 `json:"kei_depanmono_bet_5050kombinasi"`
	Kei_depanstereo_bet_5050kombinasi      float32 `json:"kei_depanstereo_bet_5050kombinasi"`
	Kei_depankembang_bet_5050kombinasi     float32 `json:"kei_depankembang_bet_5050kombinasi"`
	Kei_depankempis_bet_5050kombinasi      float32 `json:"kei_depankempis_bet_5050kombinasi"`
	Kei_depankembar_bet_5050kombinasi      float32 `json:"kei_depankembar_bet_5050kombinasi"`
	Disc_belakangmono_bet_5050kombinasi    float32 `json:"disc_belakangmono_bet_5050kombinasi"`
	Disc_belakangstereo_bet_5050kombinasi  float32 `json:"disc_belakangstereo_bet_5050kombinasi"`
	Disc_belakangkembang_bet_5050kombinasi float32 `json:"disc_belakangkembang_bet_5050kombinasi"`
	Disc_belakangkempis_bet_5050kombinasi  float32 `json:"disc_belakangkempis_bet_5050kombinasi"`
	Disc_belakangkembar_bet_5050kombinasi  float32 `json:"disc_belakangkembar_bet_5050kombinasi"`
	Disc_tengahmono_bet_5050kombinasi      float32 `json:"disc_tengahmono_bet_5050kombinasi"`
	Disc_tengahstereo_bet_5050kombinasi    float32 `json:"disc_tengahstereo_bet_5050kombinasi"`
	Disc_tengahkembang_bet_5050kombinasi   float32 `json:"disc_tengahkembang_bet_5050kombinasi"`
	Disc_tengahkempis_bet_5050kombinasi    float32 `json:"disc_tengahkempis_bet_5050kombinasi"`
	Disc_tengahkembar_bet_5050kombinasi    float32 `json:"disc_tengahkembar_bet_5050kombinasi"`
	Disc_depanmono_bet_5050kombinasi       float32 `json:"disc_depanmono_bet_5050kombinasi"`
	Disc_depanstereo_bet_5050kombinasi     float32 `json:"disc_depanstereo_bet_5050kombinasi"`
	Disc_depankembang_bet_5050kombinasi    float32 `json:"disc_depankembang_bet_5050kombinasi"`
	Disc_depankempis_bet_5050kombinasi     float32 `json:"disc_depankempis_bet_5050kombinasi"`
	Disc_depankembar_bet_5050kombinasi     float32 `json:"disc_depankembar_bet_5050kombinasi"`
	Limittotal_bet_5050kombinasi           float32 `json:"limittotal_bet_5050kombinasi"`
}
type MpasarantogelMacauKombinasi struct {
	Min_bet     float32 `json:"min_bet"`
	Max_bet     float32 `json:"max_bet"`
	Win_bet     float32 `json:"win_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Limit_total float32 `json:"limit_total"`
}
type MpasarantogelDasar struct {
	Min_bet         float32 `json:"min_bet"`
	Max_bet         float32 `json:"max_bet"`
	Kei_besar_bet   float32 `json:"kei_besar_bet"`
	Kei_kecil_bet   float32 `json:"kei_kecil_bet"`
	Kei_genap_bet   float32 `json:"kei_genap_bet"`
	Kei_ganjil_bet  float32 `json:"kei_ganjil_bet"`
	Disc_besar_bet  float32 `json:"disc_besar_bet"`
	Disc_kecil_bet  float32 `json:"disc_kecil_bet"`
	Disc_genap_bet  float32 `json:"disc_genap_bet"`
	Disc_ganjil_bet float32 `json:"disc_ganjil_bet"`
	Limit_total     float32 `json:"limit_total"`
}
type MpasarantogelShio struct {
	Min_bet     float32 `json:"min_bet"`
	Max_bet     float32 `json:"max_bet"`
	Win_bet     float32 `json:"win_bet"`
	Diskon_bet  float32 `json:"diskon_bet"`
	Limit_total float32 `json:"limit_total"`
}

func FetchAll_MclientPasaran(client_company string) (Response, error) {
	var obj Mclientpasaran
	var arraobj []Mclientpasaran
	var res Response
	var myDays = []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	statuspasaran := "ONLINE"
	msg := "Error"
	con := db.CreateCon()

	tglnow, _ := goment.New()
	daynow := tglnow.Format("d")
	intVar, _ := strconv.ParseInt(daynow, 0, 8)
	daynowhari := myDays[intVar]

	sqlpasaran := `SELECT 
		idcomppasaran, idpasarantogel, nmpasarantogel, jamtutup, jamjadwal, jamopen 
		FROM client_view_pasaran  
		WHERE statuspasaranactive = 'Y' 
		AND idcompany = ?
	`
	rowspasaran, err := con.Query(sqlpasaran, client_company)
	defer rowspasaran.Close()

	if err != nil {
		return res, err
	}
	for rowspasaran.Next() {
		statuspasaran = "ONLINE"
		var (
			idcomppasaran                                                int
			idpasarantogel, nmpasarantogel, jamtutup, jamjadwal, jamopen string
			tglkeluaran, periodekerluaran, haripasaran                   string
		)

		err = rowspasaran.Scan(
			&idcomppasaran,
			&idpasarantogel,
			&nmpasarantogel,
			&jamtutup,
			&jamjadwal,
			&jamopen)
		if err != nil {
			return res, err
		}

		sqlkeluaran := `
			SELECT 
			datekeluaran, keluaranperiode
			FROM 
				tbl_trx_keluarantogel 
			WHERE idcomppasaran = ?
			ORDER BY datekeluaran DESC
			LIMIT 1
		`
		err := con.QueryRow(sqlkeluaran, idcomppasaran).Scan(&tglkeluaran, &periodekerluaran)

		if err != nil {
			return res, err
		}

		sqlpasaranonline := `
			SELECT
				haripasaran
			FROM
				tbl_mst_company_game_pasaran_offline
			WHERE idcomppasaran = ?
			AND idcompany = ? 
			AND haripasaran = ? 
		`

		errpasaranonline := con.QueryRow(sqlpasaranonline, idcomppasaran, client_company, daynowhari).Scan(&haripasaran)

		if errpasaranonline != sql.ErrNoRows {
			taiskrg := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			jamtutup := tglnow.Format("YYYY-MM-DD") + " " + jamtutup
			jamopen := tglnow.Format("YYYY-MM-DD") + " " + jamopen
			tutup, _ := goment.New(jamtutup)
			open, _ := goment.New(jamopen)
			nowconvert := tglnow.Format("x")
			tutupconvert := tutup.Format("x")
			openconvert := open.Format("x")

			// intNow, _ := strconv.Atoi(nowconvert)
			// intTutup, _ := strconv.Atoi(tutupconvert)
			// intOpen, _ := strconv.Atoi(openconvert)

			if taiskrg >= jamtutup && taiskrg <= jamopen {
				statuspasaran = "OFFLINE"
			} else {
				statuspasaran = "ONLINE"
			}

			// log.Println(idpasarantogel + " - " + tglnow.Format("YYYY-MM-DD HH:mm:ss") + " - " + jamtutup + " - " + jamopen + " - " + statuspasaran)
			log.Println(idpasarantogel + " - " + nowconvert + " - " + tutupconvert + " - " + openconvert + " - " + statuspasaran)

		}

		obj.PasaranId = idpasarantogel
		obj.PasaranTogel = nmpasarantogel
		obj.PasaranPeriode = "#" + periodekerluaran + "-" + idpasarantogel
		obj.PasaranTglKeluaran = tglkeluaran
		obj.PasaranJamTutup = tglkeluaran + " " + jamtutup
		obj.PasaranJamJadwal = tglkeluaran + " " + jamjadwal
		obj.PasaranJamOpen = tglkeluaran + " " + jamopen
		obj.PasaranStatus = statuspasaran
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}

func FetchAll_MclientPasaranResult(client_company, pasaran_code string) (Response, error) {
	var obj MclientpasaranResult
	var arraobj []MclientpasaranResult
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT 
		A.keluaranperiode, A.datekeluaran, 
		A.keluarantogel, B.idpasarantogel 
		FROM tbl_trx_keluarantogel as A 
		JOIN tbl_mst_company_game_pasaran as B ON B.idcomppasaran = A.idcomppasaran
		WHERE B.idcompany = ? 
		AND B.idpasarantogel = ?
		AND A.keluarantogel != '' 
		ORDER BY A.datekeluaran DESC LIMIT 250
	`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	var norecord uint16 = 1
	for rowresult.Next() {
		var (
			keluaranperiode                             string
			datekeluaran, keluarantogel, idpasarantogel string
		)

		err = rowresult.Scan(
			&keluaranperiode,
			&datekeluaran,
			&keluarantogel,
			&idpasarantogel)
		if err != nil {
			return res, err
		}

		obj.No = norecord
		obj.Date = datekeluaran
		obj.Periode = idpasarantogel + "-" + keluaranperiode
		obj.Result = keluarantogel
		arraobj = append(arraobj, obj)
		msg = "Success"
		norecord = norecord + 1
	}
	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}

func CheckPasaran(client_company, pasaran_code string) (Response, error) {
	var obj MclientpasaranCheckPasaran
	var arraobj []MclientpasaranCheckPasaran
	var res Response
	var myDays = []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	statuspasaran := "ONLINE"
	msg := "Error"
	con := db.CreateCon()

	tglnow, _ := goment.New()
	daynow := tglnow.Format("d")
	intVar, _ := strconv.ParseInt(daynow, 0, 8)
	daynowhari := myDays[intVar]

	sqlpasaran := `SELECT 
		idcomppasaran, nmpasarantogel, 
		jamtutup, jamopen  
		FROM client_view_pasaran 
		WHERE idcompany = ? 
		AND idpasarantogel = ?
	`
	rowpasaran, err := con.Query(sqlpasaran, client_company, pasaran_code)
	defer rowpasaran.Close()

	if err != nil {
		return res, err
	}
	for rowpasaran.Next() {
		var (
			idcomppasaran, nmpasarantogel, jamtutup, jamopen string
			idtrxkeluaran, keluaranperiode, haripasaran      string
		)

		err = rowpasaran.Scan(
			&idcomppasaran,
			&nmpasarantogel,
			&jamtutup,
			&jamopen)
		if err != nil {
			return res, err
		}

		sqlkeluaran := `
			SELECT 
			idtrxkeluaran, keluaranperiode
			FROM 
				tbl_trx_keluarantogel 
			WHERE idcompany = ?
			AND idcomppasaran = ?
			AND keluarantogel = ''
			LIMIT 1
		`
		err := con.QueryRow(sqlkeluaran, client_company, idcomppasaran).Scan(&idtrxkeluaran, &keluaranperiode)
		if err != nil {
			return res, err
		}

		sqlpasaranonline := `
			SELECT
				haripasaran
			FROM
				tbl_mst_company_game_pasaran_offline
			WHERE idcomppasaran = ?
			AND idcompany = ? 
			AND haripasaran = ? 
		`

		errpasaranonline := con.QueryRow(sqlpasaranonline, idcomppasaran, client_company, daynowhari).Scan(&haripasaran)

		if errpasaranonline != sql.ErrNoRows {
			jamtutup := tglnow.Format("YYYY-MM-DD") + " " + jamtutup
			jamopen := tglnow.Format("YYYY-MM-DD") + " " + jamopen
			tutup, _ := goment.New(jamtutup)
			open, _ := goment.New(jamopen)
			nowconvert := tglnow.Format("x")
			tutupconvert := tutup.Format("x")
			openconvert := open.Format("x")

			intNow, _ := strconv.Atoi(nowconvert)
			intTutup, _ := strconv.Atoi(tutupconvert)
			intOpen, _ := strconv.Atoi(openconvert)

			if intNow > intTutup && intNow < intOpen {
				statuspasaran = "OFFLINE"
			}
		}

		obj.PasaranIdtansaction = idtrxkeluaran
		obj.PasaranName = nmpasarantogel
		obj.PasaranPeriode = keluaranperiode
		obj.PasaranIdcomp = idcomppasaran
		obj.PasaranStatus = statuspasaran
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func FetchAll_MinitPasaran432(client_company, pasaran_code string) (Response, error) {
	var obj Mpasarantogel432
	var arraobj []Mpasarantogel432
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT 
		1_minbet as min_bet, 
		1_maxbet4d as max4d_bet, 1_maxbet3d as max3d_bet, 
		1_maxbet2d as max2d_bet, 1_maxbet2dd as max2dd_bet, 1_maxbet2dt as max2dt_bet, 
		1_win4d as win4d_bet, 1_win3d as win3d_bet, 1_win2d as win2d_bet, 
		1_win2dd as win2dd_bet, 1_win2dt as win2dt_bet, 
		1_disc4d as disc4d_bet, 1_disc3d as disc3d_bet, 
		1_disc2d as disc2d_bet, 
		1_disc2dd as disc2dd_bet, 1_disc2dt as disc2dt_bet, 
		1_limittotal4d as limittotal4d_bet, 1_limittotal3d as limittotal3d_bet, 
		1_limittotal2d as limittotal2d_bet, 1_limittotal2dd as limittotal2dd_bet, 
		1_limittotal2dt as limittotal2dt_bet, 
		limitline_4d, limitline_3d, limitline_2d, limitline_2dd, limitline_2dt, bbfs 
		FROM tbl_mst_company_game_pasaran 
		WHERE idcompany = ? 
		AND idpasarantogel = ?
	`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet, max4d_bet, max3d_bet, max2d_bet, max2dd_bet, max2dt_bet                           float32
			disc4d_bet, disc3d_bet, disc2d_bet, disc2dd_bet, disc2dt_bet                               float32
			win4d_bet, win3d_bet, win2d_bet, win2dd_bet, win2dt_bet                                    float32
			limittotal4d_bet, limittotal3d_bet, limittotal2d_bet, limittotal2dd_bet, limittotal2dt_bet float32
			limitline_4d, limitline_3d, limitline_2d, limitline_2dd, limitline_2dt                     uint32
			bbfs                                                                                       uint8
		)

		err = rowresult.Scan(
			&min_bet, &max4d_bet, &max3d_bet, &max2d_bet, &max2dd_bet, &max2dt_bet,
			&disc4d_bet, &disc3d_bet, &disc2d_bet, &disc2dd_bet, &disc2dt_bet,
			&win4d_bet, &win3d_bet, &win2d_bet, &win2dd_bet, &win2dt_bet,
			&limittotal4d_bet, &limittotal3d_bet, &limittotal2d_bet, &limittotal2dd_bet, &limittotal2dt_bet,
			&limitline_4d, &limitline_3d, &limitline_2d, &limitline_2dd, &limitline_2dt,
			&bbfs)
		if err != nil {
			return res, err
		}
		obj.Min_bet = min_bet
		obj.Max4d_bet = max4d_bet
		obj.Max3d_bet = max3d_bet
		obj.Max2d_bet = max2d_bet
		obj.Max2dd_bet = max2dd_bet
		obj.Max2dt_bet = max2dt_bet
		obj.Disc4d_bet = disc4d_bet
		obj.Disc3d_bet = disc3d_bet
		obj.Disc2d_bet = disc2d_bet
		obj.Disc2dd_bet = disc2dd_bet
		obj.Disc2dt_bet = disc2dt_bet
		obj.Win4d_bet = win4d_bet
		obj.Win3d_bet = win3d_bet
		obj.Win2d_bet = win2d_bet
		obj.Win2dd_bet = win2dd_bet
		obj.Win2dt_bet = win2dt_bet
		obj.Limittotal4d_bet = limittotal4d_bet
		obj.Limittotal3d_bet = limittotal3d_bet
		obj.Limittotal2d_bet = limittotal2d_bet
		obj.Limittotal2dd_bet = limittotal2dd_bet
		obj.Limittotal2dt_bet = limittotal2dt_bet
		obj.Limitline_4d = limitline_4d
		obj.Limitline_3d = limitline_3d
		obj.Limitline_2d = limitline_2d
		obj.Limitline_2dd = limitline_2dd
		obj.Limitline_2dt = limitline_2dt

		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}

func FetchAll_MinitPasaranColok(client_company, pasaran_code string) (Response, error) {
	var obj MpasarantogelColok
	var arraobj []MpasarantogelColok
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT
			2_minbet as min_bet_colokbebas,
			2_maxbet as max_bet_colokbebas,
			2_disc as disc_bet_colokbebas,
			2_win as win_bet_colokbebas, 2_limitotal as limittotal_bet_colokbebas,
			3_minbet as min_bet_colokmacau,
			3_maxbet as max_bet_colokmacau,
			3_disc as disc_bet_colokmacau,
			3_win2digit as win_bet_colokmacau,
			3_win3digit as win3_bet_colokmacau,
			3_win4digit as win4_bet_colokmacau, 3_limittotal as limittotal_bet_colokmacau,
			4_minbet as min_bet_coloknaga,
			4_maxbet as max_bet_coloknaga,
			4_disc as disc_bet_coloknaga,
			4_win3digit as win_bet_coloknaga,
			4_win4digit as win4_bet_coloknaga, 4_limittotal as limittotal_bet_coloknaga,
			5_minbet as min_bet_colokjitu,
			5_maxbet as max_bet_colokjitu,
			5_desic as disc_bet_colokjitu,
			5_winas as winas_bet_colokjitu,
			5_winkop as winkop_bet_colokjitu,
			5_winkepala as winkepala_bet_colokjitu,
			5_winekor as winekor_bet_colokjitu, 5_limitotal as limittotal_bet_colokjitu
			FROM tbl_mst_company_game_pasaran
			WHERE idcompany = ?
			AND idpasarantogel = ?
		`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet_colokbebas, max_bet_colokbebas, disc_bet_colokbebas, win_bet_colokbebas, limittotal_bet_colokbebas                                                                    float32
			min_bet_colokmacau, max_bet_colokmacau, disc_bet_colokmacau, win_bet_colokmacau, win3_bet_colokmacau, win4_bet_colokmacau, limittotal_bet_colokmacau                          float32
			min_bet_coloknaga, max_bet_coloknaga, disc_bet_coloknaga, win_bet_coloknaga, win4_bet_coloknaga, limittotal_bet_coloknaga                                                     float32
			min_bet_colokjitu, max_bet_colokjitu, disc_bet_colokjitu, winas_bet_colokjitu, winkop_bet_colokjitu, winkepala_bet_colokjitu, winekor_bet_colokjitu, limittotal_bet_colokjitu float32
		)

		err = rowresult.Scan(
			&min_bet_colokbebas, &max_bet_colokbebas, &disc_bet_colokbebas, &win_bet_colokbebas, &limittotal_bet_colokbebas,
			&min_bet_colokmacau, &max_bet_colokmacau, &disc_bet_colokmacau, &win_bet_colokmacau, &win3_bet_colokmacau, &win4_bet_colokmacau, &limittotal_bet_colokmacau,
			&min_bet_coloknaga, &max_bet_coloknaga, &disc_bet_coloknaga, &win_bet_coloknaga, &win4_bet_coloknaga, &limittotal_bet_coloknaga,
			&min_bet_colokjitu, &max_bet_colokjitu, &disc_bet_colokjitu, &winas_bet_colokjitu, &winkop_bet_colokjitu,
			&winkepala_bet_colokjitu, &winekor_bet_colokjitu, &limittotal_bet_colokjitu)
		if err != nil {
			return res, err
		}
		obj.Min_bet_colokbebas = min_bet_colokbebas
		obj.Max_bet_colokbebas = max_bet_colokbebas
		obj.Disc_bet_colokbebas = disc_bet_colokbebas
		obj.Win_bet_colokbebas = win_bet_colokbebas
		obj.Limittotal_bet_colokbebas = limittotal_bet_colokbebas
		obj.Min_bet_colokmacau = min_bet_colokmacau
		obj.Max_bet_colokmacau = max_bet_colokmacau
		obj.Disc_bet_colokmacau = disc_bet_colokmacau
		obj.Win_bet_colokmacau = win_bet_colokmacau
		obj.Win3_bet_colokmacau = win3_bet_colokmacau
		obj.Win4_bet_colokmacau = win4_bet_colokmacau
		obj.Limittotal_bet_colokmacau = limittotal_bet_colokmacau
		obj.Min_bet_coloknaga = min_bet_coloknaga
		obj.Max_bet_coloknaga = max_bet_coloknaga
		obj.Disc_bet_coloknaga = disc_bet_coloknaga
		obj.Win_bet_coloknaga = win_bet_coloknaga
		obj.Win4_bet_coloknaga = win4_bet_coloknaga
		obj.Limittotal_bet_coloknaga = limittotal_bet_coloknaga
		obj.Min_bet_colokjitu = min_bet_colokjitu
		obj.Max_bet_colokjitu = max_bet_colokjitu
		obj.Disc_bet_colokjitu = disc_bet_colokjitu
		obj.Winas_bet_colokjitu = winas_bet_colokjitu
		obj.Winkop_bet_colokjitu = winkop_bet_colokjitu
		obj.Winkepala_bet_colokjitu = winkepala_bet_colokjitu
		obj.Winekor_bet_colokjitu = winekor_bet_colokjitu
		obj.Limittotal_bet_colokjitu = limittotal_bet_colokjitu

		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func FetchAll_MinitPasaran5050(client_company, pasaran_code string) (Response, error) {
	var obj Mpasarantogel5050
	var arraobj []Mpasarantogel5050
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT
			6_minbet as min_bet_5050umum,
			6_maxbet as max_bet_5050umum,
			6_keibesar as keibesar_bet_5050umum,
			6_keikecil as keikecil_bet_5050umum, 
			6_keigenap as keigenap_bet_5050umum,
			6_keiganjil as keiganjil_bet_5050umum,
			6_keitengah as keitengah_bet_5050umum,
			6_keitepi as keitepi_bet_5050umum,
			6_discbesar as discbesar_bet_5050umum,
			6_disckecil as disckecil_bet_5050umum,
			6_discgenap as discgenap_bet_5050umum,
			6_discganjil as discganjil_bet_5050umum,
			6_disctengah as disctengah_bet_5050umum,
			6_disctepi as disctepi_bet_5050umum,
			6_limittotal as limittotal_bet_5050umum,
			7_minbet as min_bet_5050special,
			7_maxbet as max_bet_5050special,
			7_keiasganjil as keiasganjil_bet_5050special,
			7_keiasgenap as keiasgenap_bet_5050special,
			7_keiasbesar as keiasbesar_bet_5050special,
			7_keiaskecil as keiaskecil_bet_5050special, 
			7_keikopganjil as keikopganjil_bet_5050special,
			7_keikopgenap as keikopgenap_bet_5050special,
			7_keikopbesar as keikopbesar_bet_5050special,
			7_keikopkecil as keikopkecil_bet_5050special,
			7_keikepalaganjil as keikepalaganjil_bet_5050special,
			7_keikepalagenap as keikepalagenap_bet_5050special, 
			7_keikepalabesar as keikepalabesar_bet_5050special,
			7_keikepalakecil as keikepalakecil_bet_5050special,
			7_keiekorganjil as keiekorganjil_bet_5050special,
			7_keiekorgenap as keiekorgenap_bet_5050special,
			7_keiekorbesar as keiekorbesar_bet_5050special,
			7_keiekorkecil as keiekorkecil_bet_5050special,
			7_discasganjil as discasganjil_bet_5050special,
			7_discasgenap as discasgenap_bet_5050special, 
			7_discasbesar as discasbesar_bet_5050special,
			7_discaskecil as discaskecil_bet_5050special,
			7_disckopganjil as disckopganjil_bet_5050special,
			7_disckopgenap as disckopgenap_bet_5050special,
			7_disckopbesar as disckopbesar_bet_5050special,
			7_disckopkecil as disckopkecil_bet_5050special,
			7_disckepalaganjil as disckepalaganjil_bet_5050special,
			7_disckepalagenap as disckepalagenap_bet_5050special,
			7_disckepalabesar as disckepalabesar_bet_5050special,
			7_disckepalakecil as disckepalakecil_bet_5050special,
			7_discekorganjil as discekorganjil_bet_5050special,
			7_discekorgenap as discekorgenap_bet_5050special,
			7_discekorbesar as discekorbesar_bet_5050special,
			7_discekorkecil as discekorkecil_bet_5050special,
			7_limittotal as limittotal_bet_5050special,
			8_minbet as min_bet_5050kombinasi,
			8_maxbet as max_bet_5050kombinasi,
			8_belakangkeimono as kei_belakangmono_bet_5050kombinasi,
			8_belakangkeistereo as kei_belakangstereo_bet_5050kombinasi,
			8_belakangkeikembang as kei_belakangkembang_bet_5050kombinasi,
			8_belakangkeikempis as kei_belakangkempis_bet_5050kombinasi,
			8_belakangkeikembar as kei_belakangkembar_bet_5050kombinasi,
			8_tengahkeimono as kei_tengahmono_bet_5050kombinasi,
			8_tengahkeistereo as kei_tengahstereo_bet_5050kombinasi,
			8_tengahkeikembang as kei_tengahkembang_bet_5050kombinasi,
			8_tengahkeikempis as kei_tengahkempis_bet_5050kombinasi,
			8_tengahkeikembar as kei_tengahkembar_bet_5050kombinasi,
			8_depankeimono as kei_depanmono_bet_5050kombinasi,
			8_depankeistereo as kei_depanstereo_bet_5050kombinasi,
			8_depankeikembang as kei_depankembang_bet_5050kombinasi,
			8_depankeikempis as kei_depankempis_bet_5050kombinasi,
			8_depankeikembar as kei_depankembar_bet_5050kombinasi,
			8_belakangdiscmono as disc_belakangmono_bet_5050kombinasi,
			8_belakangdiscstereo as disc_belakangstereo_bet_5050kombinasi,
			8_belakangdisckembang as disc_belakangkembang_bet_5050kombinasi,
			8_belakangdisckempis as disc_belakangkempis_bet_5050kombinasi,
			8_belakangdisckembar as disc_belakangkembar_bet_5050kombinasi,
			8_tengahdiscmono as disc_tengahmono_bet_5050kombinasi,
			8_tengahdiscstereo as disc_tengahstereo_bet_5050kombinasi,
			8_tengahdisckembang as disc_tengahkembang_bet_5050kombinasi,
			8_tengahdisckempis as disc_tengahkempis_bet_5050kombinasi,
			8_tengahdisckembar as disc_tengahkembar_bet_5050kombinasi,
			8_depandiscmono as disc_depanmono_bet_5050kombinasi,
			8_depandiscstereo as disc_depanstereo_bet_5050kombinasi,
			8_depandisckembang as disc_depankembang_bet_5050kombinasi,
			8_depandisckempis as disc_depankempis_bet_5050kombinasi,
			8_depandisckembar as disc_depankembar_bet_5050kombinasi,
			8_limittotal as limittotal_bet_5050kombinasi
			FROM tbl_mst_company_game_pasaran
			WHERE idcompany = ?
			AND idpasarantogel = ?
		`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet_5050umum, max_bet_5050umum                                                                                                                                                                              float32
			keibesar_bet_5050umum, keikecil_bet_5050umum, keigenap_bet_5050umum, keiganjil_bet_5050umum, keitengah_bet_5050umum, keitepi_bet_5050umum                                                                       float32
			discbesar_bet_5050umum, disckecil_bet_5050umum, discgenap_bet_5050umum, discganjil_bet_5050umum, disctengah_bet_5050umum, disctepi_bet_5050umum, limittotal_bet_5050umum                                        float32
			min_bet_5050special, max_bet_5050special                                                                                                                                                                        float32
			keiasganjil_bet_5050special, keiasgenap_bet_5050special, keiasbesar_bet_5050special, keiaskecil_bet_5050special                                                                                                 float32
			keikopganjil_bet_5050special, keikopgenap_bet_5050special, keikopbesar_bet_5050special, keikopkecil_bet_5050special                                                                                             float32
			keikepalaganjil_bet_5050special, keikepalagenap_bet_5050special, keikepalabesar_bet_5050special, keikepalakecil_bet_5050special                                                                                 float32
			keiekorganjil_bet_5050special, keiekorgenap_bet_5050special, keiekorbesar_bet_5050special, keiekorkecil_bet_5050special                                                                                         float32
			discasganjil_bet_5050special, discasgenap_bet_5050special, discasbesar_bet_5050special, discaskecil_bet_5050special                                                                                             float32
			disckopganjil_bet_5050special, disckopgenap_bet_5050special, disckopbesar_bet_5050special, disckopkecil_bet_5050special                                                                                         float32
			disckepalaganjil_bet_5050special, disckepalagenap_bet_5050special, disckepalabesar_bet_5050special, disckepalakecil_bet_5050special                                                                             float32
			discekorganjil_bet_5050special, discekorgenap_bet_5050special, discekorbesar_bet_5050special, discekorkecil_bet_5050special, limittotal_bet_5050special                                                         float32
			min_bet_5050kombinasi, max_bet_5050kombinasi                                                                                                                                                                    float32
			kei_belakangmono_bet_5050kombinasi, kei_belakangstereo_bet_5050kombinasi, kei_belakangkembang_bet_5050kombinasi, kei_belakangkempis_bet_5050kombinasi, kei_belakangkembar_bet_5050kombinasi                     float32
			kei_tengahmono_bet_5050kombinasi, kei_tengahstereo_bet_5050kombinasi, kei_tengahkembang_bet_5050kombinasi, kei_tengahkempis_bet_5050kombinasi, kei_tengahkembar_bet_5050kombinasi                               float32
			kei_depanmono_bet_5050kombinasi, kei_depanstereo_bet_5050kombinasi, kei_depankembang_bet_5050kombinasi, kei_depankempis_bet_5050kombinasi, kei_depankembar_bet_5050kombinasi                                    float32
			disc_belakangmono_bet_5050kombinasi, disc_belakangstereo_bet_5050kombinasi, disc_belakangkembang_bet_5050kombinasi, disc_belakangkempis_bet_5050kombinasi, disc_belakangkembar_bet_5050kombinasi                float32
			disc_tengahmono_bet_5050kombinasi, disc_tengahstereo_bet_5050kombinasi, disc_tengahkembang_bet_5050kombinasi, disc_tengahkempis_bet_5050kombinasi, disc_tengahkembar_bet_5050kombinasi                          float32
			disc_depanmono_bet_5050kombinasi, disc_depanstereo_bet_5050kombinasi, disc_depankembang_bet_5050kombinasi, disc_depankempis_bet_5050kombinasi, disc_depankembar_bet_5050kombinasi, limittotal_bet_5050kombinasi float32
		)

		err = rowresult.Scan(
			&min_bet_5050umum, &max_bet_5050umum,
			&keibesar_bet_5050umum, &keikecil_bet_5050umum, &keigenap_bet_5050umum, &keiganjil_bet_5050umum, &keitengah_bet_5050umum, &keitepi_bet_5050umum,
			&discbesar_bet_5050umum, &disckecil_bet_5050umum, &discgenap_bet_5050umum, &discganjil_bet_5050umum, &disctengah_bet_5050umum, &disctepi_bet_5050umum, &limittotal_bet_5050umum,
			&min_bet_5050special, &max_bet_5050special,
			&keiasganjil_bet_5050special, &keiasgenap_bet_5050special, &keiasbesar_bet_5050special, &keiaskecil_bet_5050special,
			&keikopganjil_bet_5050special, &keikopgenap_bet_5050special, &keikopbesar_bet_5050special, &keikopkecil_bet_5050special,
			&keikepalaganjil_bet_5050special, &keikepalagenap_bet_5050special, &keikepalabesar_bet_5050special, &keikepalakecil_bet_5050special,
			&keiekorganjil_bet_5050special, &keiekorgenap_bet_5050special, &keiekorbesar_bet_5050special, &keiekorkecil_bet_5050special,
			&discasganjil_bet_5050special, &discasgenap_bet_5050special, &discasbesar_bet_5050special, &discaskecil_bet_5050special,
			&disckopganjil_bet_5050special, &disckopgenap_bet_5050special, &disckopbesar_bet_5050special, &disckopkecil_bet_5050special,
			&disckepalaganjil_bet_5050special, &disckepalagenap_bet_5050special, &disckepalabesar_bet_5050special, &disckepalakecil_bet_5050special,
			&discekorganjil_bet_5050special, &discekorgenap_bet_5050special, &discekorbesar_bet_5050special, &discekorkecil_bet_5050special, &limittotal_bet_5050special,
			&min_bet_5050kombinasi, &max_bet_5050kombinasi,
			&kei_belakangmono_bet_5050kombinasi, &kei_belakangstereo_bet_5050kombinasi, &kei_belakangkembang_bet_5050kombinasi, &kei_belakangkempis_bet_5050kombinasi, &kei_belakangkembar_bet_5050kombinasi,
			&kei_tengahmono_bet_5050kombinasi, &kei_tengahstereo_bet_5050kombinasi, &kei_tengahkembang_bet_5050kombinasi, &kei_tengahkempis_bet_5050kombinasi, &kei_tengahkembar_bet_5050kombinasi,
			&kei_depanmono_bet_5050kombinasi, &kei_depanstereo_bet_5050kombinasi, &kei_depankembang_bet_5050kombinasi, &kei_depankempis_bet_5050kombinasi, &kei_depankembar_bet_5050kombinasi,
			&disc_belakangmono_bet_5050kombinasi, &disc_belakangstereo_bet_5050kombinasi, &disc_belakangkembang_bet_5050kombinasi, &disc_belakangkempis_bet_5050kombinasi, &disc_belakangkembar_bet_5050kombinasi,
			&disc_tengahmono_bet_5050kombinasi, &disc_tengahstereo_bet_5050kombinasi, &disc_tengahkembang_bet_5050kombinasi, &disc_tengahkempis_bet_5050kombinasi, &disc_tengahkembar_bet_5050kombinasi,
			&disc_depanmono_bet_5050kombinasi, &disc_depanstereo_bet_5050kombinasi, &disc_depankembang_bet_5050kombinasi, &disc_depankempis_bet_5050kombinasi, &disc_depankembar_bet_5050kombinasi,
			&limittotal_bet_5050kombinasi)
		if err != nil {
			return res, err
		}
		obj.Min_bet_5050umum = min_bet_5050umum
		obj.Max_bet_5050umum = max_bet_5050umum
		obj.Keibesar_bet_5050umum = keibesar_bet_5050umum
		obj.Keikecil_bet_5050umum = keikecil_bet_5050umum
		obj.Keigenap_bet_5050umum = keigenap_bet_5050umum
		obj.Keiganjil_bet_5050umum = keiganjil_bet_5050umum
		obj.Keitengah_bet_5050umum = keitengah_bet_5050umum
		obj.Keitepi_bet_5050umum = keitepi_bet_5050umum
		obj.Discbesar_bet_5050umum = discbesar_bet_5050umum
		obj.Disckecil_bet_5050umum = disckecil_bet_5050umum
		obj.Discgenap_bet_5050umum = discgenap_bet_5050umum
		obj.Discganjil_bet_5050umum = discganjil_bet_5050umum
		obj.Disctengah_bet_5050umum = disctengah_bet_5050umum
		obj.Disctepi_bet_5050umum = disctepi_bet_5050umum
		obj.Limittotal_bet_5050umum = limittotal_bet_5050umum

		obj.Min_bet_5050special = min_bet_5050special
		obj.Max_bet_5050special = max_bet_5050special
		obj.Keiasganjil_bet_5050special = keiasganjil_bet_5050special
		obj.Keiasgenap_bet_5050special = keiasgenap_bet_5050special
		obj.Keiasbesar_bet_5050special = keiasbesar_bet_5050special
		obj.Keiaskecil_bet_5050special = keiaskecil_bet_5050special
		obj.Keikopganjil_bet_5050special = keikopganjil_bet_5050special
		obj.Keikopgenap_bet_5050special = keikopgenap_bet_5050special
		obj.Keikopbesar_bet_5050special = keikopbesar_bet_5050special
		obj.Keikopkecil_bet_5050special = keikopkecil_bet_5050special
		obj.Keikepalaganjil_bet_5050special = keikepalaganjil_bet_5050special
		obj.Keikepalagenap_bet_5050special = keikepalagenap_bet_5050special
		obj.Keikepalabesar_bet_5050special = keikepalabesar_bet_5050special
		obj.Keikepalakecil_bet_5050special = keikepalakecil_bet_5050special
		obj.Keiekorganjil_bet_5050special = keiekorganjil_bet_5050special
		obj.Keiekorgenap_bet_5050special = keiekorgenap_bet_5050special
		obj.Keiekorbesar_bet_5050special = keiekorbesar_bet_5050special
		obj.Keiekorkecil_bet_5050special = keiekorkecil_bet_5050special
		obj.Discasganjil_bet_5050special = discasganjil_bet_5050special
		obj.Discasgenap_bet_5050special = discasgenap_bet_5050special
		obj.Discasbesar_bet_5050special = discasbesar_bet_5050special
		obj.Discaskecil_bet_5050special = discaskecil_bet_5050special
		obj.Disckopganjil_bet_5050special = disckopganjil_bet_5050special
		obj.Disckopgenap_bet_5050special = disckopgenap_bet_5050special
		obj.Disckopbesar_bet_5050special = disckopbesar_bet_5050special
		obj.Disckopkecil_bet_5050special = disckopkecil_bet_5050special
		obj.Disckepalaganjil_bet_5050special = disckepalaganjil_bet_5050special
		obj.Disckepalagenap_bet_5050special = disckepalagenap_bet_5050special
		obj.Disckepalabesar_bet_5050special = disckepalabesar_bet_5050special
		obj.Disckepalakecil_bet_5050special = disckepalakecil_bet_5050special
		obj.Discekorganjil_bet_5050special = discekorganjil_bet_5050special
		obj.Discekorgenap_bet_5050special = discekorgenap_bet_5050special
		obj.Discekorbesar_bet_5050special = discekorbesar_bet_5050special
		obj.Discekorkecil_bet_5050special = discekorkecil_bet_5050special
		obj.Limittotal_bet_5050special = limittotal_bet_5050special
		obj.Min_bet_5050kombinasi = min_bet_5050kombinasi
		obj.Max_bet_5050kombinasi = max_bet_5050kombinasi
		obj.Kei_belakangmono_bet_5050kombinasi = kei_belakangmono_bet_5050kombinasi
		obj.Kei_belakangstereo_bet_5050kombinasi = kei_belakangstereo_bet_5050kombinasi
		obj.Kei_belakangkembang_bet_5050kombinasi = kei_belakangkembang_bet_5050kombinasi
		obj.Kei_belakangkempis_bet_5050kombinasi = kei_belakangkempis_bet_5050kombinasi
		obj.Kei_belakangkembar_bet_5050kombinasi = kei_belakangkembar_bet_5050kombinasi
		obj.Kei_tengahmono_bet_5050kombinasi = kei_tengahmono_bet_5050kombinasi
		obj.Kei_tengahstereo_bet_5050kombinasi = kei_tengahstereo_bet_5050kombinasi
		obj.Kei_tengahkembang_bet_5050kombinasi = kei_tengahkembang_bet_5050kombinasi
		obj.Kei_tengahkempis_bet_5050kombinasi = kei_tengahkempis_bet_5050kombinasi
		obj.Kei_tengahkembar_bet_5050kombinasi = kei_tengahkembar_bet_5050kombinasi
		obj.Kei_depanmono_bet_5050kombinasi = kei_depanmono_bet_5050kombinasi
		obj.Kei_depanstereo_bet_5050kombinasi = kei_depanstereo_bet_5050kombinasi
		obj.Kei_depankembang_bet_5050kombinasi = kei_depankembang_bet_5050kombinasi
		obj.Kei_depankempis_bet_5050kombinasi = kei_depankempis_bet_5050kombinasi
		obj.Kei_depankembar_bet_5050kombinasi = kei_depankembar_bet_5050kombinasi
		obj.Disc_belakangmono_bet_5050kombinasi = disc_belakangmono_bet_5050kombinasi
		obj.Disc_belakangstereo_bet_5050kombinasi = disc_belakangstereo_bet_5050kombinasi
		obj.Disc_belakangkembang_bet_5050kombinasi = disc_belakangkembang_bet_5050kombinasi
		obj.Disc_belakangkempis_bet_5050kombinasi = disc_belakangkempis_bet_5050kombinasi
		obj.Disc_belakangkembar_bet_5050kombinasi = disc_belakangkembar_bet_5050kombinasi
		obj.Disc_tengahmono_bet_5050kombinasi = disc_tengahmono_bet_5050kombinasi
		obj.Disc_tengahstereo_bet_5050kombinasi = disc_tengahstereo_bet_5050kombinasi
		obj.Disc_tengahkembang_bet_5050kombinasi = disc_tengahkembang_bet_5050kombinasi
		obj.Disc_tengahkempis_bet_5050kombinasi = disc_tengahkempis_bet_5050kombinasi
		obj.Disc_tengahkembar_bet_5050kombinasi = disc_tengahkembar_bet_5050kombinasi
		obj.Disc_depanmono_bet_5050kombinasi = disc_depanmono_bet_5050kombinasi
		obj.Disc_depanstereo_bet_5050kombinasi = disc_depanstereo_bet_5050kombinasi
		obj.Disc_depankembang_bet_5050kombinasi = disc_depankembang_bet_5050kombinasi
		obj.Disc_depankempis_bet_5050kombinasi = disc_depankempis_bet_5050kombinasi
		obj.Disc_depankembar_bet_5050kombinasi = disc_depankembar_bet_5050kombinasi
		obj.Limittotal_bet_5050kombinasi = limittotal_bet_5050kombinasi

		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func FetchAll_MinitPasaranMacauKombinasi(client_company, pasaran_code string) (Response, error) {
	var obj MpasarantogelMacauKombinasi
	var arraobj []MpasarantogelMacauKombinasi
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT 
		9_minbet as min_bet, 
		9_maxbet as max_bet, 
		9_win as win_bet, 
		9_discount as diskon_bet, 
		9_limittotal as limit_total 
		FROM tbl_mst_company_game_pasaran 
		WHERE idcompany = ? 
		AND idpasarantogel = ?
	`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet, max_bet, win_bet, diskon_bet, limit_total float32
		)

		err = rowresult.Scan(
			&min_bet, &max_bet, &win_bet, &diskon_bet, &limit_total)
		if err != nil {
			return res, err
		}
		obj.Min_bet = min_bet
		obj.Max_bet = max_bet
		obj.Win_bet = win_bet
		obj.Diskon_bet = diskon_bet
		obj.Limit_total = limit_total
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func FetchAll_MinitPasaranDasar(client_company, pasaran_code string) (Response, error) {
	var obj MpasarantogelDasar
	var arraobj []MpasarantogelDasar
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT 
		10_minbet as min_bet, 
		10_maxbet as max_bet, 
		10_keibesar as kei_besar_bet, 
		10_keikecil as kei_kecil_bet, 
		10_keigenap as kei_genap_bet, 
		10_keiganjil as kei_ganjil_bet, 
		10_discbesar as disc_besar_bet, 
		10_disckecil as disc_kecil_bet, 
		10_discigenap as disc_genap_bet, 
		10_discganjil as disc_ganjil_bet,  
		10_limittotal as limit_total 
		FROM tbl_mst_company_game_pasaran 
		WHERE idcompany = ? 
		AND idpasarantogel = ?
	`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet, max_bet, kei_besar_bet, kei_kecil_bet, kei_genap_bet, kei_ganjil_bet float32
			disc_besar_bet, disc_kecil_bet, disc_genap_bet, disc_ganjil_bet, limit_total  float32
		)

		err = rowresult.Scan(
			&min_bet, &max_bet, &kei_besar_bet, &kei_kecil_bet, &kei_genap_bet,
			&kei_ganjil_bet, &disc_besar_bet, &disc_kecil_bet, &disc_genap_bet, &disc_ganjil_bet,
			&limit_total)
		if err != nil {
			return res, err
		}
		obj.Min_bet = min_bet
		obj.Max_bet = max_bet
		obj.Kei_besar_bet = kei_besar_bet
		obj.Kei_kecil_bet = kei_kecil_bet
		obj.Kei_genap_bet = kei_genap_bet
		obj.Kei_ganjil_bet = kei_ganjil_bet
		obj.Disc_besar_bet = disc_besar_bet
		obj.Disc_kecil_bet = disc_kecil_bet
		obj.Disc_genap_bet = disc_genap_bet
		obj.Disc_ganjil_bet = disc_ganjil_bet
		obj.Limit_total = limit_total
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
func FetchAll_MinitPasaranShio(client_company, pasaran_code string) (Response, error) {
	var obj MpasarantogelShio
	var arraobj []MpasarantogelShio
	var res Response
	msg := "Error"
	con := db.CreateCon()
	tglnow, _ := goment.New()

	sqlresult := `SELECT 
		11_minbet as min_bet, 
		11_maxbet as max_bet, 
		11_win as win_bet, 
		11_disc as diskon_bet, 
		11_limittotal as limit_total 
		FROM tbl_mst_company_game_pasaran 
		WHERE idcompany = ? 
		AND idpasarantogel = ?
	`
	rowresult, err := con.Query(sqlresult, client_company, pasaran_code)
	defer rowresult.Close()

	if err != nil {
		return res, err
	}
	for rowresult.Next() {
		var (
			min_bet, max_bet, win_bet, diskon_bet, limit_total float32
		)

		err = rowresult.Scan(
			&min_bet, &max_bet, &win_bet, &diskon_bet, &limit_total)
		if err != nil {
			return res, err
		}
		obj.Min_bet = min_bet
		obj.Max_bet = max_bet
		obj.Win_bet = win_bet
		obj.Diskon_bet = diskon_bet
		obj.Limit_total = limit_total
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	if len(arraobj) > 0 {
		res.Status = fiber.StatusOK
		res.Message = msg
		res.Totalrecord = len(arraobj)
		res.Record = arraobj
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	} else {
		res.Status = fiber.StatusBadRequest
		res.Message = "Not Found"
		res.Totalrecord = 0
		res.Record = nil
		res.Time = tglnow.Format("YYYY-MM-DD HH:mm:ss")
	}

	return res, nil
}
