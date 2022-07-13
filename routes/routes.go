package routes

import (
	"time"

	"bitbucket.org/isbtotogroup/api_go/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const idleTimeout = 5 * time.Second

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Path() == "/"
		},
		Format: "${time} | ${status} | ${latency} | ${ips[0]} | ${method} | ${path} - ${queryParams} ${body}\n",
	}))
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        200,
		Expiration: 200 * time.Second,
	}))
	app.Use(etag.New())
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))
	app.Use(func(c *fiber.Ctx) error {

		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		return c.Next()
	})

	app.Post("/api/servicetoken", controller.Fetch_token)
	app.Post("/api/serviceresult", controller.FetchAll_resultbypasaran)
	app.Post("/api/serviceresultall", controller.FetchAll_result)
	app.Post("/api/servicecheckpasaran", controller.Fetch_CheckPasaran)
	app.Post("/api/serviceinit", controller.FetchAll_pasaran)
	app.Post("/api/serviceconfigtogel", controller.Fetch_InitPasaran)
	app.Post("/api/servicelimittogel", controller.Fetch_LimitPasaran432)
	app.Post("/api/serviceinvoicebet", controller.Fetch_listinvoicebet)
	app.Post("/api/serviceinvoicebetdetail", controller.Fetch_listinvoicebetid)
	app.Post("/api/serviceslip", controller.Fetch_slipperiode)
	app.Post("/api/serviceslipall", controller.Fetch_slipperiodeall)
	app.Post("/api/serviceslipdetail", controller.Fetch_slipperiodedetail)

	app.Post("/api/bukumimpi", controller.FetchAll_bukumimpi)
	return app
}
