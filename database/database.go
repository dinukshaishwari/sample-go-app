package database

import (
	"sample-go-app/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var svc *dynamodb.DynamoDB
var sess *session.Session

type DB struct {
	Client dynamodbiface.DynamoDBAPI
}

func GetClient() dynamodbiface.DynamoDBAPI {
	if svc != nil {
		return svc
	}

	svc = dynamodb.New(GetSession(), aws.NewConfig())

	return svc
}

func GetSession() *session.Session {
	if sess != nil {
		return sess
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.GetConfig().GetString("aws.region")),
	}))

	return sess
}
