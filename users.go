package usercrm

import (
	"errors"
	"net/http"
	"time"

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

type UsersFormRequest struct {
	Username    string `json:"username" form:"username" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Birthday    string `json:"birthday" form:"birthday"`
	CountryCode string `json:"countryCode" form:"countryCode"`
	Gender      string `json:"gender" form:"gender"`
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

	// Parse request
	req := new(UsersFormRequest)
	if err := c.Bind(req); err != nil {
		c.Logger().Error(err)
		return err
	} else if ok, errs := cv.ValidateCtx(req, c.Request().Context()); !ok {
		// TODO custom validation for unique username, email and validate countryCode
		return c.JSON(http.StatusBadRequest, errs)
	}

	// Birthday parse in time.Time
	birthday, err := time.Parse("2006-01-02", req.Birthday)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"birthday": "DATE",
		})
	}

	// Create user record
	record := &models.User{
		Username:    req.Username,
		Email:       req.Email,
		Birthday:    birthday,
		CountryCode: req.CountryCode,
		Gender:      req.Gender,
	}

	if err := db.Create(record).Error; err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusCreated, record)
}

func (e Users) Update(c echo.Context) error {

	return nil
}

func (e Users) Remove(c echo.Context) error {

	return nil
}
