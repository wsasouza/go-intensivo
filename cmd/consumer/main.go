package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/wsasouza/gointensivo/internal/infra/database"
	"github.com/wsasouza/gointensivo/internal/usecase"
	"github.com/wsasouza/gointensivo/pkg/kafka"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close() // executa tudo e depois executa o close

	repository := database.NewOrderrepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)

	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	fmt.Println("Kafka consumer has started")
	go kafka.Consume(topics, servers, msgChanKafka)
	kafkaWorker(msgChanKafka, usecase)

}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChan {
		// processar a mensagem
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDTO)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Kafka has processed order %s\n", outputDto.ID)
	}
}
