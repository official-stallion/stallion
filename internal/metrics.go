package internal

type Metrics struct {
	LiveConnections int      `json:"live_connections"`
	DeadConnections int      `json:"dead_connections"`
	Topics          []string `json:"topics"`
}
