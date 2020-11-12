package utils

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"skeleton-code/config"
	"skeleton-code/logger"
	"time"
)

func NewKafkaProducer(kafkaCnf *config.KafkaWriteConfig) (*kafka.Writer, error) {
	conn, err := kafka.DialContext(context.Background(), "tcp", kafkaCnf.Brokers[0])
	if err != nil {
		logger.Errorf("create connection: %v", err)
		return nil, err
	}
	defer conn.Close()

	logger.Infof("kafka configuration: %+v", kafkaCnf)

	kafkaBrokers := kafkaCnf.Brokers
	if len(kafkaBrokers) == 0 {
		logger.Errorf("kafka brokers not specified!")
		return nil, err
	}

	kafkaTopic := kafkaCnf.Topic
	if len(kafkaTopic) == 0 {
		logger.Errorf("kafka topic not specified!")
		return nil, err
	}

	consumerConfig := kafka.WriterConfig{
		Brokers:      kafkaBrokers,
		Topic:        kafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: readKafkaTimeoutMS(3000),
		ReadTimeout:  readKafkaTimeoutMS(3000),
	}
	kafkaQ := kafka.NewWriter(consumerConfig)
	return kafkaQ, nil
}

func NewKafkaConsumer(kafkaCnf *config.KafkaReadConfig, startTime *time.Time) (*kafka.Reader, error) {
	conn, err := kafka.DialContext(context.Background(), "tcp", kafkaCnf.Brokers[0])
	if err != nil {
		logger.Errorf("create connection: %v", err)
		return nil, err
	}
	defer conn.Close()

	logger.Infof("kafka configuration: %+v", kafkaCnf)

	kafkaBrokers := kafkaCnf.Brokers
	if len(kafkaBrokers) == 0 {
		logger.Errorf("kafka brokers not specified!")
		return nil, err
	}

	kafkaTopic := kafkaCnf.Topic
	if len(kafkaTopic) == 0 {
		logger.Errorf("kafka topic not specified!")
		return nil, err
	}

	consumerConfig := kafka.ReaderConfig{
		Brokers:  kafkaBrokers,
		Topic:    kafkaTopic,
		GroupID:  kafkaCnf.GroupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	//startTime이 있을 경우에만 해당 시간 이후의 offset으로 시작
	if startTime != nil {
		err := setKafkaConsumerGroupOffset(conn, kafkaCnf, *startTime)
		if err != nil {
			logger.Errorf(" setKafkaConsumerGroupOffset error = %v", err)
		}
	}

	consumerConfig.MaxWait = readKafkaTimeoutMS(kafkaCnf.Timeout)
	consumerConfig.CommitInterval = readKafkaTimeoutMS(kafkaCnf.Timeout)
	kafkaQ := kafka.NewReader(consumerConfig)
	logger.Infof("start kafka status topic=%s, start offset=%v", kafkaTopic, kafkaQ.Offset())

	return kafkaQ, nil
}

func readKafkaTimeoutMS(ms int64) time.Duration {
	if ms <= 0 {
		return 10 * time.Second
	}
	return time.Duration(ms) * time.Millisecond
}

// 이 파일의 구현은 https://github.com/segmentio/kafka-go/issues/351의 의견 중 https://gist.github.com/tetafro/d46ca64e30b1a60b162ffe25ddef89e1에 게시된 소스를 기반으로 함
// SetKafkaConsumerGroupOffset Kafka Consumer Group의 offset을 특정 시간으로 설정한다. 해당 시점 이후로 메시지가 존재하지 않는다면 마지막 오프셋을 읽도록 한다.
func setKafkaConsumerGroupOffset(conn *kafka.Conn, kafkaCnf *config.KafkaReadConfig, time time.Time) error {

	// Read partitions list
	parts, err := conn.ReadPartitions(kafkaCnf.Topic)
	if err != nil {
		return fmt.Errorf("get partitions: %v", err)
	}
	// Get offsets by timestamp
	offsets := make(map[int]int64, len(parts))

	ctx := context.Background()
	for _, p := range parts {
		addr := fmt.Sprintf("%s:%d", p.Leader.Host, p.Leader.Port)

		c, err := kafka.DialLeader(ctx, "tcp", addr, kafkaCnf.Topic, p.ID)

		if err != nil {
			return fmt.Errorf("create connection to partition %d: %v", p.ID, err)
		}
		defer c.Close()

		offset, err := c.ReadOffset(time)

		if err != nil {
			return fmt.Errorf("read offset of partition %d: %v", p.ID, err)
		}
		// 추가된 부분 시작
		if offset == -1 {
			offset, err = c.ReadLastOffset()
		}
		if err != nil {
			return fmt.Errorf("read offset of partition %d: %v", p.ID, err)
		}
		// 추가된 부분 끝
		offsets[p.ID] = offset
	}

	// Set offsets for partitions
	group, err := kafka.NewConsumerGroup(kafka.ConsumerGroupConfig{
		Brokers: kafkaCnf.Brokers,
		Topics:  []string{kafkaCnf.Topic},
		ID:      kafkaCnf.GroupID,
	})

	if err != nil {
		return fmt.Errorf("create consumer group: %v", err)
	}

	defer group.Close()

	gen, err := group.Next(ctx)

	if err != nil {
		return fmt.Errorf("get next generation: %v", err)
	}

	err = gen.CommitOffsets(map[string]map[int]int64{kafkaCnf.Topic: offsets})

	if err != nil {
		return fmt.Errorf("commit offsets: %v", err)
	}

	return nil
}
