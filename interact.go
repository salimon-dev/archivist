package main

import (
	"fmt"
	"io"
	"net/http"
	"salimon/archivist/helpers"

	"github.com/labstack/echo/v4"
	"github.com/salimon-dev/gomsg"
)

func InteractHandler(ctx echo.Context) error {

	requestBody, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		fmt.Println(err)
		return helpers.InternalError(ctx)
	}

	schema, errs := gomsg.ParseInteractionSchema(requestBody)

	if errs != nil {
		return ctx.JSON(http.StatusBadRequest, errs)
	}

	// if err := ctx.Bind(payload); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	// }

	// // validation errors
	// vError, err := middlewares.ValidatePayload(*payload)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	// }
	// if vError != nil {
	// 	return ctx.JSON(http.StatusBadRequest, vError)
	// }

	// calls, err := helpers.ExtractCalls(payload)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return helpers.InternalError(ctx)
	// }

	// results := make([]types.Message, len(calls))
	// for index, call := range calls {
	// 	switch call.Type {
	// 	case "setStringValue":
	// 		result, err := actions.HandleSetStringValueAction(call)
	// 		if err != nil {
	// 			return helpers.InternalError(ctx)
	// 		}
	// 		results[index] = *result
	// 		break
	// 	}
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	return helpers.InternalError(ctx)
	// }

	return ctx.JSON(http.StatusOK, schema)
}
