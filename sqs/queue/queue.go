package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fiuskylab/aws-studies.git/internal"
)

// Queue
type Queue struct {
	Common   *internal.Common
	SQS      *sqs.SQS
	QueueURL *sqs.GetQueueUrlOutput
}

// NewQueue
func NewQueue(c *internal.Common) (*Queue, error) {
	q := Queue{
		Common: c,
	}

	return q.setSQS()
}

func (q *Queue) setSQS() (*Queue, error) {
	s := sqs.New(q.Common.Session)

	queueURL, err := s.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(q.Common.Env.SQSName),
	})

	if err != nil {
		q.Common.Logger.Error(err.Error())
		return q, err
	}
	q.QueueURL = queueURL
	q.SQS = s
	return q, nil
}
