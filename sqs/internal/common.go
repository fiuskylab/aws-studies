package internal

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Common struct {
	Session *session.Session
	Logger  *zap.Logger
	Env     Env
}

type Env struct {
	AWSRegion  string
	SQSName    string
	SQSGroupID string
}

func NewCommon() *Common {
	l, _ := zap.NewProduction()
	if err := godotenv.Load(); err != nil {
		l.Error(err.Error())
		return nil
	}

	c := Common{
		Env: Env{
			AWSRegion:  os.Getenv("AWS_REGION"),
			SQSName:    os.Getenv("SQS_NAME"),
			SQSGroupID: os.Getenv("SQS_GROUP_ID"),
		},
		Logger: l,
	}

	c.Session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(c.Env.AWSRegion),
		},
	}))

	return &c
}
