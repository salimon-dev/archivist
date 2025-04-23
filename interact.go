package main

import (
	"net/http"
	"salimon/archivist/middlewares"
	"salimon/archivist/types"

	"github.com/labstack/echo/v4"
)

func InteractHandler(ctx echo.Context) error {
	payload := new(types.InteractSchema)
	if err := ctx.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	// validation errors
	vError, err := middlewares.ValidatePayload(*payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	if vError != nil {
		return ctx.JSON(http.StatusBadRequest, vError)
	}

	return ctx.JSON(http.StatusOK, payload)
}
