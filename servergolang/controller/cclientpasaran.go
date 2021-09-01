package controller

import (
	"context"
	"encoding/json"
	"log"

	"github.com/buger/jsonparser"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiberapi/config"
	"github.com/nikitamirzani323/gofiberapi/model"
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
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Pasaran_Code    string `json:"pasaran_code"`
	Pasaran_Periode string `json:"pasaran_periode"`
}
type ClientSlipPeriode struct {
	Client_Username string `json:"client_username"`
	Client_Company  string `json:"client_company"`
	Pasaran_Code    string `json:"pasaran_code"`
}
type ClientSaveTogel struct {
	Client_Username string      `json:"client_username"`
	Client_Company  string      `json:"client_company"`
	Idtrxkeluaran   string      `json:"idtrxkeluaran"`
	Idcomppasaran   string      `json:"idcomppasaran"`
	Devicemember    string      `json:"devicemember"`
	Formipaddress   string      `json:"formipaddress"`
	Timezone        string      `json:"timezone"`
	Totalbayarbet   int         `json:"totalbayarbet"`
	List4d          interface{} `json:"list4d"`
}
type parsingjson struct {
	Record []ytRecord `json:"record"`
}
type ytRecord struct {
	PasaranId      string `json:"pasaran_id"`
	PasaranTogel   string `json:"pasaran_togel"`
	PasaranPeriode string `json:"pasaran_periode"`
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
	return c.JSON(fiber.Map{
		"status":          fiber.StatusOK,
		"token":           client.Token,
		"member_username": "developer",
		"member_company":  "MMD",
		"member_credit":   5000000,
	})
}
func FetchAll_pasaran(c *fiber.Ctx) error {
	client := new(ClientInit)

	if err := c.BodyParser(client); err != nil {
		return err
	}
	conf := config.GetConfigRedis()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.DB_HOST,
		Password: conf.DB_PASSWORD,
		DB:       conf.DB_NAME,
	})

	// rdb.Del(ctx, "listpasaran_"+client.Client_Company)
	resultredis, err := rdb.Get(ctx, "listpasaran_"+client.Client_Company).Result()
	if err == redis.Nil {
		result, err := model.FetchAll_MclientPasaran(client.Client_Company)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		// json, _ := json.Marshal(result)
		// log.Println("mysql")
		// err = rdb.Set(ctx, "listpasaran_"+client.Client_Company, json, 0).Err()
		// if err != nil {
		// 	panic(err)
		// }
		return c.JSON(result)
	} else {
		data := []byte(resultredis)
		temp, _, _, _ := jsonparser.Get(data, "record")

		jsonparser.ArrayEach(temp, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			var1, _, _, _ := jsonparser.Get(value, "pasaran_id")
			var2, _, _, _ := jsonparser.Get(value, "pasaran_togel")
			var3, _, _, _ := jsonparser.Get(value, "pasaran_periode")
			var4, _, _, _ := jsonparser.Get(value, "pasaran_tglkeluaran")
			log.Printf("%s - %s - %s - %s\n", string(var1), string(var2), string(var3), string(var4))
		})

		log.Println("cache")
		rdb.Close()
		return c.SendString(resultredis)
	}
}
func AdminDell_pasaran(c *fiber.Ctx) error {
	client := new(ClientInit)

	if err := c.BodyParser(client); err != nil {
		return err
	}
	conf := config.GetConfigRedis()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.DB_HOST,
		Password: conf.DB_PASSWORD,
		DB:       conf.DB_NAME,
	})

	rdb.Del(ctx, "listpasaran_"+client.Client_Company)
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Delete Success",
	})
}
func FetchAll_result(c *fiber.Ctx) error {
	client := new(ClientResult)

	if err := c.BodyParser(client); err != nil {
		return err
	}

	conf := config.GetConfigRedis()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.DB_HOST,
		Password: conf.DB_PASSWORD,
		DB:       conf.DB_NAME,
	})
	resultredis, err := rdb.Get(ctx, "listresult_"+client.Client_Company+"_"+client.Pasaran_Code).Result()
	if err == redis.Nil {
		result, err := model.FetchAll_MclientPasaranResult(client.Client_Company, client.Pasaran_Code)

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
			json, _ := json.Marshal(result)
			err = rdb.Set(ctx, "listresult_"+client.Client_Company+"_"+client.Pasaran_Code, json, 0).Err()
			if err != nil {
				panic(err)
			}
		}
		return c.JSON(result)
	} else {
		log.Println("cache")
		rdb.Close()
		return c.SendString(resultredis)
	}
}
func AdminDell_result(c *fiber.Ctx) error {
	client := new(ClientResult)

	if err := c.BodyParser(client); err != nil {
		return err
	}
	conf := config.GetConfigRedis()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.DB_HOST,
		Password: conf.DB_PASSWORD,
		DB:       conf.DB_NAME,
	})

	rdb.Del(ctx, "listresult_"+client.Client_Company+"_"+client.Pasaran_Code)
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Delete Success",
	})
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
	if err := c.BodyParser(client); err != nil {
		return err
	}

	result, err := model.FetchAll_MinitPasaran(client.Client_Company, client.Pasaran_Code, client.Permainan)

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
	result, err := model.Fetch_invoicebet(client.Client_Username, client.Client_Company, client.Pasaran_Code, client.Pasaran_Periode)
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
func SaveTogel(c *fiber.Ctx) error {
	client := new(ClientSaveTogel)
	if err := c.BodyParser(client); err != nil {
		return err
	}

	result, err := model.Savetransaksi(
		client.Client_Username,
		client.Client_Company, client.Idtrxkeluaran, client.Idcomppasaran, client.Devicemember, client.Formipaddress, client.Timezone, client.Totalbayarbet, client.List4d)
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
