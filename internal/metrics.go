package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

// handler will export metrics
func (m Metrics) handler(w http.ResponseWriter, _ *http.Request) {
	e := export{
		NumberOfPublish: m.NumberOfPublish,
		LiveConnections: m.LiveConnections,
		DeadConnections: m.DeadConnections,
	}

	js, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

// serve metrics http server
func (m Metrics) serve(port int) {
	go func() {
		http.HandleFunc("/metrics", m.handler)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			log.Println(fmt.Errorf("failed to start metrics server error=%w", err))
		}
	}()
}
