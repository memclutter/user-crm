package usercrm

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/user-crm/models"
)

func NewUsers(e *echo.Group) *Users {
	ep := &Users{}

	e.GET("/users", ep.List)
	e.POST("/users", ep.Create)
	e.PUT("/users/:id", ep.Update)
	e.DELETE("/users/:id", ep.Remove)

	return ep
}

type UsersListRequest struct {
	CountryCode string `query:"countryCode"`
	AgeRange    []int  `query:"ageRange[]"`
	Gender      string `query:"gender"`
	Offset      int64  `query:"offset"`
	Limit       int64  `query:"limit"`
}

type Users struct {
}

func (e Users) List(c echo.Context) error {

	// Parse request
	req := new(UsersListRequest)
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

	// Build db query
	query := db.Model(&models.User{})

	if len(req.CountryCode) > 0 {
		query = query.Where("country_code = ?", req.CountryCode)
	}

	if len(req.AgeRange) == 2 {
		// TODO filter by age
	}

	if len(req.Gender) > 0 {
		query = query.Where("gender = ?", req.Gender)
	}

	// Get total count
	totalCount := 0
	if err := query.Count(&totalCount).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	// Result
	result := make([]models.User, 0)

	// Build query
	if err := query.Offset(req.Offset).Limit(req.Limit).Find(&result).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"totalCount": totalCount,
		"items":      result,
	})
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
