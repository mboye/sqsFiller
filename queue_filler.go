package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	svc := sqs.New(
		session.New(),
		&aws.Config{Endpoint: aws.String("http://localhost:9324"),
			Region: aws.String("eu-central-1")})

	url := "http://localhost:9324/queue/requests"

	for i := 0; i < 1000; i++ {
		fmt.Printf("Inserting message %d\n", i)
		params := &sqs.SendMessageInput{
			MessageBody: aws.String(strconv.Itoa(i)),
			QueueUrl:    aws.String(url),
		}

		_, err := svc.SendMessage(params)
		if err != nil {
			println(err)
		}
	}
}
