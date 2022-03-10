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

	role := dto.Role{Id:user.RoleId}

	responseData := dto.UserList{
		Id : user.Id, 
		Name : user.Name, 
		Role: role,
		Active : user.Active,
	}
		
		response = append(response, responseData)
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Get all users successfully", 200, response)
}

func (user *userUsecase) GetUserById(id string) dto.Response {
	userData, err := user.userRepo.GetUserById(id)
	
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("Get user successfully", 200, userData)
}

func (user *userUsecase) CreateNewUser(newUser dto.User) dto.Response {
	

	userInsert := entity.User{
		Id: newUser.Id,
		Email: newUser.Email,
		Personal_number: newUser.Personal_number,
		Active: newUser.Active,
	}

	// title :="viewer"
	// InsertRole := entity.Role{
	// 	Title: title ,
	// }
		
	userData, err := user.userRepo.CreateNewUser(userInsert)
	
	 if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	newUser.Id = userData.Id
	return helpers.ResponseSuccess("New user created successfully", 200, newUser)
}

func (user *userUsecase) UpdateUserData(userUpdate dto.User, id string) dto.Response {
	
	userInsert := entity.User{
		Id: userUpdate.Id,
		Email: userUpdate.Email,
		Personal_number: userUpdate.Personal_number,
		Active: userUpdate.Active,
	}
		
	_, err := user.userRepo.UpdateUserData(userInsert, id)
	 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}

	return helpers.ResponseSuccess("User data updated successfully", 200, userUpdate)
}

func (user *userUsecase) DeleteUserById(id string) dto.Response {
	
 err := user.userRepo.DeleteUserById(id)
 
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return helpers.ResponseError("Data not found", 404)
	} else if err != nil {
		return helpers.ResponseError("Internal server error", 500)
	}
	return helpers.ResponseSuccess("User deleted successfully", 200, nil)
}
