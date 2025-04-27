package types

type MessageType string

type Message struct {
	From string `json:"from" validate:"required"`
	Type string `json:"type" validate:"required"`
	Body string `json:"body" validate:"required"`
}

type InteractSchema struct {
	Data []Message `json:"data" validate:"required,dive"`
}

type CallMeta struct {
	CallId string `json:"call_id"`
}

type ActionBody[T any] struct {
	Meta      CallMeta `json:"meta"`
	Arguments T        `json:"arguments"`
}

type ActionRequest struct {
	Meta     CallMeta `json:"meta"`
	Type     string   `json:"type"`
	Key      string   `json:"key"`
	Value    string   `json:"value"`
	Resolved bool     `json:"resolved"`
}

type InteractionSchemaValidationError struct {
	MessageIndex int    `json:"message_index"`
	Error        string `json:"error"`
	Type         string `json:"type"`
}
