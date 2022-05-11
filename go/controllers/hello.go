package controllers

import (
	"api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(ctx echo.Context) error {
	result := models.GetOne()
	return ctx.JSON(http.StatusOK, result)
}
