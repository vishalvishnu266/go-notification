package model

type KafkaData struct {
	To      string `json:"to,omitempty"`
	From    string `json:"from,omitempty"`
	Message string `json:message",omitempty"`
}
