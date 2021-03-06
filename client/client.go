package main

import (
	"../protos"
	"context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {
	// Create multiple clients and start receiving data
	var wg sync.WaitGroup

	wg.Add(1)


	clientId := int32(1) // TODO use db and real data somewhere in a distant future
	client, err := newClient(clientId)
	if err != nil {
		log.Fatal(err)
	}
	// Dispatch client goroutine
	go client.start()
	time.Sleep(time.Second * 2)

	// The wait group purpose is to avoid exiting, the clients do not exit
	wg.Wait()
}

// orderWatcherClient holds the long lived gRPC client fields
type orderWatcherClient struct {
	client protos.OrderWatcherClient // client is the long lived gRPC client
	conn   *grpc.ClientConn       // conn is the client gRPC connection
	id     int32                  // id is the client ID used for subscribing
}

// newClient creates a new client instance
func newClient(id int32) (*orderWatcherClient, error) {
	conn, err := mkConnection()
	if err != nil {
		return nil, err
	}
	return &orderWatcherClient{
		client: protos.NewOrderWatcherClient(conn),
		conn:   conn,
		id:     id,
	}, nil
}

// close is not used but is here as an example of how to close the gRPC client connection
func (c *orderWatcherClient) close() {
	if err := c.conn.Close(); err != nil {
		log.Fatal(err)
	}
}

// subscribe subscribes to messages from the gRPC server
func (c *orderWatcherClient) subscribe() (protos.OrderWatcher_SubscribeClient, error) {
	log.Printf("Subscribing client ID: %d", c.id)
	return c.client.Subscribe(context.Background(), &protos.Request{Id: c.id})
}

// unsubscribe unsubscribes to messages from the gRPC server
func (c *orderWatcherClient) unsubscribe() error {
	log.Printf("Unsubscribing client ID %d", c.id)
	_, err := c.client.Unsubscribe(context.Background(), &protos.Request{Id: c.id})
	return err
}

func (c *orderWatcherClient) start() {
	var err error
	// stream is the client side of the RPC stream
	var stream protos.OrderWatcher_SubscribeClient
	for {
		if stream == nil {
			if stream, err = c.subscribe(); err != nil {
				log.Printf("Failed to subscribe: %v", err)
				c.sleep()
				// Retry on failure
				continue
			}
		}
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			// Clearing the stream will force the client to resubscribe on next iteration
			stream = nil
			c.sleep()
			// Retry on failure
			continue
		}
		log.Printf("ETA: %d, Order: %d, Store: %d, Items: %+q\\n", response.Eta, response.Order, response.Store, response.Items)

	}
}

// sleep is used to give the server time to unsubscribe the client and reset the stream
func (c *orderWatcherClient) sleep() {
	time.Sleep(time.Second * 5)
}

func mkConnection() (*grpc.ClientConn, error) {
	return grpc.Dial("127.0.0.1:50051", []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}...)
}
