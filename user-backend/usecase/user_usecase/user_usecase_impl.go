package user_usecase

import (
	"errors"
	"go-api/helpers"
	"go-api/models/dto"
	"go-api/models/entity"

	"gorm.io/gorm"
)

func (user *userUsecase) GetAllUsers() dto.Response {
	userlist, err := user.userRepo.GetAllUsers()
	response := []dto.UserList{}
	for _, user := range userlist {
	role := dto.Role{Id:user.RoleID, Title: user.Title}
	responseData := dto.UserList{
		Id : user.ID, 
		Name : user.Name, 
		Role: role,
		Active : user.Active,
	}	
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, response)
}

func (user *userUsecase) GetUserById(id string) dto.Response {
	userData, err := user.userRepo.GetUserById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	role := dto.Role{
		Id: userData.RoleID, 
		Title: userData.Title,
	}

	userResponse := map[string]interface{}{
		"id": userData.ID,
		"name": userData.Name,
		"email": userData.Email,
		"role": role,	
		"personalNumber": userData.Personal_number,
		"active": userData.Active,
	}
	return helpers.ResponseSuccess("ok", nil, userResponse)
}

func (user *userUsecase) CreateNewUser(newUser dto.User) dto.Response {
	userInsert := entity.User{
		ID: newUser.Id,
		Name: newUser.Name,
		Email: newUser.Email,
		Personal_number: newUser.Personal_number,
		Password: newUser.Password,
	}

	userData, _, err := user.userRepo.CreateNewUser(userInsert)
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID} )
}

func (user *userUsecase) UpdateUserData(userUpdate dto.User, id string) dto.Response {
	userInsert := entity.User{
		Name: userUpdate.Name,
		Email: userUpdate.Email,
		Personal_number: userUpdate.Personal_number,
		Active: userUpdate.Active,
		Password: userUpdate.Password,
		RoleID: userUpdate.Role.Id,
	}
	_, err := user.userRepo.UpdateUserData(userInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}

	userUpdate.Id = id
	return helpers.ResponseSuccess("ok", nil, map[string]interface{}{"id": id})
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {
	
 err := user.userRepo.DeleteUserById(id)
 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", err)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", err)
	}
	return helpers.ResponseSuccess("ok", nil, nil)
}
