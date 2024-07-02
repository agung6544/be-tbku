package url

import (
	"github.com/agung6544/be-tbku/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	page.Get("/ayam", controller.GetAyam)
	page.Get("/ayam/:id", controller.GetAyamID)
	page.Post("/ins", controller.InsertDataAyam)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteAyamByID)

	page.Get("/order", controller.GetOrder)
	page.Get("/order/:id", controller.GetOrderID)
	page.Post("/insorder", controller.InsertDataOrder)
	page.Put("/uporder/:id", controller.UpdateDataOrder)
	page.Delete("/delorder/:id", controller.DeleteOrderByID)
	page.Get("/docs/*", swagger.HandlerDefault)
}