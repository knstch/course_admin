package usermanagement

import (
	"context"

	adminerror "github.com/knstch/course_admin/internal/app/admin_error"
	"github.com/knstch/course_admin/internal/app/validation"
	"github.com/knstch/course_admin/internal/domain/entity"
)

type UserManagementService struct {
	manager UserManager
}

type UserManager interface {
	GetAllUserData(ctx context.Context, firstName, surname, phoneNumber, email, active, isVerified string) ([]entity.UserData, *adminerror.ErrorResponse)
}

func NewUserManagementService() *UserManagementService {
	return &UserManagementService{}
}

func (user UserManagementService) GetAllUserData(ctx context.Context,
	firstName, surname, phoneNumber, email, active, isVerified string) ([]entity.UserData, *adminerror.ErrorResponse) {

	if err := validation.NewUserFiltersToValidate(firstName, surname, phoneNumber, email, active, isVerified).Validate(ctx); err != nil {
		return nil, err
	}

	userData, err := user.manager.GetAllUserData(ctx, firstName, surname, phoneNumber, email, active, isVerified)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
