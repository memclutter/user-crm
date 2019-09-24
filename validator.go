package usercrm

import (
	"context"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/memclutter/user-crm/models"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() (*CustomValidator, error) {
	v := validator.New()

	// Use json tag value as validation errors key
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	// Country code custom validator
	if err := v.RegisterValidationCtx("country_code", func(ctx context.Context, fl validator.FieldLevel) bool {
		db := ctx.Value("db").(*gorm.DB)
		countryCode := fl.Field().String()
		totalCount := 0
		if err := db.Model(&models.Country{}).Where("code = ?", countryCode).Count(&totalCount).Error; err != nil {
			// TODO handle validate database error
			panic(err)
			return false
		}

		// Validation pass if code not exist
		return totalCount == 0
	}); err != nil {
		return nil, err
	}

	return &CustomValidator{validator: v}, nil
}

func (cv CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (cv CustomValidator) ValidateCtx(i interface{}, ctx context.Context) (bool, map[string]string) {
	vErr := make(map[string]string)

	if err := cv.validator.StructCtx(ctx, i); err != nil {
		e := err.(validator.ValidationErrors)

		for _, fe := range e {

			// Generate error code
			code := strings.ToUpper(fe.Tag())
			key := fe.Field()

			// Assign error code
			vErr[key] = code
		}
	}

	return len(vErr) == 0, vErr
}
