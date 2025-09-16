package main

import (
	"database/sql"
	"log"
	"net/http"

	appis "github.com/varlaalekya/goproject/api" // package name is `appis` in /api

	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL DSN
	dsn := "root:root@tcp(127.0.0.1:3306)/orderservice?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Try to start Kafka producer, but continue if Kafka isn't running
	producer, err := initKafkaProducer()
	if err != nil {
		log.Printf("WARN: Kafka not available (%v). Continuing without producer.", err)
		producer = nil
	}
	if producer != nil {
		defer producer.Close()
	}

	// Register HTTP routes
	appis.RegisterRoutes(db, producer)

	// Start the HTTP server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	return sarama.NewSyncProducer(brokerList, config)
}

// command to check which ports are listening on windows: netstat -aon
