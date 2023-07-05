package controller

import (
	"net/http"
	"sadiq/Go_Rest_API/models"
	"sadiq/Go_Rest_API/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		response := models.Response{
			StatusCode:   http.StatusBadRequest,
			Status:       "error",
			Message:      "Invalid request payload",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := c.userService.CreateUser(&user)
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusInternalServerError,
			Status:       "Failure",
			Message:      "Failed to create user",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.Response{
		StatusCode:   http.StatusCreated,
		Status:       "success",
		Message:      "User created successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: []models.User{user},
		Error:        err,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusInternalServerError,
			Status:       "Failure",
			Message:      "Failed to get users",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	if len(users) == 0 {
		response := models.Response{
			StatusCode:   http.StatusNotFound,
			Status:       "error",
			Message:      "No users found",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err,
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := models.Response{
		StatusCode:   http.StatusOK,
		Status:       "success",
		Message:      "Users retrieved successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: users,
		Error:        err,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusNotFound,
			Status:       "error",
			Message:      "User not found",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := models.Response{
		StatusCode:   http.StatusOK,
		Status:       "success",
		Message:      "User retrieved successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: user,
		Error:        err,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := c.userService.GetUserByEmail(email)
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusNotFound,
			Status:       "error",
			Message:      "User not found",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := models.Response{
		StatusCode:   http.StatusOK,
		Status:       "success",
		Message:      "User retrieved successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: user,
		Error:        err,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusNotFound,
			Status:       "error",
			Message:      "User not found",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err,
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		response := models.Response{
			StatusCode:   http.StatusBadRequest,
			Status:       "error",
			Message:      "Invalid request payload",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err := c.userService.UpdateUser(user); err != nil {
		response := models.Response{
			StatusCode:   http.StatusInternalServerError,
			Status:       "Failure",
			Message:      "Failed to update user",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.Response{
		StatusCode:   http.StatusOK,
		Status:       "success",
		Message:      "User details updated successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: user,
		Error:        err,
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		response := models.Response{
			StatusCode:   http.StatusNotFound,
			Status:       "error",
			Message:      "User not found",
			RequestURL:   ctx.Request.URL.Path,
			ResponseData: nil,
			Error:        err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	c.userService.DeleteUser(userID)

	response := models.Response{
		StatusCode:   http.StatusOK,
		Status:       "success",
		Message:      "Below User details were deleted successfully",
		RequestURL:   ctx.Request.URL.Path,
		ResponseData: user,
		Error:        "",
	}

	ctx.JSON(http.StatusOK, response)
}
