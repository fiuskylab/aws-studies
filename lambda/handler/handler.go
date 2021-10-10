package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	QUEUE_URL  = `https://sqs.us-east-2.amazonaws.com/337201879866/dynamo-events`
	QUEUE_NAME = `dynamo-events`
	REGION     = `us-east-2`
)

func HandleRequest(ctx context.Context, event events.DynamoDBEvent) {
	if len(event.Records) == 0 {
		log.Println("Received 0 records")
		return
	}

	var err error

	awssession := session.Must(session.New(&aws.Config{
		Region:   aws.String(REGION),
		LogLevel: aws.LogLevel(aws.LogDebug),
	}), err)

	if err != nil {
		log.Println(err.Error())
		return
	}

	SQSClient := sqs.New(awssession)

	queueURL, err := SQSClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(QUEUE_NAME),
	})

	log.Println(*queueURL.QueueUrl)

	if err != nil {
		log.Println(err.Error())
		return
	}

	e := event.Records[0].Change.NewImage

	b, err := json.Marshal(e)

	if err != nil {
		log.Println(err.Error())
		return
	}

	if _, err = SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(b)),
		QueueUrl:    queueURL.QueueUrl,
	}); err != nil {
		log.Println(err.Error())
		return
	}
}

/*
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
* */
