package stan

import (
	"consumer/internal/core/domain"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/nats-io/stan.go"
)

type Stan struct {
	conn stan.Conn
	sub  stan.Subscription
}

type Credentials struct {
	ClusterID string //os.Getenv("NATS_STAN_CLUSTER_ID"),
	ClientID  string
	StorePort string
}

func NewStanConnection(cr *Credentials) (*Stan, error) {
	natsUrl := fmt.Sprintf("nats://localhost:%s/", cr.StorePort)

	conn, err := stan.Connect(cr.ClusterID, cr.ClientID, stan.NatsURL(natsUrl))
	if err != nil {
		return nil, err
	}

	return &Stan{
		conn,
		nil,
	}, nil
}

func (nsc *Stan) Subscribe() error {
	channelName := "wb_orders"
	sub, err := nsc.conn.Subscribe(channelName, onMessage, stan.DeliverAllAvailable()) // stan.DeliverAllAvailable()

	if err != nil {
		return err
	}

	nsc.sub = sub

	return nil
}

func onMessage(m *stan.Msg) {
	var order *domain.Order
	err := json.Unmarshal(m.Data, &order)

	if err != nil {
		slog.Error("Error while unmarshaling the message", "error", err)
	}

	fmt.Println(order)
}

func (nsc *Stan) Close() {
	nsc.conn.Close()
	nsc.sub.Close()
}
