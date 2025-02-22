package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/agung6544/be-tbku/config"
	inimodel "github.com/agung6544/packagetbku/model"
	cek "github.com/agung6544/packagetbku/module"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetOrder godoc
// @Summary Get All Data Order.
// @Description Mengambil semua data order.
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /order [get]
func GetOrder(c *fiber.Ctx) error {
	ps := cek.GetAllOrder(config.Ulbimongoconn, "orderku")
	return c.JSON(ps)
}

// GetOrderID godoc
// @Summary Get By ID Data Order.
// @Description Ambil per ID data order.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Order
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /order/{id} [get]
func GetOrderID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetOrderFromID(objID, config.Ulbimongoconn, "orderku")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataOrder godoc
// @Summary Insert data order.
// @Description Input data order.
// @Tags Order
// @Accept json
// @Produce json
// @Param request body ReqOrder true "Payload Body [RAW]"
// @Success 200 {object} Order
// @Failure 400
// @Failure 500
// @Router /insorder [post]
func InsertDataOrder(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var order inimodel.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertOrder(db, "orderku",
		order.Ayam,
		order.Nama_Pemesan,
		order.Alamat,)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateDataOrder godoc
// @Summary Update data order.
// @Description Ubah data order.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqOrder true "Payload Body [RAW]"
// @Success 200 {object} Order
// @Failure 400
// @Failure 500
// @Router /uporder/{id} [put]
func UpdateDataOrder(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var order inimodel.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateOrder function with the parsed ID and the Ayam object
	err = cek.UpdateOrder(db, "orderku",
		objectID,
		order.Ayam,
		order.Nama_Pemesan,
		order.Alamat,)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteOrderByID godoc
// @Summary Delete data order.
// @Description Hapus data order.
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delorder/{id} [delete]
func DeleteOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteOrderByID(objID, config.Ulbimongoconn, "orderku")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}