package service

import (
	"time"
	"context"
	"github.com/segmentio/kafka-go"
)

const (
	topic         = "teste-kafka"
	brokerAddress = "localhost:9092"
	partition = 0
)

func SendToKafka(message string) {
	
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}