package endpoints

import "github.com/labstack/echo"

func NewCountries(e *echo.Echo) *Countries {
	ep := &Countries{}

	e.GET("/countries", ep.List)
	e.POST("/countries", ep.Create)
	e.PUT("/countries/:code", ep.Update)
	e.DELETE("/countries/:code", ep.Remove)

	return ep
}

type Countries struct {
}

func (e Countries) List(c echo.Context) error {

	return nil
}

func (e Countries) Create(c echo.Context) error {

	return nil
}

func (e Countries) Update(c echo.Context) error {

	return nil
}

func (e Countries) Remove(c echo.Context) error {

	return nil
}
