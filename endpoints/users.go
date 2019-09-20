package endpoints

import "github.com/labstack/echo"

func NewUsers(e *echo.Echo) *Users {
	ep := &Users{}

	e.GET("/users", ep.List)
	e.POST("/users", ep.Create)
	e.PUT("/users/:id", ep.Update)
	e.DELETE("/users/:id", ep.Remove)

	return ep
}

type Users struct {
}

func (e Users) List(c echo.Context) error {

	return nil
}

func (e Users) Create(c echo.Context) error {

	return nil
}

func (e Users) Update(c echo.Context) error {

	return nil
}

func (e Users) Remove(c echo.Context) error {

	return nil
}
