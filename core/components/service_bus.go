package components

import (
  "encoding/json"
  "log"
  "time"

  "github.com/nats-io/nats.go"
)

type ServiceBusInterface interface {
  Connect(connectionString string)
  Connected() bool
  Close()
  Drain()
  Subscribe(topic string, handler ServiceBusHandler)
  Request(topic string, body interface{}, timeout time.Duration) (Response, error)
}

type ServiceBusMessageInterface interface {
  Respond(body interface{}) error
}

type ServiceBus struct {
  ServiceBusInterface
  conn *nats.Conn
}

type ServiceBusMessage struct {
  ServiceBusMessageInterface
  nats.Msg
}

type ServiceBusHandler = nats.MsgHandler

var serviceBus *ServiceBus

func GetServiceBus() *ServiceBus {
  if serviceBus == nil {
    serviceBus = &ServiceBus{}
  }
  return serviceBus
}

type Response struct {
  RawBody []byte
  Body    interface{}
  Error   string
}

type Empty = Response

func (s *ServiceBus) Connect(connectionString string) {
  conn, err := nats.Connect(connectionString)
  if err != nil {
    panic(err)
  }
  s.conn = conn
}

func (s *ServiceBus) Connected() bool {
  if s.conn == nil {
    return false
  }
  return s.conn.IsConnected()
}

func (s *ServiceBus) Subscribe(topic string, handler ServiceBusHandler) {
  s.conn.Subscribe(topic, handler)
}

func (s *ServiceBus) Drain() {
  if s.conn == nil {
    return
  }

  err := s.conn.Drain()
  if err != nil {
    log.Print(err)
  }
}

func (s *ServiceBus) Close() {
  if s.conn == nil {
    return
  }

  s.conn.Close()
}

func (s *ServiceBus) Request(topic string, body interface{}, timeout time.Duration) (Response, error) {
  buffer, err := json.Marshal(body)
  msg, err := s.conn.Request(topic, buffer, timeout)
  if err != nil {
    return Empty{}, err
  }

  var response = Response{RawBody: msg.Data}
  err = json.Unmarshal(msg.Data, &response.Body)
  return response, err
}

func (s *ServiceBusMessage) Respond(body interface{}) error {
  buffer, err := json.Marshal(body)
  if err != nil {
    return err
  }

  err = s.Respond(buffer)
  return err
}
