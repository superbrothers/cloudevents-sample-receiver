package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudevents/sdk-go"
	"github.com/spf13/cobra"
)

const (
	defaultPort = 8080
	defaultPath = "/"
)

type ReceiverOptions struct {
	Port int
	Path string
}

func NewReceiverOptions() *ReceiverOptions {
	return &ReceiverOptions{
		Port: defaultPort,
		Path: defaultPath,
	}
}

func (o *ReceiverOptions) Run() error {
	ctx := context.Background()

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithPort(o.Port),
		cloudevents.WithPath(o.Path),
	)
	if err != nil {
		return fmt.Errorf("Failed to create transport: %v", err)
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		return fmt.Errorf("Failed to create client: %v", err)
	}

	log.Printf("Starting to listen on: :%d%s", o.Port, o.Path)
	if err := c.StartReceiver(ctx, gotEvent); err != nil {
		return fmt.Errorf("Failed to start receiver: %v", err)
	}

	return nil
}

func gotEvent(ctx context.Context, event cloudevents.Event) error {
	fmt.Printf("%+v\n", cloudevents.HTTPTransportContextFrom(ctx))
	fmt.Printf("Event:\n%s\n", event)
	fmt.Printf("----------------------------\n")

	return nil
}

func main() {
	opts := NewReceiverOptions()

	cmd := &cobra.Command{
		Use: fmt.Sprintf("cloudevents-sample-receiver [--port=%d] [--path=%s]", opts.Port, opts.Path),
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.Port, "port", opts.Port, "The port on which to run the receiver")
	cmd.Flags().StringVar(&opts.Path, "path", opts.Path, "The path on which to run the receiver")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
