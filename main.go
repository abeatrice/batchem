package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sess *session.Session
var svc *sqs.SQS

func main() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = sqs.New(sess)
	fmt.Println(listQueues())
}

func listQueues() string {
	result, err := svc.ListQueues(nil)
	check(err)
	return result.String()
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
