package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

type Server interface {
	Handle(conn net.Conn)
}

// server is our broker service.
type server struct {
	auth    auth
	metrics *Metrics

	prefix int
	broker *broker
}

// NewServer returns a new broker server.
func NewServer(metrics int, user string, pass string) Server {
	s := &server{
		auth: auth{
			username: user,
			password: pass,
		},
		prefix: 101,
		metrics: &Metrics{
			NumberOfPublish: 0,
			LiveConnections: 0,
			DeadConnections: 0,
		},
	}

	// setting up the server broker and starting it
	s.broker = newBroker(
		make(chan message),
		make(chan subscribeChannel),
		make(chan unsubscribeChannel),
		make(chan int),
		s.metrics,
	)
	go s.broker.start()
	go s.serveMetrics(metrics)

	return s
}

// Handle will handle the clients.
func (s *server) Handle(conn net.Conn) {
	w := newWorker(
		s.prefix,
		s.auth,
		conn,
		make(chan message),
		s.broker.receiveChannel,
		s.broker.subscribeChannel,
		s.broker.unsubscribeChannel,
		s.broker.terminateChannel,
	)

	logInfo("new client joined", fmt.Sprintf("id=%d", s.prefix))

	s.metrics.LiveConnections++
	s.prefix++

	go w.start()
}

// metricsHandler will export metrics
func (s *server) metricsHandler(w http.ResponseWriter, _ *http.Request) {
	e := export{
		NumberOfPublish: s.metrics.NumberOfPublish,
		LiveConnections: s.metrics.LiveConnections,
		DeadConnections: s.metrics.DeadConnections,
		Topics:          s.broker.getTopics(),
	}

	js, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

// serveMetrics http server
func (s *server) serveMetrics(port int) {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", s.metricsHandler)

	log.Println(fmt.Sprintf("metrics server started on %d ...", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Println(fmt.Errorf("failed to start metrics server error=%w", err))
	}
}
