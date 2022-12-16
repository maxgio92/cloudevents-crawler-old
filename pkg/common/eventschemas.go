package common

type EventID string

type CommonSource struct {
	SourceEvent `json:"source_event"`
}

type SourceEvent struct {
	ID   string      `json:"id"`
	Data interface{} `json:"data"`
}
