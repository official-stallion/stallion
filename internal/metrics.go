package internal

type (
	Metrics struct {
		NumberOfPublish int
		LiveConnections int
		DeadConnections int
	}

	export struct {
		NumberOfPublish int      `json:"number_of_publish"`
		LiveConnections int      `json:"live_connections"`
		DeadConnections int      `json:"dead_connections"`
		Topics          []string `json:"topics"`
	}
)
