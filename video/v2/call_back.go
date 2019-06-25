package v2

type CallBack struct {
	Version   string      `json:"version"`
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}
