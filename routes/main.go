package routes

import (
	"github.com/labstack/echo/v4"
)

func Load(app *echo.Echo) {
	version1 := app.Group("/v1")
	CrmContact(version1.Group("/contact"))
}
