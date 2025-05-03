package actions

import (
	"salimon/archivist/db"
	"salimon/archivist/types"
	"time"

	"github.com/google/uuid"
	"github.com/salimon-dev/gomsg"
)

func HandleSetStringValueAction(request *gomsg.Message, user *types.User) *gomsg.Message {
	var result gomsg.Message
	result.Type = "actionResult"
	result.From = "archivist"
	result.Meta = request.Meta
	result.Result = &gomsg.ActionResult{}

	if request.Type != "setStringValue" {
		result.Result.Status = "failure"
		result.Result.Message = "invalid message type"
		return &result
	}
	if request.Parameters == nil {
		result.Result.Status = "failure"
		result.Result.Message = "missing parameters"
		return &result
	}
	if request.Parameters.RecordKey == nil {
		result.Result.Status = "failure"
		result.Result.Message = "missing parameter record_key"
		return &result
	}
	if request.Parameters.StringValue == nil {
		result.Result.Status = "failure"
		result.Result.Message = "missing parameter string_value"
		return &result
	}

	// check if record exists

	recordKey := request.Parameters.RecordKey
	data := request.Parameters.StringValue

	record, err := db.FindRecord("user_id = ? AND name = ?", user.Id, recordKey)

	if err != nil {
		result.Result.Status = "failure"
		result.Result.Message = "error finding record"
		return &result
	}

	if record == nil {
		record = &types.Record{
			Id:        uuid.New(),
			UserId:    user.Id,
			Network:   user.Network,
			Type:      types.RecordTypeString,
			Name:      *recordKey,
			Data:      *data,
			ExpiresAt: nil,
			CreateAt:  time.Now(),
			UpdatedAt: time.Now(),
		}
		db.InsertRecord(record)
	} else {
		record.Data = *request.Parameters.StringValue
		record.UpdatedAt = time.Now()
		db.UpdateRecord(record)
	}

	result.Result.Status = "success"
	result.Result.Message = "string value updated successfully"

	return &result
}
