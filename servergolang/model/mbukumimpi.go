package model

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiberapi/config"
	"github.com/nikitamirzani323/gofiberapi/db"
	"github.com/nikitamirzani323/gofiberapi/helpers"
)

type Mbukumimpi struct {
	Tipe  string `json:"bukumimpi_tipe"`
	Nama  string `json:"bukumimpi_nama"`
	Nomor string `json:"bukumimpi_nomor"`
}

func FetchAll_Mbukumimpi(tipe, nama string) (helpers.Response, error) {
	var obj Mbukumimpi
	var arraobj []Mbukumimpi
	var res helpers.Response
	render_page := time.Now()
	ctx := context.Background()
	con := db.CreateCon()
	sql_bukumimpi := ""

	sql_bukumimpi += ""
	sql_bukumimpi += "SELECT "
	sql_bukumimpi += "typebukumimpi, nmbukumimpi, nmrbukumimpi "
	sql_bukumimpi += "FROM " + config.DB_tbl_bukumimpi + " "
	if tipe != "" {
		sql_bukumimpi += "WHERE typebukumimpi='" + tipe + "' "
	}
	if nama != "" {
		sql_bukumimpi += "WHERE nmbukumimpi LIKE '%" + nama + "%' "
	}
	sql_bukumimpi += "ORDER BY nmbukumimpi ASC LIMIT 50 "

	rows, err := con.QueryContext(ctx, sql_bukumimpi)
	defer rows.Close()
	helpers.ErrorCheck(err)
	for rows.Next() {
		var (
			typebukumimpi_db, nmbukumimpi_db, nmrbukumimpi_db string
		)
		err = rows.Scan(&typebukumimpi_db, &nmbukumimpi_db, &nmrbukumimpi_db)
		helpers.ErrorCheck(err)
		obj.Tipe = typebukumimpi_db
		obj.Nama = nmbukumimpi_db
		obj.Nomor = nmrbukumimpi_db
		arraobj = append(arraobj, obj)
	}
	res.Status = fiber.StatusOK
	res.Message = "Success"
	res.Totalrecord = len(arraobj)
	res.Record = arraobj
	res.Time = time.Since(render_page).String()

	return res, nil
}
