package eventstore

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"kardexService/kardex"
	"kardexService/protobuf"
)

type gRpcClient struct {
	gRpcUri   string
	event     string
	aggregate string
}

func NewGRpcClient(gRpcUri string) kardex.EventStoreClient {
	return &gRpcClient{
		gRpcUri:   gRpcUri,
		event:     "order-submitted",
		aggregate: "kardex",
	}
}

func (g gRpcClient) Register(ctx context.Context, newKardex protobuf.RegisterKardexCommand) error {
	conn, err := grpc.Dial(g.gRpcUri, grpc.WithInsecure()) // grpc.WithInsecure() Only for testing without security (SSL)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := protobuf.NewEventStoreClient(conn)
	orderJSON, _ := json.Marshal(newKardex)
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	event := &protobuf.Event{
		EventId:       newUUID.String(),
		EventType:     g.event,
		AggregateId:   newKardex.KardexId,
		AggregateType: g.aggregate,
		EventData:     string(orderJSON),
		Channel:       g.event,
	}
	resp, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC kardexService")
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "Error from RPC kardexService")
	}
}
