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
	GroupID  *string
}

// NewQueue
func NewQueue(c *internal.Common) (*Queue, error) {
	q := Queue{
		Common:  c,
		GroupID: aws.String(c.Env.SQSGroupID),
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

func (q *Queue) Push(msg string) error {
	_, err := q.SQS.SendMessage(&sqs.SendMessageInput{
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"message": {
				DataType:    aws.String("String"),
				StringValue: aws.String(msg),
			},
		},
		MessageBody:            aws.String("aaaaaaaaaa"),
		QueueUrl:               q.QueueURL.QueueUrl,
		MessageGroupId:         q.GroupID,
		MessageDeduplicationId: q.GroupID,
	})
	return err
}

func (q *Queue) Pop() (*sqs.Message, error) {
	msgOutput, err := q.SQS.ReceiveMessage(&sqs.ReceiveMessageInput{
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl: q.QueueURL.QueueUrl,
	})

	return msgOutput.Messages[0], err
}
