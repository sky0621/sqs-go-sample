package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	ep := flag.String("ep", "endpoint", "endpoint")
	qn := flag.String("qn", "queuename", "queuename")
	flag.Parse()

	// Credentialは環境変数セット済の前提
	awsCfg := &aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: ep,
	}
	awsCfg.Credentials = credentials.NewEnvCredentials()

	sess, err := session.NewSession(awsCfg)
	if err != nil {
		panic(err)
	}

	cli := sqs.New(sess)

	getInput := &sqs.GetQueueUrlInput{QueueName: qn}
	gquRes, err := cli.GetQueueUrl(getInput)
	if err != nil {
		panic(err)
	}

	input := &sqs.SendMessageInput{
		QueueUrl:    gquRes.QueueUrl,
		MessageBody: aws.String("Test Message"),
	}
	output, err := cli.SendMessage(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
