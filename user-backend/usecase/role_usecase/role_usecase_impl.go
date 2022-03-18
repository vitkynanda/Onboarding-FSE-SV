package role_usecase

import (
	"go-api/helpers"
	"go-api/models/dto"
)

func (role *roleUsecase) GetAllRole() dto.Response {
	roles, err := role.roleRepo.GetAllRole()

	

	if err != nil {
	return	helpers.ResponseError("Data not found", err.Error(), 404)
	}

	return helpers.ResponseSuccess("ok", nil, roles, 200)
}
