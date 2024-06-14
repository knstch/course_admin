package validation

import (
	"context"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	adminerror "github.com/knstch/course_admin/internal/app/admin_error"
)

const (
	errFieldAcceptsOnlyLetters = "допустимы только буквы"
	errBadBool                 = "допустимы значения только true/fasle"
	errBadEmail                = "email передан неправильно"
)

var (
	lettersRegex = regexp.MustCompile(`^\p{L}+$`)
	phoneRegex   = regexp.MustCompile(`^\+?\d{1,20}$`)
	emailRegex   = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	bools = []string{
		"true",
		"false",
	}

	boolsInterfaces = stringSliceTOInterfaceSlice(bools)
)

type UserFiltersToValidate struct {
	firstName   string
	surname     string
	phoneNumber string
	active      string
	email       string
	isVerified  string
}

func stringSliceTOInterfaceSlice(values []string) []interface{} {
	interfaces := make([]interface{}, len(values))
	for i := range values {
		interfaces[i] = values[i]
	}

	return interfaces
}

func NewUserFiltersToValidate(firstName, surname, phoneNumber, email, active, isVerified string) *UserFiltersToValidate {
	return &UserFiltersToValidate{
		firstName:   firstName,
		surname:     surname,
		phoneNumber: phoneNumber,
		active:      active,
		email:       email,
		isVerified:  isVerified,
	}
}

func (userFilters *UserFiltersToValidate) Validate(ctx context.Context) *adminerror.ErrorResponse {
	if err := validation.ValidateStructWithContext(ctx, userFilters,
		validation.Field(&userFilters.firstName,
			validation.Match(lettersRegex).Error(errFieldAcceptsOnlyLetters),
		),
		validation.Field(&userFilters.surname,
			validation.Match(lettersRegex).Error(errFieldAcceptsOnlyLetters),
		),
		validation.Field(&userFilters.phoneNumber,
			validation.Match(phoneRegex).Error("номер телефона передан неверно, введите его в фромате 79123456789"),
		),
		validation.Field(&userFilters.isVerified,
			validation.In(boolsInterfaces...).Error(errBadBool),
		),
		validation.Field(&userFilters.active,
			validation.In(boolsInterfaces...).Error(errBadBool),
		),
		validation.Field(&userFilters.email,
			validation.Match(emailRegex).Error(errBadEmail),
		),
	); err != nil {
		return adminerror.CreateErrorResponse(err, 400)
	}

	return nil
}
