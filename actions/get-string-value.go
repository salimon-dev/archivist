package actions

import (
	"salimon/archivist/db"
	"salimon/archivist/types"

	"github.com/salimon-dev/gomsg"
)

func HandleGetStringValueAction(request *gomsg.Message, user *types.User) *gomsg.Message {
	var result gomsg.Message
	result.Type = "actionResult"
	result.From = "archivist"
	result.Meta = request.Meta
	result.Result = &gomsg.ActionResult{}

	if request.Type != "getStringValue" {
		result.Result.Status = "failure"
		result.Result.Message = "invalid message type"
		return &result
	}
	if request.Parameters == nil {
		result.Result.Status = "failure"
		result.Result.Message = "missing parameters"
		return &result
	}
	if request.Parameters.RecordKey == "" {
		result.Result.Status = "failure"
		result.Result.Message = "missing parameter record_key"
		return &result
	}

	record, err := db.FindRecord("user_id = ? AND name = ?", user.Id, request.Parameters.RecordKey)

	if err != nil {
		result.Result.Status = "failure"
		result.Result.Message = "failed to find the record"
		return &result
	}

	if record == nil {
		result.Result.Status = "failure"
		result.Result.Message = "no such record found"
		return &result
	}

	result.Result.Status = "success"
	result.Result.Message = record.Data

	return &result
}
