package bus

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

var (
	SC stan.Conn
)

func CreateSTANConnection() {
	id := uuid.New().String()
	SC, _ = stan.Connect("ticketing", id, stan.NatsURL("http://localhost:4222"))
	//handle error later
}

func CreateSTANListener(subject string, q string, fun func(m *stan.Msg)) {
	_, err := SC.QueueSubscribe(subject, q, fun, stan.DeliverAllAvailable(), stan.SetManualAckMode(), stan.DurableName(q), stan.AckWait(5*time.Second))
	if err != nil {
		fmt.Println(err)
	}
}

func STANPublish(subject string, data []byte) {
	SC.Publish(subject, data)
	//handle error later
}
