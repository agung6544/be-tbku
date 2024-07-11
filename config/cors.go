package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://yoginara.github.io",
	"http://127.0.0.1:5500",
	"http://127.0.0.1:5501",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowHeaders:     "Origin,Login,Content-Type",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}