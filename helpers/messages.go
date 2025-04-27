package helpers

import (
	"errors"
	"fmt"
	"salimon/archivist/types"

	"github.com/salimon-dev/gomsg"
)

func extractActionResults(schema *gomsg.InteractionSchema) (*[]string, error) {
	callIds := make([]string, len(schema.Data))
	index := 0
	for i, message := range schema.Data {
		if message.Type != "actionResult" {
			continue
		}
		if message.Meta == nil {
			return nil, errors.New(fmt.Sprintf("NoActionId:%d", i))
		}
		if message.Meta.ActionId == "" {
			return nil, errors.New(fmt.Sprintf("EmptyActionId:%d", i))
		}
		callIds[index] = message.Meta.ActionId
		index++
	}
	result := make([]string, index)
	for i := 0; i < index; i++ {
		result[i] = callIds[i]
	}
	return &result, nil
}

// func isActionResolved(message *gomsg.Message, callIds *[]string) (bool, error) {
// 	var body types.ActionBody[interface{}]
// 	err := json.Unmarshal([]byte(message.Body), &body)
// 	if err != nil {
// 		return false, err
// 	}
// 	callId := body.Meta.CallId
// 	for _, id := range *callIds {
// 		if id == callId {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }

func ExtractCalls(schema *gomsg.InteractionSchema) (*[]types.ActionRequest, error) {
	resolvedCalls, err := extractActionResults(schema)
	if err != nil {
		return nil, err
	}

	fmt.Println(resolvedCalls)
	calls := make([]types.ActionRequest, len(schema.Data))
	// index := 0

	return &calls, nil

	// for _, message := range schema.Data {
	// 	var call *types.ActionRequest
	// 	var err error
	// 	switch message.Type {
	// 	case "setStringValue":
	// 		resolved, err := isActionResolved(&message, resolvedCalls)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if !resolved {
	// 			call, err = parseSetStringValue(&message)
	// 		}
	// 		break
	// 	case "getStringValue":
	// 		resolved, err := isActionResolved(&message, resolvedCalls)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if !resolved {
	// 			call, err = parseGetStringValue(message)
	// 		}
	// 		break
	// 	case "removeStringValue":
	// 		resolved, err := isActionResolved(&message, resolvedCalls)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if !resolved {
	// 			call, err = praseRemoveStringValue(message)
	// 		}
	// 		break
	// 	}
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if call == nil {
	// 		continue
	// 	}
	// 	calls[index] = *call
	// 	index++

	// }
	// result := make([]types.ActionRequest, index)
	// for i := 0; i < index; i++ {
	// 	result[i] = calls[i]
	// }
	// return result, nil
}

// func parseSetStringValue(message *gomsg.Message) (*types.ActionRequest, error) {
// 	var result types.ActionRequest
// 	result.Type = "setStringValue"
// 	type Arguments struct {
// 		Key   string `json:"key"`
// 		Value string `json:"value"`
// 	}

// 	var body types.ActionBody[Arguments]

// 	err := json.Unmarshal([]byte(message.Body), &body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result.Meta = body.Meta
// 	result.Key = body.Arguments.Key
// 	result.Value = body.Arguments.Value
// 	return &result, nil
// }

// func parseGetStringValue(message gomsg.Message) (*types.ActionRequest, error) {
// 	var result types.ActionRequest
// 	result.Type = "getStringValue"

// 	var body types.ActionBody[string]

// 	err := json.Unmarshal([]byte(message.Body), &body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result.Meta = body.Meta
// 	result.Key = body.Arguments
// 	return &result, nil
// }

// func praseRemoveStringValue(message gomsg.Message) (*types.ActionRequest, error) {
// 	var result types.ActionRequest
// 	result.Type = "removeStringValue"

// 	var body types.ActionBody[string]

// 	err := json.Unmarshal([]byte(message.Body), &body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result.Meta = body.Meta
// 	result.Key = body.Arguments
// 	return &result, nil
// }

// // returns the call ID from action result
// func praseActionResult(message *gomsg.Message) (string, error) {
// 	var body types.ActionBody[string]
// 	err := json.Unmarshal([]byte(message.Body), &body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return body.Meta.CallId, nil
// }
