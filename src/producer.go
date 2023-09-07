package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/higor-tavares/helo-kafka/src/utils"
	"github.com/higor-tavares/helo-kafka/src/service"
)

var (
	message chan string
	port int
	baseUrl string
)

func init() {
	port = 8085
	baseUrl = fmt.Sprintf("http://localhost:%d",port)
}

func main() {
	message = make(chan string)
	defer close(message)
	go sendMessage(message)
	http.HandleFunc("/produce/", ProduceMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
func ProduceMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respondWith(w, http.StatusMethodNotAllowed, utils.Headers{
			"Allow":"POST",
		})
		return
	}
	msg := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(msg)
	message <- string(msg)
	respondWith(w, http.StatusOK, nil)
}

func sendMessage(messages <-chan string) {
	for message := range messages {
		service.SendToKafka(message)
		fmt.Printf("The message [%s] was sent successfuly\n", message)
	}
}
