package queue

import "github.com/IBM/sarama"

func ProduceKafkaMessage(topic, message string, producer sarama.SyncProducer) error {
	if producer == nil {
		return nil // no-op when Kafka isn't running
	}
	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	_, _, err := producer.SendMessage(msg)
	return err
}
