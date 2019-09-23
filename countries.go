package usercrm

import (
	"errors"
	"net/http"

	"github.com/memclutter/user-crm/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewCountries(e *echo.Group) *Countries {
	ep := &Countries{}

	e.GET("/countries", ep.List)
	e.POST("/countries", ep.Create)
	e.PUT("/countries/:code", ep.Update)
	e.DELETE("/countries/:code", ep.Remove)

	return ep
}

type CountriesListRequest struct {
	Offset int64 `query:"offset"`
	Limit  int64 `query:"limit"`
}

type CountriesFormRequest struct {
	Code string `json:"code" form:"code"`
	Name string `json:"name" form:"name"`
}

type Countries struct {
}

func (e Countries) List(c echo.Context) error {

	// Parse request
	req := new(CountriesListRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return err
	}

	// Get database instance
	db, ok := c.Get("db").(*gorm.DB)
	if !ok {
		err := errors.New("database instance missing in context")
		c.Logger().Error(err)
		return err
	}

	// Get total count
	totalCount := 0
	if err := db.Model(&models.Country{}).Count(&totalCount).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	// Result
	result := make([]models.Country, 0)

	// Build query
	if err := db.Offset(req.Offset).Limit(req.Limit).Find(&result).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"totalCount": totalCount,
		"items":      result,
	})
}

func (e Countries) Create(c echo.Context) error {

	// Parse request
	req := new(CountriesFormRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return err
	} else if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// TODO validation

	// Get database instance
	db, ok := c.Get("db").(*gorm.DB)
	if !ok {
		err := errors.New("database instance missing in context")
		c.Logger().Error(err)
		return err
	}

	_ = db

	// TODO Insert record
	// TODO return response
	return nil
}

func (e Countries) Update(c echo.Context) error {

	return nil
}

func (e Countries) Remove(c echo.Context) error {

	return nil
}
