package connections

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConfig struct {
	Protocol string
	Login    string
	Password string
	Host     string
	Port     string
}

func NewRabbitConn(conf *RabbitConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s",
		conf.Protocol,
		conf.Login,
		conf.Password,
		conf.Host,
		conf.Port,
	))
	if err != nil {
		return nil, fmt.Errorf("conn to rabbit: %w", err)
	}

	return conn, nil
}
