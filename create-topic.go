package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	bootstrapServers := "localhost:9092"
	topic := "topic-with-golang"
	numParts := 2

	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})

	if err != nil {
		fmt.Printf("Failet to create admin client: %s\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur, err := time.ParseDuration("60s")

	if err != nil {
		panic("ParseDuration(60s)")
	}

	results, err := admin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:         topic,
			NumPartitions: numParts,
		}},
		kafka.SetAdminOperationTimeout(maxDur))

	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		os.Exit(1)
	}

	for _, result := range results {
		fmt.Printf("%s\n", &result)
	}

	admin.Close()
}
