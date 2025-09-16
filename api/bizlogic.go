package api

import (
<<<<<<< HEAD
    "database/sql"
    "fmt"

    "github.com/varlaalekya/goproject/dataservice"
    "github.com/varlaalekya/goproject/model"
    "github.com/varlaalekya/goproject/queue"

    "github.com/IBM/sarama"
)


type IBizLogic interface {
	CreateOrderLogic(order model.Order) error
	UpdateOrderLogic(order model.Order) error
	DeleteOrderLogic(id int) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: producer}
}

func (bl *BizLogic) CreateOrderLogic(order model.Order) error {
	if err := dataservice.CreateOrder(bl.DB, order); err != nil {
		return err
	}

	msg := fmt.Sprintf("Order created for %s (id=%d)", order.Customer_Name, order.Id)
	if err := queue.ProduceKafkaMessage("Order_created_topic", msg, bl.Producer); err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}
	return nil
}

func (bl *BizLogic) UpdateOrderLogic(order model.Order) error {
	if err := dataservice.UpdateOrder(bl.DB, order); err != nil {
		return err
	}

	msg := fmt.Sprintf("Order updated (id=%d)", order.Id)
	if err := queue.ProduceKafkaMessage("Order_updated_topic", msg, bl.Producer); err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}
	return nil
}

func (bl *BizLogic) DeleteOrderLogic(id int) error {
	if err := dataservice.DeleteOrder(bl.DB, id); err != nil {
		return err
	}

	msg := fmt.Sprintf("Order deleted (id=%d)", id)
	if err := queue.ProduceKafkaMessage("Order_deleted_topic", msg, bl.Producer); err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}
	return nil
}
=======
	"context"
	"database/sql"

	"apigolang/dataservice"
	"apigolang/model"
)

type IBizLogic interface {
	CreateStudent(s model.Student) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateStudent(s model.Student) error {
	if err := dataservice.InsertStudent(context.Background(), bl.DB, s.Name, s.Age, s.Grade); err != nil {
		return err
	}
	return nil
}
>>>>>>> origin/main
