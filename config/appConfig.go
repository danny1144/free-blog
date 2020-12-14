package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"os"
)

var engine = html.New("./views", ".html")
var AppConfig = fiber.Config{
	Prefork:       false,
	StrictRouting: true,
	ServerHeader:  "Fiber",
	CaseSensitive: true,
	Views:         engine,
}

var LoggerConfig = logger.Config{
	//Format:     "${pid} ${status} - ${method} ${path}\n",
	Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	TimeFormat: "02-Jan-2006",
	TimeZone:   "Asia/Chongqing",
	Output:     os.Stdout,
}
