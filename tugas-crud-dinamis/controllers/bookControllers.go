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
func GetBookController(c echo.Context) error {
	books, err := database.GetBook()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get all book",
		"book": books,
	})
}

// find user by id
func GetBookControllerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := database.GetBookId(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get book by id",
		"book": books,
	})
}

// create new user
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	// DB.Save for save to db
	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// delete function
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DeleteBook(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success delete book by id",
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	author := c.FormValue("author")

	books := models.Book{
		Title:  title,
		Author: author,
	}

	if err := config.DB.Where("id = ?", id).Updates(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success update book by id",
		"book": books,
	})
}
