package usercrm

import (
	"context"
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
	Code string `json:"code" form:"code" validate:"required,country_code"`
	Name string `json:"name" form:"name" validate:"required"`
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

	// Get custom validator instance
	cv, ok := c.Get("validator").(*CustomValidator)
	if !ok {
		err := errors.New("validator instance missing in context")
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

	// Create validation context
	ctx := context.WithValue(c.Request().Context(), "db", db.Model(&models.Country{}))

	// Parse request
	req := new(CountriesFormRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return err
	} else if ok, errs := cv.ValidateCtx(req, ctx); !ok {
		return c.JSON(http.StatusBadRequest, errs)
	}

	record := &models.Country{
		Code: req.Code,
		Name: req.Name,
	}

	if err := db.Create(record).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusCreated, record)
}

func (e Countries) Update(c echo.Context) error {

	// Parse country code from path
	code := c.Param("code")

	// Get custom validator instance
	cv, ok := c.Get("validator").(*CustomValidator)
	if !ok {
		err := errors.New("validator instance missing in context")
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

	// Find record
	record := &models.Country{}
	if err := db.Model(&record).Where("code = ?", code).First(record).Error; err == gorm.ErrRecordNotFound {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		c.Logger().Error(err)
		return err
	}

	// Create validation context
	ctx := context.WithValue(c.Request().Context(), "db", db.Model(&models.Country{}).Where("code != ?", code))

	// Parse request
	req := new(CountriesFormRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return err
	} else if ok, errs := cv.ValidateCtx(req, ctx); !ok {
		return c.JSON(http.StatusBadRequest, errs)
	}

	record.Code = req.Code
	record.Name = req.Name

	if err := db.Save(record).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, record)
}

func (e Countries) Remove(c echo.Context) error {

	// Parse country code from path
	code := c.Param("code")

	// Get database instance
	db, ok := c.Get("db").(*gorm.DB)
	if !ok {
		err := errors.New("database instance missing in context")
		c.Logger().Error(err)
		return err
	}

	// Find record
	record := &models.Country{}
	if err := db.Model(record).Where("code = ?", code).First(record).Error; err == gorm.ErrRecordNotFound {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		c.Logger().Error(err)
		return err
	}

	// Remove record
	if err := db.Delete(record).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
