package controllers

import (
	"day4/tugas-crud-dinamis/config"
	"day4/tugas-crud-dinamis/lib/database"
	"day4/tugas-crud-dinamis/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all user functino
func GetUserController(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": users})
}

// find user by id
func GetUserControllerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := database.GetUserId(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get user by id",
		"user": users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// DB.Save for save to db
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete function
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DeleteUser(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	email := c.FormValue("email")
	name := c.FormValue("name")

	users := models.User{
		Name:  name,
		Email: email,
	}

	if err := config.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success update user by id",
		"user": users,
	})
}

func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(user)

	token, err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": token,
	})
}

// func GetUserDetailController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	users, err := database.GetDetailUser(uint(id))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"status": "succes",
// 		"user":   users,
// 	})
// }
