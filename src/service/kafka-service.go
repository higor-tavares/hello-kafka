package service

import (
	"time"
	"context"
	"github.com/segmentio/kafka-go"
	"os"
)

var (
	brokerAddr string
)
func init(){
	brokerAddr = os.Getenv("KAFKA_BROKER")
}

func SetUp(topic string, partition int) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddr, topic, partition)
	conn.SetWriteDeadline(time.Now().Add(60*time.Second))
	return conn, err
}
 
func SendToKafka(conn *kafka.Conn, message string) error {
	_, err := conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)
	return err
}