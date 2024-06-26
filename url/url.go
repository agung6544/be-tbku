package url

import (
	"github.com/agung6544/be-tbku/controller"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/swagger" // swagger handler
)

func Web(page *fiber.App) {
	page.Get("/ayam", controller.GetAyam)
	page.Get("/ayam/:id", controller.GetAyamID) //menampilkan data presensi berdasarkan id

	page.Post("/ins", controller.InsertDataAyam)
	page.Put("/update/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeleteAyamByID)
	// page.Get("/docs/*", swagger.HandlerDefault)
}