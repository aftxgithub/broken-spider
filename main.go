package main

import (
	"log"
	"net"
	"time"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/thealamu/broken-spider/pkg/brokenspider"
)

func main() {
	setRoutes()
	srv := http.Server{
		Addr:    getAddrFromEnv(),
		Handler: http.DefaultServeMux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
	}

	log.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func setRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// serve the spider on /spider
	http.Handle("/spider", newSpiderHandler())
}

type spiderHandler struct {
	spider   *brokenspider.Spider
	upgrader websocket.Upgrader
}

func (s *spiderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// pull the url from query string
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	// upgrade the connection to a websocket
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade to websocket failed: ", err)
		http.Error(w, "upgrade to ws failed", http.StatusInternalServerError)
	}
	defer conn.Close()

	// walk the received url
	statusesChan := make(chan brokenspider.LinkStatus)
	go func() {
		log.Println("Walking", url)
		s.spider.Walk(url, statusesChan)
	}()

	for status := range statusesChan {
		log.Println(status)
		// send the status to the client
		conn.WriteJSON(status)
	}
}

func newSpiderHandler() *spiderHandler {
	return &spiderHandler{
		spider: brokenspider.New(),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func getAddrFromEnv() string {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	return net.JoinHostPort("", port)
}
