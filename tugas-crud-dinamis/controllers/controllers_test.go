package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	e := echo.New()

	t.Run("Get User 2xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users")

		// u := &UsersModel{}
		// h := NewHandler(u)

		// var a string
		// a = `{"name": "insan"}`

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// assert.Equal(t, a, rec.Body.String())
			// assert.Equal(t, userJSON, rec.Body.String())
		}

	})

	t.Run("Create User 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/")

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Update User 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.PUT, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Get User id 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, GetUserControllerById(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Delete User id 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}

func TestGetUser_Invalid(t *testing.T) {
	e := echo.New()

	t.Run("Get User 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users")

		// u := &UsersModel{}
		// h := NewHandler(u)

		// var a string
		// a = `{"name": "insan"}`

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			// assert.Equal(t, a, rec.Body.String())
			// assert.Equal(t, userJSON, rec.Body.String())
		}

	})

	t.Run("Create User 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/")

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Update User 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.PUT, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Get User id 4xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, GetUserControllerById(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Delete User 4d 2xx", func(t *testing.T) {
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("v1/users/:id")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

}
