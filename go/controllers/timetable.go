package controllers

import (
	"api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllTimetables(ctx echo.Context) error {
	timetables, err := models.FetchTimetables()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, timetables)
}

func GetOneTimetable(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	timetable, err := models.FetchTimetable(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, timetable)
}

func GetTimetableData(ctx echo.Context) error {
	timetable_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := models.FetchTimetableData(timetable_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, data)
}
