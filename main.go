package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var sess *session.Session
var svc *sqs.SQS
var wg sync.WaitGroup
var url string

func main() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = sqs.New(sess)
	url = "https://sqs.us-east-1.amazonaws.com/803551335240/batchem"

	// push()
	messages := poll()
	process(messages)

	fmt.Printf("[%s] Done\n", time.Now().Format("2006-01-02 15:04:05"))
}

func listQueues() string {
	r, err := svc.ListQueues(nil)
	check(err)
	return r.String()
}

func push() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.SendMessage(&sqs.SendMessageInput{
				MessageBody: aws.String("test item"),
				QueueUrl:    aws.String(url),
			})
			check(err)
		}()
	}
	wg.Wait()
}

func poll() []*sqs.Message {
	res, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(url),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(20),
	})
	check(err)
	return res.Messages
}

func process(messages []*sqs.Message) {
	for _, message := range messages {
		wg.Add(1)
		go func(m *sqs.Message) {
			fmt.Printf("%s\n", *m.Body)
			fmt.Printf("%s\n", *m.MessageId)
		}(message)
	}
	wg.Wait()
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
