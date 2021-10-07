package main

import (
	"fmt"

	"github.com/fiuskylab/aws-studies.git/internal"
	"github.com/fiuskylab/aws-studies.git/sqs/queue"
)

func main() {
	c := internal.NewCommon()

	s, err := queue.NewQueue(c)

	if err != nil {
		return
	}

	if err := s.Push("salve"); err != nil {
		c.Logger.Error(err.Error())
	}

	msg, err := s.Pop()

	if err != nil {
		c.Logger.Error(err.Error())
	}

	fmt.Println(msg)
}
