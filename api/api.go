package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sbusso/ipdata/ipdata"
)

// Start starts the ipdata API service using the passed params
func Start(port string) {
	// Create the ipdata client
	ic, err := ipdata.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer ic.Close()

	// Create a new echo Echo and bind middleware
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Bind all API endpoint handlers
	e.GET("/lookup/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Lookup(c.Param("ip")))
	})

	e.GET("/continent/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Continent(c.Param("ip")))
	})

	// Listen on the passed port
	e.Logger.Fatal(e.Start(":" + port))
}
