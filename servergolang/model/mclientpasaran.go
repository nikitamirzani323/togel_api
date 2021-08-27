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
