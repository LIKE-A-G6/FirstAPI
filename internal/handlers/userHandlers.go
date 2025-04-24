package handlers

import (
	"Projects/internal/userService"
	"Projects/internal/web/users"
	"context"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (u UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := u.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &userToCreate.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (u UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	updateData := request.Body

	userToUpdate := userService.User{
		Email:    *updateData.Email,
		Password: *updateData.Password,
	}

	updatedUser, err := u.Service.UpdateUserByID(uint(id), userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.User{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return users.PatchUsersId200JSONResponse(response), nil

}

func (u UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	err := u.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
