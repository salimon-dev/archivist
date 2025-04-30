package main

import (
	"fmt"
	"io"
	"net/http"
	"salimon/archivist/actions"
	"salimon/archivist/helpers"
	"salimon/archivist/types"

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

	actionMessages := gomsg.ExtractUnresolvedActionMessages(&schema.Data)

	if actionMessages == nil || len(actionMessages) == 0 {
		return ctx.JSON(http.StatusOK, gomsg.InteractionSchema{
			Data: []gomsg.Message{},
		})
	}

	data := make([]gomsg.Message, len(actionMessages))

	user := ctx.Get("user").(*types.User)

	for index, action := range actionMessages {
		switch action.Type {
		case "setStringValue":
			data[index] = *actions.HandleSetStringValueAction(action, user)
		case "getStringValue":
			data[index] = *actions.HandleGetStringValueAction(action, user)
		}
	}

	result := gomsg.InteractionSchema{Data: data}
	return ctx.JSON(http.StatusOK, result)
}
