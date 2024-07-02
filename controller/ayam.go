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

// GetAyam godoc
// @Summary Get All Data Ayam.
// @Description Mengambil semua data ayam.
// @Tags Ayam
// @Accept json
// @Produce json
// @Success 200 {object} Ayam
// @Router /ayam [get]
func GetAyam(c *fiber.Ctx) error {
	ps := cek.GetAllAyam(config.Ulbimongoconn, "ayamku")
	return c.JSON(ps)
}

// GetAyamID godoc
// @Summary Get By ID Data Ayam.
// @Description Ambil per ID data ayam.
// @Tags Ayam
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Ayam
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /ayam/{id} [get]
func GetAyamID(c *fiber.Ctx) error {
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
	ps, err := cek.GetAyamFromID(objID, config.Ulbimongoconn, "ayamku")
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

// InsertDataAyam godoc
// @Summary Insert data ayam.
// @Description Input data ayam.
// @Tags Ayam
// @Accept json
// @Produce json
// @Param request body ReqAyam true "Payload Body [RAW]"
// @Success 200 {object} Ayam
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertDataAyam(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var ayam inimodel.Ayam
	if err := c.BodyParser(&ayam); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertAyam(db, "ayamku",
		ayam.Jenis,
		ayam.Umur,
		ayam.Bobot,
		ayam.Tinggi,
		ayam.Jenis_Kelamin,
		ayam.Harga)
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


func UpdateData(c *fiber.Ctx) error {
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
	var ayam inimodel.Ayam
	if err := c.BodyParser(&ayam); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdateAyam function with the parsed ID and the Ayam object
	err = cek.UpdateAyam(db, "ayamku",
		objectID,
		ayam.Jenis,
		ayam.Umur,
		ayam.Bobot,
		ayam.Tinggi,
		ayam.Jenis_Kelamin,
		ayam.Harga)
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


func DeleteAyamByID(c *fiber.Ctx) error {
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

	err = cek.DeleteAyamByID(objID, config.Ulbimongoconn, "ayamku")
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