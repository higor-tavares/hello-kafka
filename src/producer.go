package main 

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"net/http"
	"github.com/higor-tavares/hello-kafka/src/utils"
	"github.com/higor-tavares/hello-kafka/src/service"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

var (
	messageChanel chan string
	port int
	baseUrl string
	topic string
	conn *kafka.Conn
)

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	port,_ = strconv.Atoi(os.Getenv("PORT"))
	baseUrl = fmt.Sprintf("http://localhost:%d",port)
	topic = os.Getenv("SAMPLE_TOPIC")
}

func main() {
	var err error
	messageChanel = make(chan string)
	conn, err = service.SetUp(topic, 0);
	if err != nil {
		log.Fatal("Error while trying to connect with Kafk [%v]", err)
	}
	defer close(messageChanel)
	go messageListener(messageChanel)
	fmt.Printf("server is runing on %s ...\n",baseUrl)
	http.HandleFunc("/produce/", ProduceMessage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func ProduceMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.RespondWith(w, http.StatusMethodNotAllowed, utils.Headers{
			"Allow":"POST",
		})
		return
	}
	msg := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(msg)
	messageChanel <- string(msg)
	utils.RespondWith(w, http.StatusOK, nil)
}

func messageListener(messages <-chan string) {
	for message := range messages {
		err := service.SendToKafka(conn, message)
		if err != nil {
			fmt.Printf("Error while sending the message %s to kafka: %s", message, err)
		} else {
		fmt.Printf("The message [%s] was sent successfuly\n", message)
		}
	}
}
