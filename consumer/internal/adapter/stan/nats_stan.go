package stan

import (
	"consumer/internal/core/domain"
	"consumer/internal/core/port"
	"consumer/internal/core/service"
	"context"
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

func (nsc *Stan) Subscribe(ctx context.Context, ordSvc *service.OrderService) error {
	channelName := "wb_orders"
	sub, err := nsc.conn.Subscribe(channelName, onMessage(ctx, ordSvc), stan.StartWithLastReceived()) // stan.DeliverAllAvailable()

	if err != nil {
		return err
	}

	nsc.sub = sub

	return nil
}

func onMessage(ctx context.Context, ordSvc port.OrderService) func(m *stan.Msg) {
	return func(m *stan.Msg) {
		order, err := UnmarshalMessage(m)

		if err != nil {
			slog.Error("Error while unmarshaling the message", "error", err)
		}

		err = ordSvc.SaveOrder(ctx, order)
		if err != nil {
			slog.Error("Error while creating the order", "error", err)
		}
	}
}

func UnmarshalMessage(m *stan.Msg) (*domain.Order, error) {
	order := domain.Order{}

	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		slog.Error("Error while unmarshalling message to model", "error", err.Error())
		return nil, err
	}

	return &order, nil
}

func (nsc *Stan) Close() {
	nsc.conn.Close()
	nsc.sub.Close()
}
