package cmd

import (
	"ecommerce-ums/helpers"
	"strings"

	"github.com/IBM/sarama"
)

func ServerKafkaConsumer() { 
	brokers := strings.Split(helpers.GetEnv("KAFKA_BROKER", "localhost:9092"),"")
	topic := helpers.GetEnv("KAFKA_TOPIC", "test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Fatal("Failed to connect to kafka: as consumer ", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		helpers.Logger.Fatal("Failed to create consumer partition 01 ", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		helpers.Logger.Info("Message Received ", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
}