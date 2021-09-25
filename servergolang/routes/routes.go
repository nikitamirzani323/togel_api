package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/gofiberapi/controller"
)

const idleTimeout = 5 * time.Second

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Post("/api/servicetoken", controller.Fetch_token)
	app.Post("/api/serviceresult", controller.FetchAll_result)
	app.Post("/api/servicecheckpasaran", controller.Fetch_CheckPasaran)
	app.Post("/api/serviceinit", controller.FetchAll_pasaran)
	app.Post("/api/serviceconfigtogel", controller.Fetch_InitPasaran)
	app.Post("/api/servicelimittogel", controller.Fetch_LimitPasaran432)
	app.Post("/api/serviceinvoicebet", controller.Fetch_listinvoicebet)
	app.Post("/api/serviceslip", controller.Fetch_slipperiode)
	app.Post("/api/serviceslipdetail", controller.Fetch_slipperiodedetail)
	app.Post("/api/savetransaksi", controller.SaveTogel)

	app.Post("/api/bukumimpi", controller.FetchAll_bukumimpi)

	app.Post("/api/deleteredispasaran", controller.AdminDell_pasaran)
	app.Post("/api/deleteredisresult", controller.AdminDell_result)
	return app
}
