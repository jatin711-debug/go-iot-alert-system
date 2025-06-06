package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaLogger struct {
	writer *kafka.Writer
	logger *zap.Logger
}

func NewKafkaLogger(brokers []string, topic string, logger *zap.Logger) *KafkaLogger {
	return &KafkaLogger{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Topic:                  topic,
			AllowAutoTopicCreation: true,
			Balancer:               &kafka.LeastBytes{},
			RequiredAcks:           kafka.RequireAll,
			Async:                  false,
			BatchTimeout:           500 * time.Millisecond,
		},
		logger: logger,
	}
}

func (k *KafkaLogger) WriteLog(ctx context.Context, key string, message string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
		Time:  time.Now(),
	}

	err := k.writer.WriteMessages(ctx, msg)
	if err != nil {
		k.logger.Error("failed to write log to Kafka", zap.Error(err))
		return err
	}

	k.logger.Info("log written to Kafka", zap.String("key", key))
	return nil
}

func (k *KafkaLogger) LogInfo(message string) {
	k.WriteLog(context.Background(), "INFO", message)
}

func (k *KafkaLogger) LogError(message string) {
	k.WriteLog(context.Background(), "ERROR", message)
}

func (k *KafkaLogger) LogDebug(message string) {
	k.WriteLog(context.Background(), "DEBUG", message)
}

func (k *KafkaLogger) Close() error {
	return k.writer.Close()
}
