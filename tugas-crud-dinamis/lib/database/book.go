package database

import (
	"day4/tugas-crud-dinamis/config"
	"day4/tugas-crud-dinamis/models"
)

func GetBook() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil

}

func GetBookId(id uint) (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books, id).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func DeleteBook(id uint) (interface{}, error) {
	var books []models.Book

	if err := config.DB.Delete(&books, id).Error; err != nil {
		return nil, err
	}

	return books, nil
}
