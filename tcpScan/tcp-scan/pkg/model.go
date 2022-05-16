package pkg

import (
	"fmt"
	"net"
	"time"
)

type Task struct {
	Host    string
	Port    string
	Timeout time.Duration
	Status  bool
	Error   error
}

func NewTask(host, port string, timeout ...time.Duration) *Task {
	task := &Task{
		Host:    host,
		Port:    port,
		Timeout: time.Second * 1,
	}
	if len(timeout) > 0 {
		task.Timeout = timeout[0]
	}
	return task
}

func (t *Task) Do() error {
	address := fmt.Sprintf("%s:%s", t.Host, t.Port)
	conn, err := net.DialTimeout("tcp", address, t.Timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
