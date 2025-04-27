package actions

import (
	"encoding/json"
	"errors"
	"salimon/archivist/types"
)

func HandleSetStringValueAction(request types.ActionRequest) (*types.Message, error) {
	if request.Type != "setStringValue" {
		return nil, errors.New("invalid action type")
	}
	if request.Key == "" {
		return nil, errors.New("key cannot be empty")
	}
	if request.Value == "" {
		return nil, errors.New("value cannot be empty")
	}
	var result types.Message
	result.Type = "actionResult"

	result.From = "archivist"

	type Arguments struct {
		Result  string `json:"result"`
		Message string `json:"message"`
	}

	body := types.ActionBody[Arguments]{
		Meta: request.Meta,
		Arguments: Arguments{
			Result:  "success",
			Message: "string value set successfully",
		},
	}
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	result.Body = string(bodyStr)
	return &result, nil
}
