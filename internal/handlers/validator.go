package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
)

type Validate struct {
	validate *validator.Validate
}

func GetValidator() (*Validate, error) {
	return &Validate{validate: validator.New()}, nil
}

func (v *Validate) InitValidator() error {

	if err := v.validate.RegisterValidation("slug", SlugValidator); err != nil {
		return errors.Wrap(err, "slug validator error")
	}
	if err := v.validate.RegisterValidation("chance", ChanceValidator); err != nil {
		return errors.Wrap(err, "chance validator error")
	}
	if err := v.validate.RegisterValidation("datem", DateMValidator); err != nil {
		return errors.Wrap(err, "datem validator error")
	}
	if err := v.validate.RegisterValidation("slugslice", SlugSliceValidator); err != nil {
		return errors.Wrap(err, "slugslice validator error")
	}
	if err := v.validate.RegisterValidation("ttl", TtlValidator); err != nil {
		return errors.Wrap(err, "ttl validator error")
	}
	return nil
}

func SlugValidator(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(string)
	if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", value); !matched || err != nil {
		return false
	}
	return true
}

func ChanceValidator(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(string)
	if matched, err := regexp.MatchString("([0-1])|([.][0-9]+)", value); !matched || err != nil {
		return false
	}
	valueF, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	if valueF > 1 {
		return false
	}
	return true
}

// (0[1-9]|1[1,2])(\/|-)(19|20)\d{2}
func DateMValidator(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(string)
	if matched, err := regexp.MatchString("(19|20)\\d{2}(\\/|-)(0[1-9]|1[1,2])", value); !matched || err != nil {
		return false
	}
	return true
}

func SlugSliceValidator(fl validator.FieldLevel) bool {
	slice := fl.Field().Interface().([]string)
	if len(slice) != 0 {
		for _, v := range slice {
			if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", v); !matched || err != nil {
				return false
			}
		}
	}
	return true
}

func TtlValidator(fl validator.FieldLevel) bool {
	ttl := fl.Field().Interface().(map[string]string)
	for key, value := range ttl {
		if matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", key); !matched || err != nil {
			return false
		}
		if matched, err := regexp.MatchString("^[a-zA-z0-9\\s]+$", value); !matched || err != nil {
			return false
		}
	}
	return true
}
