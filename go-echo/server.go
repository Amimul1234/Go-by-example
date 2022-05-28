package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {

	e := echo.New()

	e.GET("/", getHelloWorld)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.POST("/save", save)

	defer e.Logger.Fatal(e.Start(":1323"))
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "Name: "+name+", email: "+email)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "Team: "+team+", member: "+member)
}

func getHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
